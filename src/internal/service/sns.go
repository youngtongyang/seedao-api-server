package service

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/seedao/seedao-api-server/internal/model"
	"github.com/seedao/seedao-api-server/internal/pkg/redisstore"
	"github.com/seedao/seedao-api-server/internal/pkg/snsutil"
)

const snsCheckCacheTTL = 60 * time.Second

type SNSService struct {
	db    *gorm.DB
	redis *redisstore.Store
}

func NewSNSService(db *gorm.DB, redis *redisstore.Store) *SNSService {
	return &SNSService{db: db, redis: redis}
}

type SNSCheckResult struct {
	SNSName   string `json:"sns_name"`
	Available bool   `json:"available"`
	Reason    string `json:"reason,omitempty"`
}

func (s *SNSService) Check(ctx context.Context, snsName string) (*SNSCheckResult, error) {
	normalized := snsutil.Normalize(snsName)
	if err := snsutil.ValidateFormat(snsName); err != nil {
		return &SNSCheckResult{
			SNSName:   normalized,
			Available: false,
			Reason:    snsutil.ReasonInvalidFormat,
		}, nil
	}

	if cached, err := s.redis.GetSNSCheckCache(ctx, normalized); err == nil && cached != "" {
		var result SNSCheckResult
		if json.Unmarshal([]byte(cached), &result) == nil {
			return &result, nil
		}
	}

	result, err := s.checkDB(normalized)
	if err != nil {
		return nil, err
	}

	if data, err := json.Marshal(result); err == nil {
		_ = s.redis.SetSNSCheckCache(ctx, normalized, string(data), snsCheckCacheTTL)
	}
	return result, nil
}

func (s *SNSService) checkDB(name string) (*SNSCheckResult, error) {
	var snsCount int64
	if err := s.db.Model(&model.SNSName{}).Where("LOWER(name) = ?", name).Count(&snsCount).Error; err != nil {
		return nil, err
	}
	if snsCount > 0 {
		return &SNSCheckResult{
			SNSName:   name,
			Available: false,
			Reason:    snsutil.ReasonTakenInApp,
		}, nil
	}

	registryNames := snsutil.RegistryLookupNames(name)
	var registry model.SNSChainRegistry
	err := s.db.Where("LOWER(sns_name) IN ? AND claimed = ?", registryNames, false).First(&registry).Error
	if err == nil {
		return &SNSCheckResult{
			SNSName:   name,
			Available: false,
			Reason:    snsutil.ReasonTakenInRegistry,
		}, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &SNSCheckResult{SNSName: name, Available: true}, nil
}

func (s *SNSService) IsAvailable(name string) (bool, error) {
	result, err := s.checkDB(snsutil.Normalize(name))
	if err != nil {
		return false, err
	}
	return result.Available, nil
}
