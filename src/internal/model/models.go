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

	PayChannelWechatNative = "wechat_native"
	PayChannelWechatH5     = "wechat_h5"

	PaymentOrderStatusPending = "pending"
	PaymentOrderStatusPaid    = "paid"
	PaymentOrderStatusClosed  = "closed"

	PaymentRefundStatusPending = "pending"
	PaymentRefundStatusSuccess = "success"
	PaymentRefundStatusFailed  = "failed"
)

type User struct {
	ID                  int64          `gorm:"primaryKey"`
	RegisterSource      string         `gorm:"size:32;not null"`
	Status              string         `gorm:"size:16;not null;default:active"`
	Avatar              string         `gorm:"size:512"`
	Bio                 string         `gorm:"type:text"`
	MembershipExpiresAt *time.Time     `gorm:"index"`
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
	ID              int64      `gorm:"primaryKey"`
	OrderNo         string     `gorm:"size:64;uniqueIndex;not null"`
	UserID          int64      `gorm:"index;not null"`
	ProductType     string     `gorm:"size:32;not null"`
	ProductDays     int        `gorm:"not null;default:365"`
	Amount          int        `gorm:"not null"`
	Currency        string     `gorm:"size:8;not null;default:CNY"`
	Status          string     `gorm:"size:16;not null;default:pending;index"`
	PayChannel      string     `gorm:"size:16;not null;default:wechat_native"`
	CodeURL         *string    `gorm:"type:text"`
	WxTransactionID *string    `gorm:"size:64;index"`
	PaidAt          *time.Time
	ExpiredAt       time.Time  `gorm:"not null"`
	ClosedAt        *time.Time
	NotifyPayload   *string    `gorm:"type:text"`
	CreatedAt       time.Time  `gorm:"not null"`
	UpdatedAt       time.Time  `gorm:"not null"`
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
