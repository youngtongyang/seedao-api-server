package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/seedao/seedao-api-server/internal/config"
	"github.com/seedao/seedao-api-server/internal/model"
	"github.com/seedao/seedao-api-server/internal/pkg/alipay"
	"github.com/seedao/seedao-api-server/internal/pkg/redisstore"
	"github.com/seedao/seedao-api-server/internal/pkg/wechatpay"
)

const paymentNotifyLockTTL = 30 * time.Second

var (
	ErrPaymentOrderNotFound   = errors.New("payment order not found")
	ErrPaymentOrderForbidden  = errors.New("payment order forbidden")
	ErrPaymentOrderNotPending = errors.New("payment order not pending")
	ErrPaymentProductInvalid  = errors.New("payment product invalid")
	ErrPaymentOrderExpired    = errors.New("payment order expired")
	ErrPaymentChannelInvalid  = errors.New("payment channel invalid")
)

type PaymentService struct {
	db      *gorm.DB
	redis   *redisstore.Store
	wechat  wechatpay.Client
	alipay  alipay.Client
	cfg     config.WechatPayConfig
	devMode bool
	coupon  *CouponService
	course  *CourseService
}

func NewPaymentService(db *gorm.DB, redis *redisstore.Store, wechat wechatpay.Client, alipayClient alipay.Client, cfg config.WechatPayConfig, devMode bool, coupon *CouponService, course *CourseService) *PaymentService {
	return &PaymentService{
		db:      db,
		redis:   redis,
		wechat:  wechat,
		alipay:  alipayClient,
		cfg:     cfg,
		devMode: devMode,
		coupon:  coupon,
		course:  course,
	}
}

type CreatePaymentOrderInput struct {
	ProductType string
	Days        int
	PayChannel  string
	ClientIP    string
	CouponCode  string
}

type PaymentOrderView struct {
	OrderNo             string `json:"order_no"`
	ProductType         string `json:"product_type"`
	ProductDays         int    `json:"product_days"`
	PayChannel          string `json:"pay_channel"`
	OriginalAmount      int    `json:"original_amount"`
	DiscountAmount      int    `json:"discount_amount,omitempty"`
	CouponCode          string `json:"coupon_code,omitempty"`
	Amount              int    `json:"amount"`
	Currency            string `json:"currency"`
	Status              string `json:"status"`
	CodeURL             string `json:"code_url,omitempty"`
	H5URL               string `json:"h5_url,omitempty"`
	PayURL              string `json:"pay_url,omitempty"`
	ExpiredAt           string `json:"expired_at"`
	PaidAt              string `json:"paid_at,omitempty"`
	MembershipExpiresAt string `json:"membership_expires_at,omitempty"`
	WxTransactionID     string `json:"wx_transaction_id,omitempty"`
}

func (s *PaymentService) PreviewCoupon(input CouponQuoteInput) (*CouponQuote, error) {
	productType := input.ProductType
	if productType == "" {
		productType = model.ProductTypeMembershipRenew
	}
	if productType != model.ProductTypeMembershipRenew {
		return nil, ErrPaymentProductInvalid
	}

	days := input.Days
	if days <= 0 {
		days = model.DefaultMembershipDays
	}

	input.ProductType = productType
	input.Days = days

	originalAmount := s.calcAmountFen(days)
	return s.coupon.Preview(input, originalAmount)
}

