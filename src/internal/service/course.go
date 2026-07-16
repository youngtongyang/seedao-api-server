package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/seedao/seedao-api-server/internal/config"
	"github.com/seedao/seedao-api-server/internal/model"
	"github.com/seedao/seedao-api-server/internal/pkg/alipay"
	"github.com/seedao/seedao-api-server/internal/pkg/ordertoken"
	"github.com/seedao/seedao-api-server/internal/pkg/wechatpay"
)

var (
	ErrCourseNotFound                    = errors.New("course not found")
	ErrCourseSKUNotFound                 = errors.New("course sku not found")
	ErrCourseOffSale                     = errors.New("course off sale")
	ErrCourseMemberDiscountRequiresLogin = errors.New("course member discount requires login")
	ErrCourseOrderTokenRequired          = errors.New("course order token required")
	ErrCourseOrderTokenInvalid           = errors.New("course order token invalid")
	ErrCourseOrderAlreadyClaimed         = errors.New("course order already claimed")
)

type CourseService struct {
	db         *gorm.DB
	wechat     wechatpay.Client
	alipay     alipay.Client
	membership config.MembershipConfig
	devMode    bool
}

func NewCourseService(db *gorm.DB, wechat wechatpay.Client, alipayClient alipay.Client, membership config.MembershipConfig, devMode bool) *CourseService {
	return &CourseService{db: db, wechat: wechat, alipay: alipayClient, membership: membership, devMode: devMode}
}

type CourseProductRef struct {
	CourseID string `json:"course_id"`
	SKUID    string `json:"sku_id"`
}

type CourseCheckoutInput struct {
	CourseID            string
	SKUID               string
	ChannelCode         string
	ApplyMemberDiscount bool
	UserID              int64
	UserKnown           bool
}

type CourseDiscountLine struct {
	Type    string `json:"type"`
	Label   string `json:"label"`
	Amount  int    `json:"amount"`
	Applied bool   `json:"applied"`
}

type CourseMemberDiscountPreview struct {
	Amount         int    `json:"amount"`
	Applied        bool   `json:"applied"`
	RequiresLogin  bool   `json:"requires_login"`
	Hint           string `json:"hint,omitempty"`
}

type CourseCheckoutPreview struct {
	CourseID       string                      `json:"course_id"`
	SKUID          string                      `json:"sku_id"`
	OriginalAmount int                         `json:"original_amount"`
	Discounts      []CourseDiscountLine        `json:"discounts"`
	FinalAmount    int                         `json:"final_amount"`
	AmountIfMember int                         `json:"amount_if_member"`
	MemberDiscount CourseMemberDiscountPreview `json:"member_discount"`
	Currency       string                      `json:"currency"`
}

type CreateCourseOrderInput struct {
	CourseCheckoutInput
	PayChannel string
	ClientIP   string
}

type CourseOrderView struct {
	OrderNo        string               `json:"order_no"`
	OrderToken     string               `json:"order_token,omitempty"`
	ProductType    string               `json:"product_type"`
	CourseID       string               `json:"course_id"`
	SKUID          string               `json:"sku_id"`
	PayChannel     string               `json:"pay_channel"`
	OriginalAmount int                  `json:"original_amount"`
	DiscountAmount int                  `json:"discount_amount"`
	Discounts      []CourseDiscountLine `json:"discounts,omitempty"`
	Amount         int                  `json:"amount"`
	Currency       string               `json:"currency"`
	Status         string               `json:"status"`
	CodeURL        string               `json:"code_url,omitempty"`
	H5URL          string               `json:"h5_url,omitempty"`
	PayURL         string               `json:"pay_url,omitempty"`
	ExpiredAt      string               `json:"expired_at"`
	PaidAt         string               `json:"paid_at,omitempty"`
	WxTransactionID string              `json:"wx_transaction_id,omitempty"`
	EnrollmentStatus string             `json:"enrollment_status,omitempty"`
}

type courseDiscountSnapshot struct {
	Items              []CourseDiscountLine `json:"items"`
	AmountIfNonMember  int                  `json:"amount_if_non_member"`
	PricedAt           string               `json:"priced_at"`
}

type coursePricingResult struct {
	OriginalAmount int
	DiscountAmount int
	FinalAmount    int
	AmountIfMember int
	Discounts      []CourseDiscountLine
	MemberPreview  CourseMemberDiscountPreview
}

