package model

import "time"

// UserProfile mirrors 1.0 profile fields keyed by SNS (from sns_chain_registry).
type UserProfile struct {
	SNSName         string    `gorm:"column:sns_name;size:64;primaryKey"`
	Bio             string    `gorm:"column:bio;type:text;not null;default:''"`
	Email           string    `gorm:"column:email;size:255;not null;default:''"`
	TwitterProfile  string    `gorm:"column:twitter_profile;size:512;not null;default:''"`
	Mirror          string    `gorm:"column:mirror;size:512;not null;default:''"`
	SCRAmount       *string   `gorm:"column:scr_amount;type:numeric(78,30)"`
	CreatedAt       time.Time `gorm:"column:created_at;not null"`
	UpdatedAt       time.Time `gorm:"column:updated_at;not null"`
}

func (UserProfile) TableName() string { return "user_profiles" }

// UserAsset mirrors 1.0 assets[] (WANG / SEE, etc.).
type UserAsset struct {
	ID                int64     `gorm:"column:id;primaryKey;autoIncrement"`
	SNSName           string    `gorm:"column:sns_name;size:64;not null;uniqueIndex:uq_user_assets_sns_asset"`
	AssetName         string    `gorm:"column:asset_name;size:32;not null;uniqueIndex:uq_user_assets_sns_asset"`
	DealtAmount       string    `gorm:"column:dealt_amount;type:numeric(78,30);not null;default:0"`
	ProcessingAmount  string    `gorm:"column:processing_amount;type:numeric(78,30);not null;default:0"`
	Claimed           bool      `gorm:"column:claimed;not null;default:false"`
	SourceCreateTS    *int64    `gorm:"column:source_create_ts"`
	CreatedAt         time.Time `gorm:"column:created_at;not null"`
	UpdatedAt         time.Time `gorm:"column:updated_at;not null"`
}

func (UserAsset) TableName() string { return "user_assets" }

// UserRole mirrors 1.0 sp.roles[].
type UserRole struct {
	SNSName   string    `gorm:"column:sns_name;size:64;primaryKey"`
	Role      string    `gorm:"column:role;size:64;primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
}

func (UserRole) TableName() string { return "user_roles" }

// UserSeed mirrors 1.0 sp.seed[].
type UserSeed struct {
	ID            int64     `gorm:"column:id;primaryKey;autoIncrement"`
	SNSName       string    `gorm:"column:sns_name;size:64;not null;uniqueIndex:uq_user_seeds_key"`
	TokenID       string    `gorm:"column:token_id;size:64;not null;uniqueIndex:uq_user_seeds_key"`
	ContractAddr  string    `gorm:"column:contract_addr;size:42;not null;uniqueIndex:uq_user_seeds_key"`
	ContractType  string    `gorm:"column:contract_type;size:16;not null;default:''"`
	TokenAmount   string    `gorm:"column:token_amount;size:32;not null;default:''"`
	ImageURISrc   string    `gorm:"column:image_uri_src;type:text;not null;default:''"`
	ImageURIOSS   string    `gorm:"column:image_uri_oss;size:1024;not null;default:''"`
	CreatedAt     time.Time `gorm:"column:created_at;not null"`
	UpdatedAt     time.Time `gorm:"column:updated_at;not null"`
}

func (UserSeed) TableName() string { return "user_seeds" }

// UserSBT mirrors 1.0 sp.sbt[].
type UserSBT struct {
	ID              int64     `gorm:"column:id;primaryKey;autoIncrement"`
	SNSName         string    `gorm:"column:sns_name;size:64;not null;uniqueIndex:uq_user_sbts_key"`
	TokenID         string    `gorm:"column:token_id;size:64;not null;uniqueIndex:uq_user_sbts_key"`
	ContractAddr    string    `gorm:"column:contract_addr;size:42;not null;uniqueIndex:uq_user_sbts_key"`
	ContractType    string    `gorm:"column:contract_type;size:16;not null;default:''"`
	TokenAmount     string    `gorm:"column:token_amount;size:32;not null;default:''"`
	CollectionName  string    `gorm:"column:collection_name;size:255;not null;default:''"`
	DisplayName     string    `gorm:"column:display_name;size:512;not null;default:''"`
	Description     string    `gorm:"column:description;type:text;not null;default:''"`
	ImageURISrc     string    `gorm:"column:image_uri_src;type:text;not null;default:''"`
	ImageURIOSS     string    `gorm:"column:image_uri_oss;size:1024;not null;default:''"`
	CreatedAt       time.Time `gorm:"column:created_at;not null"`
	UpdatedAt       time.Time `gorm:"column:updated_at;not null"`
}

func (UserSBT) TableName() string { return "user_sbts" }

// RewardApplication mirrors 1.0 new_reward applications.
type RewardApplication struct {
	ApplicationID     int64     `gorm:"column:application_id;primaryKey"`
	SeasonName        string    `gorm:"column:season_name;size:32;not null;default:''"`
	EntityName        string    `gorm:"column:entity_name;size:512;not null;default:''"`
	EntityType        string    `gorm:"column:entity_type;size:64;not null;default:''"`
	AssetName         string    `gorm:"column:asset_name;size:32;not null;default:''"`
	Amount            *string   `gorm:"column:amount;type:numeric(78,30)"`
	BudgetSource      string    `gorm:"column:budget_source;size:512;not null;default:''"`
	Status            string    `gorm:"column:status;size:32;not null;default:''"`
	TargetUserWallet  string    `gorm:"column:target_user_wallet;size:42;not null;default:''"`
	TargetSNSName     *string   `gorm:"column:target_sns_name;size:64"`
	ApplicantWallet   string    `gorm:"column:applicant_wallet;size:42;not null;default:''"`
	ApplicantSNSName  *string   `gorm:"column:applicant_sns_name;size:64"`
	SourceCreateTS    *int64    `gorm:"column:source_create_ts"`
	SourceUpdateTS    *int64    `gorm:"column:source_update_ts"`
	CreatedAt         time.Time `gorm:"column:created_at;not null"`
	UpdatedAt         time.Time `gorm:"column:updated_at;not null"`
}

func (RewardApplication) TableName() string { return "reward_applications" }

// OSSObject caches uploaded image URLs keyed by source URL hash.
type OSSObject struct {
	SrcURLHash  string    `gorm:"column:src_url_hash;size:64;primaryKey"`
	SrcURL      string    `gorm:"column:src_url;type:text;not null"`
	OSSURL      string    `gorm:"column:oss_url;size:1024;not null"`
	ContentType string    `gorm:"column:content_type;size:128;not null;default:''"`
	CreatedAt   time.Time `gorm:"column:created_at;not null"`
}

func (OSSObject) TableName() string { return "oss_objects" }