func (s *PaymentService) CreateOrder(ctx context.Context, userID int64, input CreatePaymentOrderInput) (*PaymentOrderView, error) {
	productType := input.ProductType
	if productType == "" {
		productType = model.ProductTypeMembershipRenew
	}
	if productType != model.ProductTypeMembershipRenew {
		return nil, ErrPaymentProductInvalid
	}

	days := input.Days
	if days <= 0 {
		days = model.DefaultMembershipDays
	}

	payChannel := input.PayChannel
	if payChannel == "" {
		payChannel = model.PayChannelWechatNative
	}
	if !model.ValidPayChannel(payChannel) {
		return nil, ErrPaymentChannelInvalid
	}

	var user model.User
	if err := s.db.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	originalAmountFen := s.calcAmountFen(days)
	finalAmountFen := originalAmountFen
	var discountAmount int64
	var storedCouponCode *string

	if input.CouponCode != "" {
		discount, code, err := s.couponReserveQuote(userID, productType, days, input.CouponCode, originalAmountFen)
		if err != nil {
			return nil, err
		}
		discountAmount = discount
		finalAmountFen = originalAmountFen - discountAmount
		storedCouponCode = &code
	}

	now := time.Now()
	expiredAt := now.Add(time.Duration(s.cfg.OrderExpireMinutes) * time.Minute)
	orderNo := generateMembershipOrderNo(userID, now)

	var payURL string
	switch payChannel {
	case model.PayChannelWechatNative:
		nativeResult, err := s.wechat.CreateNativeOrder(ctx, wechatpay.NativeOrderInput{
			OrderNo:     orderNo,
			Description: s.cfg.Description,
			AmountFen:   finalAmountFen,
			ExpireAt:    expiredAt,
		})
		if err != nil {
			return nil, err
		}
		payURL = nativeResult.CodeURL
	case model.PayChannelWechatH5:
		h5Result, err := s.wechat.CreateH5Order(ctx, wechatpay.H5OrderInput{
			OrderNo:     orderNo,
			Description: s.cfg.Description,
			AmountFen:   finalAmountFen,
			ExpireAt:    expiredAt,
			ClientIP:    input.ClientIP,
		})
		if err != nil {
			return nil, err
		}
		payURL = h5Result.H5URL
	case model.PayChannelAlipayH5:
		wapResult, err := s.alipay.CreateWapOrder(ctx, alipay.WapOrderInput{
			OrderNo:   orderNo,
			Subject:   s.cfg.Description,
			AmountFen: finalAmountFen,
			ExpireAt:  expiredAt,
		})
		if err != nil {
			return nil, err
		}
		payURL = wapResult.PayURL
	}

	order := model.PaymentOrder{
		OrderNo:        orderNo,
		UserID:         &userID,
		ProductType:    productType,
		ProductDays:    days,
		OriginalAmount: int(originalAmountFen),
		DiscountAmount: int(discountAmount),
		CouponCode:     storedCouponCode,
		Amount:         int(finalAmountFen),
		Currency:       "CNY",
		Status:         model.PaymentOrderStatusPending,
		PayChannel:     payChannel,
		CodeURL:        &payURL,
		ExpiredAt:      expiredAt,
		IsAnonymous:    false,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&order).Error; err != nil {
			return err
		}
		if input.CouponCode == "" {
			return nil
		}
		_, _, err := s.coupon.Reserve(tx, CouponQuoteInput{
			UserID:      userID,
			UserKnown:   true,
			ProductType: productType,
			Days:        days,
			CouponCode:  input.CouponCode,
		}, orderNo, order.ID, originalAmountFen)
		return err
	})
	if err != nil {
		_ = s.closeProviderOrder(ctx, payChannel, orderNo)
		return nil, err
	}

	return toPaymentOrderView(&order, ""), nil
}

func (s *PaymentService) couponReserveQuote(userID int64, productType string, days int, couponCode string, originalAmount int64) (int64, string, error) {
	quote, err := s.coupon.Preview(CouponQuoteInput{
		UserID:      userID,
		UserKnown:   true,
		ProductType: productType,
		Days:        days,
		CouponCode:  couponCode,
	}, originalAmount)
	if err != nil {
		return 0, "", err
	}
	return int64(quote.DiscountAmount), quote.CouponCode, nil
}

func (s *PaymentService) GetOrder(userID int64, orderNo string) (*PaymentOrderView, error) {
	order, err := s.findOwnedMembershipOrder(userID, orderNo)
	if err != nil {
		return nil, err
	}

	membershipExpiresAt := ""
	if order.Status == model.PaymentOrderStatusPaid {
		var user model.User
		if order.UserID != nil {
			if err := s.db.First(&user, *order.UserID).Error; err == nil {
				membershipExpiresAt = formatMembershipExpiry(user.MembershipExpiresAt)
			}
		}
	}

	return toPaymentOrderView(order, membershipExpiresAt), nil
}