func (s *CourseService) PreviewCheckout(input CourseCheckoutInput) (*CourseCheckoutPreview, error) {
	pricing, err := s.calculatePricing(input)
	if err != nil {
		return nil, err
	}

	finalAmount := pricing.FinalAmount
	if input.ApplyMemberDiscount && input.UserKnown && pricing.MemberPreview.Applied {
		finalAmount = pricing.AmountIfMember
	} else if !input.ApplyMemberDiscount || !input.UserKnown {
		finalAmount = pricing.FinalAmount
	}

	return &CourseCheckoutPreview{
		CourseID:       input.CourseID,
		SKUID:          input.SKUID,
		OriginalAmount: pricing.OriginalAmount,
		Discounts:      pricing.Discounts,
		FinalAmount:    finalAmount,
		AmountIfMember: pricing.AmountIfMember,
		MemberDiscount: pricing.MemberPreview,
		Currency:       "CNY",
	}, nil
}

func (s *CourseService) CreateOrder(ctx context.Context, input CreateCourseOrderInput) (*CourseOrderView, error) {
	if input.ApplyMemberDiscount && !input.UserKnown {
		return nil, ErrCourseMemberDiscountRequiresLogin
	}

	pricing, err := s.calculatePricing(input.CourseCheckoutInput)
	if err != nil {
		return nil, err
	}

	finalAmount := pricing.FinalAmount
	if input.ApplyMemberDiscount {
		if !pricing.MemberPreview.Applied {
			return nil, ErrCourseMemberDiscountRequiresLogin
		}
		finalAmount = pricing.AmountIfMember
	}

	payChannel := input.PayChannel
	if payChannel == "" {
		payChannel = model.PayChannelWechatNative
	}
	if !model.ValidPayChannel(payChannel) {
		return nil, ErrPaymentChannelInvalid
	}

	now := time.Now()
	expiredAt := now.Add(time.Duration(s.membership.OrderExpireMinutes) * time.Minute)
	orderNo := generateCourseOrderNo(now)

	var userID *int64
	isAnonymous := !input.UserKnown
	var plainToken string
	var tokenHash *string
	if input.UserKnown {
		userID = &input.UserID
		isAnonymous = false
	} else {
		plain, hash, err := ordertoken.Generate()
		if err != nil {
			return nil, err
		}
		plainToken = plain
		tokenHash = &hash
	}

	productRefBytes, err := json.Marshal(CourseProductRef{
		CourseID: input.CourseID,
		SKUID:    input.SKUID,
	})
	if err != nil {
		return nil, err
	}
	productRef := string(productRefBytes)

	snapshotBytes, err := json.Marshal(courseDiscountSnapshot{
		Items:             pricing.Discounts,
		AmountIfNonMember: pricing.FinalAmount,
		PricedAt:          now.UTC().Format(time.RFC3339),
	})
	if err != nil {
		return nil, err
	}
	snapshot := string(snapshotBytes)

	var channelCode *string
	if strings.TrimSpace(input.ChannelCode) != "" {
		normalized := normalizeChannelCode(input.ChannelCode)
		channelCode = &normalized
	}

	description := fmt.Sprintf("SeeDAO 课程 %s", input.CourseID)
	var payURL string
	switch payChannel {
	case model.PayChannelWechatNative:
		nativeResult, err := s.wechat.CreateNativeOrder(ctx, wechatpay.NativeOrderInput{
			OrderNo:     orderNo,
			Description: description,
			AmountFen:   int64(finalAmount),
			ExpireAt:    expiredAt,
		})
		if err != nil {
			return nil, err
		}
		payURL = nativeResult.CodeURL
	case model.PayChannelWechatH5:
		h5Result, err := s.wechat.CreateH5Order(ctx, wechatpay.H5OrderInput{
			OrderNo:     orderNo,
			Description: description,
			AmountFen:   int64(finalAmount),
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
			Subject:   description,
			AmountFen: int64(finalAmount),
			ExpireAt:  expiredAt,
		})
		if err != nil {
			return nil, err
		}
		payURL = wapResult.PayURL
	case model.PayChannelAlipayPage:
		pageResult, err := s.alipay.CreatePageOrder(ctx, alipay.PageOrderInput{
			OrderNo:   orderNo,
			Subject:   description,
			AmountFen: int64(finalAmount),
			ExpireAt:  expiredAt,
		})
		if err != nil {
			return nil, err
		}
		payURL = pageResult.PayURL
	}

	order := model.PaymentOrder{
		OrderNo:          orderNo,
		UserID:           userID,
		ProductType:      model.ProductTypeCoursePurchase,
		ProductDays:      0,
		ProductRef:       &productRef,
		OriginalAmount:   pricing.OriginalAmount,
		DiscountAmount:   pricing.OriginalAmount - finalAmount,
		DiscountSnapshot: &snapshot,
		ChannelCode:      channelCode,
		OrderTokenHash:   tokenHash,
		IsAnonymous:      isAnonymous,
		Amount:           finalAmount,
		Currency:         "CNY",
		Status:           model.PaymentOrderStatusPending,
		PayChannel:       payChannel,
		CodeURL:          &payURL,
		ExpiredAt:        expiredAt,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	if err := s.db.Create(&order).Error; err != nil {
		_ = s.closeProviderOrder(ctx, payChannel, orderNo)
		return nil, err
	}

	view := toCourseOrderView(&order, pricing.Discounts, "")
	view.OrderToken = plainToken
	return view, nil
}

func (s *CourseService) GetOrder(userID int64, userKnown bool, orderNo, token string) (*CourseOrderView, error) {
	order, err := s.findAccessibleOrder(userID, userKnown, orderNo, token)
	if err != nil {
		return nil, err
	}
	return s.buildOrderView(order)
}

func (s *CourseService) CloseOrder(ctx context.Context, userID int64, userKnown bool, orderNo, token string) (*CourseOrderView, error) {
	order, err := s.findAccessibleOrder(userID, userKnown, orderNo, token)
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
	if err := s.db.Save(order).Error; err != nil {
		return nil, err
	}

	return s.buildOrderView(order)
}

func (s *CourseService) ClaimOrder(userID int64, orderNo, token string) (*CourseOrderView, error) {
	if strings.TrimSpace(token) == "" {
		return nil, ErrCourseOrderTokenRequired
	}

	var order model.PaymentOrder
	if err := s.db.Where("order_no = ? AND product_type = ?", orderNo, model.ProductTypeCoursePurchase).
		First(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPaymentOrderNotFound
		}
		return nil, err
	}

	if order.OrderTokenHash == nil || !ordertoken.Verify(token, *order.OrderTokenHash) {
		return nil, ErrCourseOrderTokenInvalid
	}
	if order.UserID != nil {
		if *order.UserID == userID {
			return s.buildOrderView(&order)
		}
		return nil, ErrCourseOrderAlreadyClaimed
	}

	now := time.Now()
	order.UserID = &userID
	order.IsAnonymous = false
	order.UpdatedAt = now

	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&order).Error; err != nil {
			return err
		}

		var enrollment model.CourseEnrollment
		err := tx.Where("order_id = ?", order.ID).First(&enrollment).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		if err != nil {
			return err
		}

		enrollment.UserID = &userID
		enrollment.UpdatedAt = now
		return tx.Save(&enrollment).Error
	})
	if err != nil {
		return nil, err
	}

	return s.buildOrderView(&order)
}

