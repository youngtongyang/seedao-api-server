package service

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/seedao/seedao-api-server/internal/model"
)

var (
	ErrCouponInvalid         = errors.New("coupon invalid")
	ErrCouponNotApplicable   = errors.New("coupon not applicable")
	ErrCouponUsageExceeded   = errors.New("coupon usage exceeded")
	ErrCouponAmountTooLow    = errors.New("coupon amount too low")
)

type CouponService struct {
	db *gorm.DB
}

func NewCouponService(db *gorm.DB) *CouponService {
	return &CouponService{db: db}
}

type CouponQuoteInput struct {
	UserID      int64
	ProductType string
	Days        int
	CouponCode  string
}

type CouponQuote struct {
	CouponCode     string `json:"coupon_code"`
	OriginalAmount int    `json:"original_amount"`
	DiscountAmount int    `json:"discount_amount"`
	FinalAmount    int    `json:"final_amount"`
	Currency       string `json:"currency"`
}

func (s *CouponService) Preview(input CouponQuoteInput, originalAmount int64) (*CouponQuote, error) {
	if strings.TrimSpace(input.CouponCode) == "" {
		return nil, ErrInvalidParam
	}

	discount, coupon, err := s.resolveDiscount(input, originalAmount, nil)
	if err != nil {
		return nil, err
	}

	return &CouponQuote{
		CouponCode:     coupon.Code,
		OriginalAmount: int(originalAmount),
		DiscountAmount: int(discount),
		FinalAmount:    int(originalAmount - discount),
		Currency:       "CNY",
	}, nil
}

func (s *CouponService) Reserve(tx *gorm.DB, input CouponQuoteInput, orderNo string, orderID int64, originalAmount int64) (int64, string, error) {
	discount, coupon, err := s.resolveDiscount(input, originalAmount, tx)
	if err != nil {
		return 0, "", err
	}

	now := time.Now()
	redemption := model.CouponRedemption{
		CouponID:       coupon.ID,
		UserID:         input.UserID,
		OrderID:        &orderID,
		OrderNo:        orderNo,
		DiscountAmount: int(discount),
		Status:         model.CouponRedemptionReserved,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	if err := tx.Create(&redemption).Error; err != nil {
		return 0, "", err
	}

	return discount, coupon.Code, nil
}

func (s *CouponService) ReleaseByOrderID(tx *gorm.DB, orderID int64) error {
	var redemption model.CouponRedemption
	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("order_id = ? AND status = ?", orderID, model.CouponRedemptionReserved).
		First(&redemption).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if err != nil {
		return err
	}

	redemption.Status = model.CouponRedemptionReleased
	redemption.UpdatedAt = time.Now()
	return tx.Save(&redemption).Error
}

func (s *CouponService) ConsumeByOrderID(tx *gorm.DB, orderID int64) error {
	var redemption model.CouponRedemption
	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("order_id = ? AND status = ?", orderID, model.CouponRedemptionReserved).
		First(&redemption).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if err != nil {
		return err
	}

	var coupon model.Coupon
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&coupon, redemption.CouponID).Error; err != nil {
		return err
	}

	now := time.Now()
	redemption.Status = model.CouponRedemptionConsumed
	redemption.UpdatedAt = now
	if err := tx.Save(&redemption).Error; err != nil {
		return err
	}

	coupon.UsedCount++
	coupon.UpdatedAt = now
	return tx.Save(&coupon).Error
}

func (s *CouponService) resolveDiscount(input CouponQuoteInput, originalAmount int64, tx *gorm.DB) (int64, *model.Coupon, error) {
	db := s.db
	if tx != nil {
		db = tx
	}

	code := normalizeCouponCode(input.CouponCode)
	if code == "" {
		return 0, nil, ErrCouponInvalid
	}

	var coupon model.Coupon
	err := db.Where("code = ?", code).First(&coupon).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, nil, ErrCouponInvalid
	}
	if err != nil {
		return 0, nil, err
	}

	now := time.Now()
	if coupon.Status != model.CouponStatusActive {
		return 0, nil, ErrCouponInvalid
	}
	if now.Before(coupon.ValidFrom) || now.After(coupon.ValidUntil) {
		return 0, nil, ErrCouponInvalid
	}

	productType := input.ProductType
	if productType == "" {
		productType = model.ProductTypeMembershipRenew
	}
	if coupon.ProductType != nil && *coupon.ProductType != productType {
		return 0, nil, ErrCouponNotApplicable
	}

	if originalAmount < int64(coupon.MinAmountFen) {
		return 0, nil, ErrCouponNotApplicable
	}

	if err := s.checkUsageLimits(db, &coupon, input.UserID); err != nil {
		return 0, nil, err
	}

	discount, err := calcCouponDiscount(&coupon, originalAmount)
	if err != nil {
		return 0, nil, err
	}

	return discount, &coupon, nil
}

func (s *CouponService) checkUsageLimits(db *gorm.DB, coupon *model.Coupon, userID int64) error {
	var activeTotal int64
	if err := db.Model(&model.CouponRedemption{}).
		Where("coupon_id = ? AND status IN ?", coupon.ID, []string{
			model.CouponRedemptionReserved,
			model.CouponRedemptionConsumed,
		}).
		Count(&activeTotal).Error; err != nil {
		return err
	}
	if coupon.MaxTotalUses != nil && activeTotal >= int64(*coupon.MaxTotalUses) {
		return ErrCouponUsageExceeded
	}

	var userUses int64
	if err := db.Model(&model.CouponRedemption{}).
		Where("coupon_id = ? AND user_id = ? AND status IN ?", coupon.ID, userID, []string{
			model.CouponRedemptionReserved,
			model.CouponRedemptionConsumed,
		}).
		Count(&userUses).Error; err != nil {
		return err
	}
	if int(userUses) >= coupon.MaxUsesPerUser {
		return ErrCouponUsageExceeded
	}

	return nil
}

func calcCouponDiscount(coupon *model.Coupon, originalAmount int64) (int64, error) {
	var discount int64
	switch coupon.DiscountType {
	case model.CouponDiscountFixed:
		discount = int64(coupon.DiscountValue)
	case model.CouponDiscountPercent:
		if coupon.DiscountValue <= 0 || coupon.DiscountValue >= 100 {
			return 0, ErrCouponInvalid
		}
		discount = originalAmount * int64(coupon.DiscountValue) / 100
	default:
		return 0, ErrCouponInvalid
	}

	if discount <= 0 {
		return 0, ErrCouponNotApplicable
	}
	if discount >= originalAmount {
		return 0, ErrCouponAmountTooLow
	}
	if originalAmount-discount < model.MinPaymentAmountFen {
		return 0, ErrCouponAmountTooLow
	}

	return discount, nil
}

func normalizeCouponCode(code string) string {
	return strings.ToUpper(strings.TrimSpace(code))
}