func (s *PaymentService) ListOrders(userID int64, limit int) ([]PaymentOrderView, error) {
	if limit <= 0 || limit > 50 {
		limit = 20
	}

	var orders []model.PaymentOrder
	if err := s.db.Where("user_id = ? AND product_type = ?", userID, model.ProductTypeMembershipRenew).
		Order("id DESC").
		Limit(limit).
		Find(&orders).Error; err != nil {
		return nil, err
	}

	out := make([]PaymentOrderView, 0, len(orders))
	for i := range orders {
		out = append(out, *toPaymentOrderView(&orders[i], ""))
	}
	return out, nil
}

func (s *PaymentService) CloseOrder(ctx context.Context, userID int64, orderNo string) (*PaymentOrderView, error) {
	order, err := s.findOwnedMembershipOrder(userID, orderNo)
	if err != nil {
		return nil, err
	}
	if order.Status != model.PaymentOrderStatusPending {
		return nil, ErrPaymentOrderNotPending
	}

	if err := s.closeProviderOrder(ctx, order.PayChannel, orderNo); err != nil {
		return nil, err
	}

	now := time.Now()
	order.Status = model.PaymentOrderStatusClosed
	order.ClosedAt = &now
	order.UpdatedAt = now
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(order).Error; err != nil {
			return err
		}
		return s.coupon.ReleaseByOrderID(tx, order.ID)
	}); err != nil {
		return nil, err
	}

	return toPaymentOrderView(order, ""), nil
}

func (s *PaymentService) HandleWechatNotify(ctx context.Context, req *http.Request) error {
	notifyResult, err := s.wechat.ParseNotify(req)
	if err != nil {
		return err
	}
	return s.CompleteOrderPayment(ctx, notifyResult.OutTradeNo, notifyResult.TransactionID, notifyResult.AmountFen, notifyResult.RawJSON, notifyResult.SuccessTime)
}

func (s *PaymentService) HandleAlipayNotify(ctx context.Context, req *http.Request) error {
	notifyResult, err := s.alipay.ParseNotify(req)
	if err != nil {
		return err
	}
	return s.CompleteOrderPayment(ctx, notifyResult.OutTradeNo, notifyResult.TransactionID, notifyResult.AmountFen, notifyResult.RawJSON, notifyResult.SuccessTime)
}

func (s *PaymentService) closeProviderOrder(ctx context.Context, payChannel, orderNo string) error {
	switch payChannel {
	case model.PayChannelAlipayH5:
		return s.alipay.CloseOrder(ctx, orderNo)
	default:
		return s.wechat.CloseOrder(ctx, orderNo)
	}
}

func (s *PaymentService) SimulatePayDev(ctx context.Context, userID int64, orderNo string) (*PaymentOrderView, error) {
	if !s.devMode {
		return nil, ErrInvalidParam
	}

	order, err := s.findOwnedMembershipOrder(userID, orderNo)
	if err != nil {
		return nil, err
	}
	if order.Status != model.PaymentOrderStatusPending {
		return nil, ErrPaymentOrderNotPending
	}
	if time.Now().After(order.ExpiredAt) {
		return nil, ErrPaymentOrderExpired
	}

	txID := "dev_" + orderNo
	raw := fmt.Sprintf(`{"dev_mode":true,"order_no":"%s"}`, orderNo)
	if err := s.CompleteOrderPayment(ctx, orderNo, txID, int64(order.Amount), raw, time.Now()); err != nil {
		return nil, err
	}

	return s.GetOrder(userID, orderNo)
}