func (s *CourseService) SimulatePayDev(ctx context.Context, userID int64, userKnown bool, orderNo, token string, completeFn func(context.Context, string, string, int64, string, time.Time) error) (*CourseOrderView, error) {
	if !s.devMode {
		return nil, ErrInvalidParam
	}

	order, err := s.findAccessibleOrder(userID, userKnown, orderNo, token)
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
	if err := completeFn(ctx, orderNo, txID, int64(order.Amount), raw, time.Now()); err != nil {
		return nil, err
	}

	var updated model.PaymentOrder
	if err := s.db.Where("order_no = ?", orderNo).First(&updated).Error; err != nil {
		return nil, err
	}
	return s.buildOrderView(&updated)
}

func (s *CourseService) FulfillPaidOrder(tx *gorm.DB, order *model.PaymentOrder) error {
	if order.ProductType != model.ProductTypeCoursePurchase {
		return nil
	}

	var existing model.CourseEnrollment
	err := tx.Where("order_id = ?", order.ID).First(&existing).Error
	if err == nil {
		return nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	ref, err := parseCourseProductRef(order.ProductRef)
	if err != nil {
		return err
	}

	now := time.Now()
	enrollment := model.CourseEnrollment{
		OrderID:   order.ID,
		UserID:    order.UserID,
		CourseID:  ref.CourseID,
		SKUID:     ref.SKUID,
		Status:    model.CourseEnrollmentActive,
		CreatedAt: now,
		UpdatedAt: now,
	}
	return tx.Create(&enrollment).Error
}

func (s *CourseService) calculatePricing(input CourseCheckoutInput) (*coursePricingResult, error) {
	courseID := strings.TrimSpace(input.CourseID)
	skuID := strings.TrimSpace(input.SKUID)
	if courseID == "" || skuID == "" {
		return nil, ErrInvalidParam
	}

	var course model.Course
	if err := s.db.Where("id = ?", courseID).First(&course).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCourseNotFound
		}
		return nil, err
	}
	if course.Status != model.CourseStatusOnSale {
		return nil, ErrCourseOffSale
	}

	var sku model.CourseSKU
	if err := s.db.Where("course_id = ? AND sku_id = ?", courseID, skuID).First(&sku).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCourseSKUNotFound
		}
		return nil, err
	}
	if sku.Status != model.CourseStatusOnSale {
		return nil, ErrCourseOffSale
	}

	var rules []model.CourseDiscountRule
	if err := s.db.Where("course_id = ? AND sku_id = ? AND status = ?", courseID, skuID, model.CourseStatusOnSale).
		Find(&rules).Error; err != nil {
		return nil, err
	}

	now := time.Now()
	channelCode := normalizeChannelCode(input.ChannelCode)
	isMember := input.UserKnown && s.isActiveMember(input.UserID, now)

	base := sku.BasePriceFen
	var lines []CourseDiscountLine
	totalNonMember := 0
	memberAmount := 0

	for _, rule := range rules {
		switch rule.RuleType {
		case model.CourseDiscountEarlyBird:
			if !ruleTimeValid(&rule, now) {
				continue
			}
			amount := calcRuleDiscount(&rule, base)
			if amount <= 0 {
				continue
			}
			lines = append(lines, CourseDiscountLine{
				Type:    model.CourseDiscountEarlyBird,
				Label:   "早鸟价",
				Amount:  amount,
				Applied: true,
			})
			totalNonMember += amount
		case model.CourseDiscountChannel:
			if channelCode == "" || rule.ChannelCode == nil || normalizeChannelCode(*rule.ChannelCode) != channelCode {
				continue
			}
			if !ruleTimeValid(&rule, now) {
				continue
			}
			amount := calcRuleDiscount(&rule, base)
			if amount <= 0 {
				continue
			}
			lines = append(lines, CourseDiscountLine{
				Type:    model.CourseDiscountChannel,
				Label:   fmt.Sprintf("渠道 %s", channelCode),
				Amount:  amount,
				Applied: true,
			})
			totalNonMember += amount
		case model.CourseDiscountMember:
			amount := calcRuleDiscount(&rule, base)
			if amount <= 0 {
				continue
			}
			memberAmount = amount
			if isMember && input.ApplyMemberDiscount {
				lines = append(lines, CourseDiscountLine{
					Type:    model.CourseDiscountMember,
					Label:   "会员价",
					Amount:  amount,
					Applied: true,
				})
			}
		}
	}

	finalNonMember := clampPaymentAmount(base - totalNonMember)
	amountIfMember := clampPaymentAmount(base - totalNonMember - memberAmount)

	memberPreview := CourseMemberDiscountPreview{
		Amount:        memberAmount,
		Applied:       isMember && input.ApplyMemberDiscount && memberAmount > 0,
		RequiresLogin: !input.UserKnown,
	}
	if memberAmount > 0 {
		yuan := fmt.Sprintf("%.2f", float64(amountIfMember)/100)
		if input.UserKnown && isMember {
			memberPreview.Hint = fmt.Sprintf("会员价 %s 元", yuan)
		} else {
			memberPreview.Hint = fmt.Sprintf("会员价 %s 元，登录后可享", yuan)
		}
	}

	return &coursePricingResult{
		OriginalAmount: base,
		DiscountAmount: base - finalNonMember,
		FinalAmount:    finalNonMember,
		AmountIfMember: amountIfMember,
		Discounts:      lines,
		MemberPreview:  memberPreview,
	}, nil
}

