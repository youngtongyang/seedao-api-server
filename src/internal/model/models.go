package model

import (
	"time"

	"gorm.io/gorm"
)

const (
	RegisterSourceNewSNS      = "new_sns"
	RegisterSourceExistingSNS = "existing_sns"

	UserStatusActive    = "active"
	UserStatusSuspended = "suspended"

	MintStatusReserved = "reserved"
	MintStatusOnChain  = "on_chain"

	DefaultMembershipDays = 365

	ProductTypeMembershipRenew = "membership_renew"
	ProductTypeCoursePurchase  = "course_purchase"

	PayChannelWechatNative = "wechat_native"
	PayChannelWechatH5     = "wechat_h5"
	PayChannelAlipayH5     = "alipay_h5"
	PayChannelAlipayPage   = "alipay_page"

	PaymentOrderStatusPending = "pending"
	PaymentOrderStatusPaid    = "paid"
	PaymentOrderStatusClosed  = "closed"

	PaymentRefundStatusPending = "pending"
	PaymentRefundStatusSuccess = "success"
	PaymentRefundStatusFailed  = "failed"

	CouponDiscountFixed   = "fixed"
	CouponDiscountPercent = "percent"

	CouponStatusActive   = "active"
	CouponStatusDisabled = "disabled"

	CouponRedemptionReserved = "reserved"
	CouponRedemptionConsumed = "consumed"
	CouponRedemptionReleased = "released"

	CourseStatusOnSale  = "on_sale"
	CourseStatusOffSale = "off_sale"

	CourseDiscountEarlyBird = "early_bird"
	CourseDiscountMember    = "member"
	CourseDiscountChannel   = "channel"

	CourseEnrollmentActive = "active"

	OAuthProviderFeishu = "feishu"

	OAuthBindSourceBotToken = "bot_token"
	OAuthBindSourceBotCode  = "bot_code"

	MinPaymentAmountFen = 1
)

func ValidPayChannel(payChannel string) bool {
	switch payChannel {
	case PayChannelWechatNative, PayChannelWechatH5, PayChannelAlipayH5, PayChannelAlipayPage:
		return true
	default:
		return false
	}
}

type User struct {
	ID                  int64          `gorm:"primaryKey"`
	RegisterSource      string         `gorm:"size:32;not null"`
	Status              string         `gorm:"size:16;not null;default:active"`
	Avatar              string         `gorm:"size:512"`
	Bio                 string         `gorm:"type:text"`
	MembershipExpiresAt *time.Time     `gorm:"index"`
	IsTest              bool           `gorm:"column:is_test;not null;default:false;index"`
	CreatedAt           time.Time      `gorm:"not null"`
	UpdatedAt           time.Time      `gorm:"not null"`
	DeletedAt           gorm.DeletedAt `gorm:"index"`

	Phone   *UserPhone `gorm:"foreignKey:UserID"`
	SNSName *SNSName   `gorm:"foreignKey:UserID"`
}

func (User) TableName() string { return "users" }

type UserPhone struct {
	ID         int64     `gorm:"primaryKey"`
	UserID     int64     `gorm:"uniqueIndex;not null"`
	Phone      string    `gorm:"size:20;uniqueIndex;not null"`
	PhoneHash  string    `gorm:"size:64;not null"`
	VerifiedAt time.Time `gorm:"not null"`
	CreatedAt  time.Time `gorm:"not null"`
}

func (UserPhone) TableName() string { return "user_phones" }

type SNSName struct {
	ID          int64      `gorm:"primaryKey"`
	Name        string     `gorm:"size:64;uniqueIndex;not null"`
	UserID      int64      `gorm:"uniqueIndex;not null"`
	MintStatus  string     `gorm:"size:16;not null"`
	Wallet      *string    `gorm:"size:42;uniqueIndex"`
	ReservedAt  *time.Time
	CreatedAt   time.Time  `gorm:"not null"`
}

func (SNSName) TableName() string { return "sns_names" }

