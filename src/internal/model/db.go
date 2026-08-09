package model

import (
	"fmt"
	"sync/atomic"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var sqliteMemDBSeq atomic.Uint64

func NewDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, fmt.Errorf("connect database: %w", err)
	}
	return db, nil
}

func NewSQLiteDB(dsn string) (*gorm.DB, error) {
	if dsn == "" || dsn == ":memory:" {
		// Unique name per call so tests stay isolated; cache=shared + MaxOpenConns(1)
		// lets concurrent queries in one process see the same schema/data.
		dsn = fmt.Sprintf("file:memdb-%d?mode=memory&cache=shared", sqliteMemDBSeq.Add(1))
	}
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, fmt.Errorf("connect sqlite: %w", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("sqlite sql db: %w", err)
	}
	sqlDB.SetMaxOpenConns(1)
	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&UserPhone{},
		&SNSName{},
		&SNSChainRegistry{},
		&SMSSendLog{},
		&PaymentOrder{},
		&PaymentRefund{},
		&Coupon{},
		&CouponRedemption{},
		&Course{},
		&CourseSKU{},
		&CourseDiscountRule{},
		&CourseEnrollment{},
		&Proposal{},
		&ProposalVote{},
		&ProposalComment{},
		&UserOAuthBinding{},
		&UserProfile{},
		&UserAsset{},
		&UserRole{},
		&UserSeed{},
		&UserSBT{},
		&RewardApplication{},
		&OSSObject{},
	)
}
