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
