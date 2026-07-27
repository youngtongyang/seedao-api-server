package model

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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
	if dsn == "" {
		dsn = ":memory:"
	}
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		return nil, fmt.Errorf("connect sqlite: %w", err)
	}
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
	)
}
