package service

import (
	"gorm.io/gorm"

	"github.com/seedao/seedao-api-server/internal/model"
)

type StatsService struct {
	db *gorm.DB
}

func NewStatsService(db *gorm.DB) *StatsService {
	return &StatsService{db: db}
}

type UserCountResult struct {
	TotalUserCount int64 `json:"total_user_count"`
}

func (s *StatsService) GetUserCount() (*UserCountResult, error) {
	var legacySNSCount int64
	if err := s.db.Model(&model.SNSChainRegistry{}).Count(&legacySNSCount).Error; err != nil {
		return nil, err
	}

	var newSNSUserCount int64
	if err := s.db.Model(&model.User{}).
		Where("register_source = ?", model.RegisterSourceNewSNS).
		Count(&newSNSUserCount).Error; err != nil {
		return nil, err
	}

	return &UserCountResult{
		TotalUserCount: legacySNSCount + newSNSUserCount,
	}, nil
}
