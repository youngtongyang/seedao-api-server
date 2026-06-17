package service

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"time"

	"gorm.io/gorm"

	"github.com/seedao/seedao-api-server/internal/model"
	"github.com/seedao/seedao-api-server/internal/pkg/aliyunsms"
	"github.com/seedao/seedao-api-server/internal/pkg/phone"
	"github.com/seedao/seedao-api-server/internal/pkg/redisstore"
)

const (
	smsCodeTTL      = 5 * time.Minute
	smsPhoneRateTTL = 60 * time.Second
	smsIPRateTTL    = time.Hour
	smsIPRateLimit  = 20
)

type SMSService struct {
	db     *gorm.DB
	redis  *redisstore.Store
	sender aliyunsms.Sender
}

func NewSMSService(db *gorm.DB, redis *redisstore.Store, sender aliyunsms.Sender) *SMSService {
	return &SMSService{db: db, redis: redis, sender: sender}
}

type SendSMSResult struct {
	RetryAfter int `json:"retry_after"`
}

func (s *SMSService) Send(ctx context.Context, phoneNum, ip string) (*SendSMSResult, error) {
	if err := phone.Validate(phoneNum); err != nil {
		return nil, ErrInvalidParam
	}

	phoneHash := phone.Hash(phoneNum)

	limited, err := s.redis.PhoneRateLimited(ctx, phoneHash)
	if err != nil {
		return nil, err
	}
	if limited {
		return nil, ErrTooManyRequests
	}

	count, err := s.redis.IncrIPRate(ctx, ip, smsIPRateTTL)
	if err != nil {
		return nil, err
	}
	if count > smsIPRateLimit {
		return nil, ErrTooManyRequests
	}

	code, err := generateSMSCode()
	if err != nil {
		return nil, err
	}

	if err := s.sender.Send(phoneNum, code); err != nil {
		_ = s.logSend(phoneHash, ip, "failed")
		return nil, err
	}

	if err := s.redis.SetSMSCode(ctx, phoneHash, code, smsCodeTTL); err != nil {
		return nil, err
	}
	if err := s.redis.SetPhoneRateLimit(ctx, phoneHash, smsPhoneRateTTL); err != nil {
		return nil, err
	}
	_ = s.logSend(phoneHash, ip, "success")

	return &SendSMSResult{RetryAfter: 60}, nil
}

func (s *SMSService) Verify(ctx context.Context, phoneNum, code string) error {
	if err := phone.Validate(phoneNum); err != nil {
		return ErrInvalidParam
	}

	phoneHash := phone.Hash(phoneNum)
	stored, err := s.redis.GetSMSCode(ctx, phoneHash)
	if err != nil {
		return ErrInvalidSMSCode
	}
	if stored != code {
		return ErrInvalidSMSCode
	}
	return s.redis.DeleteSMSCode(ctx, phoneHash)
}

func (s *SMSService) logSend(phoneHash, ip, status string) error {
	return s.db.Create(&model.SMSSendLog{
		PhoneHash: phoneHash,
		IP:        ip,
		Status:    status,
		CreatedAt: time.Now(),
	}).Error
}

func generateSMSCode() (string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return "", fmt.Errorf("generate code: %w", err)
	}
	return fmt.Sprintf("%06d", n.Int64()), nil
}

func phoneExists(db *gorm.DB, phoneNum string) (bool, error) {
	var count int64
	err := db.Model(&model.UserPhone{}).Where("phone = ?", phoneNum).Count(&count).Error
	return count > 0, err
}

func getUserByPhone(db *gorm.DB, phoneNum string) (*model.User, error) {
	var up model.UserPhone
	if err := db.Where("phone = ?", phoneNum).First(&up).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	var user model.User
	if err := db.Preload("SNSName").First(&user, up.UserID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