func (s *CourseService) isActiveMember(userID int64, now time.Time) bool {
	var user model.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return false
	}
	return user.MembershipExpiresAt != nil && user.MembershipExpiresAt.After(now)
}

func (s *CourseService) findAccessibleOrder(userID int64, userKnown bool, orderNo, token string) (*model.PaymentOrder, error) {
	var order model.PaymentOrder
	if err := s.db.Where("order_no = ? AND product_type = ?", orderNo, model.ProductTypeCoursePurchase).
		First(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPaymentOrderNotFound
		}
		return nil, err
	}

	if userKnown && order.UserID != nil {
		if *order.UserID != userID {
			return nil, ErrPaymentOrderForbidden
		}
		return &order, nil
	}

	if order.OrderTokenHash == nil {
		return nil, ErrCourseOrderTokenRequired
	}
	if !ordertoken.Verify(token, *order.OrderTokenHash) {
		return nil, ErrCourseOrderTokenInvalid
	}
	return &order, nil
}

func (s *CourseService) buildOrderView(order *model.PaymentOrder) (*CourseOrderView, error) {
	if _, err := parseCourseProductRef(order.ProductRef); err != nil {
		return nil, err
	}

	var discounts []CourseDiscountLine
	if order.DiscountSnapshot != nil {
		var snapshot courseDiscountSnapshot
		if err := json.Unmarshal([]byte(*order.DiscountSnapshot), &snapshot); err == nil {
			discounts = snapshot.Items
		}
	}

	enrollmentStatus := ""
	var enrollment model.CourseEnrollment
	if err := s.db.Where("order_id = ?", order.ID).First(&enrollment).Error; err == nil {
		enrollmentStatus = enrollment.Status
	}

	return toCourseOrderView(order, discounts, enrollmentStatus), nil
}