type SNSChainRegistry struct {
	ID        int64     `gorm:"primaryKey"`
	Wallet    string    `gorm:"size:42;uniqueIndex;not null"`
	SNSName   string    `gorm:"size:64;uniqueIndex;not null;column:sns_name"`
	Claimed   bool      `gorm:"not null;default:false"`
	SyncedAt  time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (SNSChainRegistry) TableName() string { return "sns_chain_registry" }

type SMSSendLog struct {
	ID        int64     `gorm:"primaryKey"`
	PhoneHash string    `gorm:"size:64;not null;index"`
	IP        string    `gorm:"size:45"`
	Status    string    `gorm:"size:16;not null"`
	CreatedAt time.Time `gorm:"not null"`
}

func (SMSSendLog) TableName() string { return "sms_send_logs" }

type PaymentOrder struct {
	ID               int64      `gorm:"primaryKey"`
	OrderNo          string     `gorm:"size:64;uniqueIndex;not null"`
	UserID           *int64     `gorm:"index"`
	ProductType      string     `gorm:"size:32;not null;index"`
	ProductDays      int        `gorm:"not null;default:365"`
	ProductRef       *string    `gorm:"type:text"`
	OriginalAmount   int        `gorm:"not null;default:0"`
	DiscountAmount   int        `gorm:"not null;default:0"`
	CouponCode       *string    `gorm:"size:64"`
	DiscountSnapshot *string    `gorm:"type:text"`
	ChannelCode      *string    `gorm:"size:64"`
	OrderTokenHash   *string    `gorm:"size:64;index"`
	IsAnonymous      bool       `gorm:"not null;default:false"`
	Amount           int        `gorm:"not null"`
	Currency         string     `gorm:"size:8;not null;default:CNY"`
	Status           string     `gorm:"size:16;not null;default:pending;index"`
	PayChannel       string     `gorm:"size:16;not null;default:wechat_native"`
	CodeURL          *string    `gorm:"type:text"`
	WxTransactionID  *string    `gorm:"size:64;index"`
	PaidAt           *time.Time
	ExpiredAt        time.Time `gorm:"not null"`
	ClosedAt         *time.Time
	NotifyPayload    *string   `gorm:"type:text"`
	CreatedAt        time.Time `gorm:"not null"`
	UpdatedAt        time.Time `gorm:"not null"`
}

func (PaymentOrder) TableName() string { return "payment_orders" }

type PaymentRefund struct {
	ID            int64     `gorm:"primaryKey"`
	OrderID       int64     `gorm:"index;not null"`
	RefundNo      string    `gorm:"size:64;uniqueIndex;not null"`
	WxRefundID    *string   `gorm:"size:64"`
	Amount        int       `gorm:"not null"`
	Reason        string    `gorm:"size:256"`
	Status        string    `gorm:"size:16;not null;default:pending"`
	NotifyPayload *string   `gorm:"type:text"`
	CreatedAt     time.Time `gorm:"not null"`
	UpdatedAt     time.Time `gorm:"not null"`
}

func (PaymentRefund) TableName() string { return "payment_refunds" }

type Coupon struct {
	ID             int64     `gorm:"primaryKey"`
	Code           string    `gorm:"size:64;uniqueIndex;not null"`
	DiscountType   string    `gorm:"size:16;not null;default:fixed"`
	DiscountValue  int       `gorm:"not null"`
	ProductType    *string   `gorm:"size:32"`
	MinAmountFen   int       `gorm:"not null;default:0"`
	MaxTotalUses   *int
	UsedCount      int       `gorm:"not null;default:0"`
	MaxUsesPerUser int       `gorm:"not null;default:1"`
	ValidFrom      time.Time `gorm:"not null"`
	ValidUntil     time.Time `gorm:"not null"`
	Status         string    `gorm:"size:16;not null;default:active"`
	CreatedAt      time.Time `gorm:"not null"`
	UpdatedAt      time.Time `gorm:"not null"`
}

func (Coupon) TableName() string { return "coupons" }

type CouponRedemption struct {
	ID             int64     `gorm:"primaryKey"`
	CouponID       int64     `gorm:"index;not null"`
	UserID         int64     `gorm:"index;not null"`
	OrderID        *int64    `gorm:"index"`
	OrderNo        string    `gorm:"size:64;not null;index"`
	DiscountAmount int       `gorm:"not null"`
	Status         string    `gorm:"size:16;not null;default:reserved;index"`
	CreatedAt      time.Time `gorm:"not null"`
	UpdatedAt      time.Time `gorm:"not null"`
}

func (CouponRedemption) TableName() string { return "coupon_redemptions" }

type Course struct {
	ID        string    `gorm:"primaryKey;size:64"`
	Title     string    `gorm:"size:256;not null"`
	Status    string    `gorm:"size:16;not null;default:on_sale"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (Course) TableName() string { return "courses" }

type CourseSKU struct {
	ID           int64     `gorm:"primaryKey"`
	CourseID     string    `gorm:"size:64;not null;index"`
	SKUID        string    `gorm:"column:sku_id;size:64;not null"`
	Name         string    `gorm:"size:128;not null"`
	BasePriceFen int       `gorm:"not null"`
	Status       string    `gorm:"size:16;not null;default:on_sale"`
	CreatedAt    time.Time `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"not null"`
}

func (CourseSKU) TableName() string { return "course_skus" }

type CourseDiscountRule struct {
	ID            int64      `gorm:"primaryKey"`
	CourseID      string     `gorm:"size:64;not null;index"`
	SKUID         string     `gorm:"column:sku_id;size:64;not null"`
	RuleType      string     `gorm:"size:16;not null"`
	DiscountType  string     `gorm:"size:16;not null;default:fixed"`
	DiscountValue int        `gorm:"not null"`
	ChannelCode   *string    `gorm:"size:64"`
	ValidFrom     *time.Time
	ValidUntil    *time.Time
	Status        string     `gorm:"size:16;not null;default:active"`
	CreatedAt     time.Time  `gorm:"not null"`
	UpdatedAt     time.Time  `gorm:"not null"`
}

func (CourseDiscountRule) TableName() string { return "course_discount_rules" }

type CourseEnrollment struct {
	ID        int64     `gorm:"primaryKey"`
	OrderID   int64     `gorm:"uniqueIndex;not null"`
	UserID    *int64    `gorm:"index"`
	CourseID  string    `gorm:"size:64;not null;index"`
	SKUID     string    `gorm:"column:sku_id;size:64;not null"`
	Status    string    `gorm:"size:16;not null;default:active"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (CourseEnrollment) TableName() string { return "course_enrollments" }

type UserOAuthBinding struct {
	ID            int64          `gorm:"primaryKey"`
	UserID        int64          `gorm:"not null;index"`
	Provider      string         `gorm:"size:32;not null"`
	ProviderAppID string         `gorm:"size:64;not null"`
	OpenID        string         `gorm:"size:128;not null"`
	UnionID       *string        `gorm:"size:128"`
	DisplayName   string         `gorm:"size:128;not null;default:''"`
	BindSource    string         `gorm:"size:32;not null"`
	BoundAt       time.Time      `gorm:"not null"`
	CreatedAt     time.Time      `gorm:"not null"`
	UpdatedAt     time.Time      `gorm:"not null"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func (UserOAuthBinding) TableName() string { return "user_oauth_bindings" }