func (s *PaymentService) CompleteOrderPayment(ctx context.Context, orderNo, wxTransactionID string, amountFen int64, rawJSON string, paidAt time.Time) error {
	acquired, err := s.redis.TryPaymentNotifyLock(ctx, orderNo, paymentNotifyLockTTL)
	if err != nil {
		return err
	}
	if !acquired {
		return nil
	}
	defer func() {
		_ = s.redis.ReleasePaymentNotifyLock(context.Background(), orderNo)
	}()

	return s.db.Transaction(func(tx *gorm.DB) error {
		var order model.PaymentOrder
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("order_no = ?", orderNo).
			First(&order).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrPaymentOrderNotFound
			}
			return err
		}

		if order.Status == model.PaymentOrderStatusPaid {
			return nil
		}
		if order.Status != model.PaymentOrderStatusPending {
			return ErrPaymentOrderNotPending
		}
		if int64(order.Amount) != amountFen {
			return fmt.Errorf("payment amount mismatch: order=%d notify=%d", order.Amount, amountFen)
		}

		now := time.Now()
		switch order.ProductType {
		case model.ProductTypeMembershipRenew:
			if order.UserID == nil {
				return fmt.Errorf("membership order missing user_id")
			}
			var user model.User
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&user, *order.UserID).Error; err != nil {
				return err
			}
			expiresAt := ExtendMembershipExpiry(user.MembershipExpiresAt, now, order.ProductDays)
			user.MembershipExpiresAt = &expiresAt
			user.UpdatedAt = now
			if err := tx.Save(&user).Error; err != nil {
				return err
			}
			if err := s.coupon.ConsumeByOrderID(tx, order.ID); err != nil {
				return err
			}
		case model.ProductTypeCoursePurchase:
			if err := s.course.FulfillPaidOrder(tx, &order); err != nil {
				return err
			}
		default:
			return ErrPaymentProductInvalid
		}

		order.Status = model.PaymentOrderStatusPaid
		order.WxTransactionID = &wxTransactionID
		order.PaidAt = &paidAt
		order.NotifyPayload = &rawJSON
		order.UpdatedAt = now
		return tx.Save(&order).Error
	})
}

func (s *PaymentService) findOwnedMembershipOrder(userID int64, orderNo string) (*model.PaymentOrder, error) {
	var order model.PaymentOrder
	err := s.db.Where("order_no = ? AND product_type = ?", orderNo, model.ProductTypeMembershipRenew).
		First(&order).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrPaymentOrderNotFound
	}
	if err != nil {
		return nil, err
	}
	if order.UserID == nil || *order.UserID != userID {
		return nil, ErrPaymentOrderForbidden
	}
	return &order, nil
}

func (s *PaymentService) calcAmountFen(days int) int64 {
	base := s.cfg.MembershipPriceFen
	if days == model.DefaultMembershipDays {
		return base
	}
	return base * int64(days) / int64(model.DefaultMembershipDays)
}

func generateMembershipOrderNo(userID int64, now time.Time) string {
	buf := make([]byte, 4)
	_, _ = rand.Read(buf)
	return fmt.Sprintf("M%s%06d%s", now.Format("20060102150405"), userID%1000000, hex.EncodeToString(buf))
}

func toPaymentOrderView(order *model.PaymentOrder, membershipExpiresAt string) *PaymentOrderView {
	originalAmount := order.OriginalAmount
	if originalAmount == 0 {
		originalAmount = order.Amount
	}

	view := &PaymentOrderView{
		OrderNo:        order.OrderNo,
		ProductType:    order.ProductType,
		ProductDays:    order.ProductDays,
		PayChannel:     order.PayChannel,
		OriginalAmount: originalAmount,
		DiscountAmount: order.DiscountAmount,
		Amount:         order.Amount,
		Currency:       order.Currency,
		Status:         order.Status,
		ExpiredAt:      order.ExpiredAt.UTC().Format(time.RFC3339),
	}
	if order.CouponCode != nil {
		view.CouponCode = *order.CouponCode
	}
	if order.CodeURL != nil && order.Status == model.PaymentOrderStatusPending {
		switch order.PayChannel {
		case model.PayChannelWechatH5:
			view.H5URL = *order.CodeURL
		case model.PayChannelAlipayH5:
			view.PayURL = *order.CodeURL
		default:
			view.CodeURL = *order.CodeURL
		}
	}
	if order.PaidAt != nil {
		view.PaidAt = order.PaidAt.UTC().Format(time.RFC3339)
	}
	if order.WxTransactionID != nil {
		view.WxTransactionID = *order.WxTransactionID
	}
	if membershipExpiresAt != "" {
		view.MembershipExpiresAt = membershipExpiresAt
	}
	return view
}