func parseCourseProductRef(raw *string) (*CourseProductRef, error) {
	if raw == nil || strings.TrimSpace(*raw) == "" {
		return nil, ErrInvalidParam
	}
	var ref CourseProductRef
	if err := json.Unmarshal([]byte(*raw), &ref); err != nil {
		return nil, err
	}
	if ref.CourseID == "" || ref.SKUID == "" {
		return nil, ErrInvalidParam
	}
	return &ref, nil
}

func toCourseOrderView(order *model.PaymentOrder, discounts []CourseDiscountLine, enrollmentStatus string) *CourseOrderView {
	ref, _ := parseCourseProductRef(order.ProductRef)
	courseID, skuID := "", ""
	if ref != nil {
		courseID = ref.CourseID
		skuID = ref.SKUID
	}

	view := &CourseOrderView{
		OrderNo:          order.OrderNo,
		ProductType:      order.ProductType,
		CourseID:         courseID,
		SKUID:            skuID,
		PayChannel:       order.PayChannel,
		OriginalAmount:   order.OriginalAmount,
		DiscountAmount:   order.DiscountAmount,
		Discounts:        discounts,
		Amount:           order.Amount,
		Currency:         order.Currency,
		Status:           order.Status,
		ExpiredAt:        order.ExpiredAt.UTC().Format(time.RFC3339),
		EnrollmentStatus: enrollmentStatus,
	}
	if order.CodeURL != nil && order.Status == model.PaymentOrderStatusPending {
		switch order.PayChannel {
		case model.PayChannelWechatH5:
			view.H5URL = *order.CodeURL
		case model.PayChannelAlipayH5, model.PayChannelAlipayPage:
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
	return view
}

func ruleTimeValid(rule *model.CourseDiscountRule, now time.Time) bool {
	if rule.ValidFrom != nil && now.Before(*rule.ValidFrom) {
		return false
	}
	if rule.ValidUntil != nil && now.After(*rule.ValidUntil) {
		return false
	}
	return true
}

func calcRuleDiscount(rule *model.CourseDiscountRule, base int) int {
	switch rule.DiscountType {
	case model.CouponDiscountFixed:
		return rule.DiscountValue
	case model.CouponDiscountPercent:
		if rule.DiscountValue <= 0 || rule.DiscountValue >= 100 {
			return 0
		}
		return base * rule.DiscountValue / 100
	default:
		return 0
	}
}

func clampPaymentAmount(amount int) int {
	if amount < model.MinPaymentAmountFen {
		return model.MinPaymentAmountFen
	}
	return amount
}

func normalizeChannelCode(code string) string {
	return strings.ToUpper(strings.TrimSpace(code))
}

func (s *CourseService) closeProviderOrder(ctx context.Context, payChannel, orderNo string) error {
	switch payChannel {
	case model.PayChannelAlipayH5, model.PayChannelAlipayPage:
		return s.alipay.CloseOrder(ctx, orderNo)
	default:
		return s.wechat.CloseOrder(ctx, orderNo)
	}
}

func generateCourseOrderNo(now time.Time) string {
	buf := make([]byte, 4)
	_, _ = rand.Read(buf)
	return fmt.Sprintf("C%s%s", now.Format("20060102150405"), hex.EncodeToString(buf))
}
