package service

import (
	"errors"

	"gorm.io/gorm"

	"github.com/seedao/seedao-api-server/internal/model"
	"github.com/seedao/seedao-api-server/internal/pkg/phone"
)

type ProfileService struct {
	db *gorm.DB
}

func NewProfileService(db *gorm.DB) *ProfileService {
	return &ProfileService{db: db}
}

type ProfileResult struct {
	ID             int64  `json:"id"`
	SNSName        string `json:"sns_name"`
	Phone          string `json:"phone"`
	RegisterSource string `json:"register_source"`
	Avatar         string `json:"avatar"`
	Bio            string `json:"bio"`
	CreatedAt      string `json:"created_at"`
}

func (s *ProfileService) GetByUserID(userID int64) (*ProfileResult, error) {
	var user model.User
	err := s.db.Preload("Phone").Preload("SNSName").First(&user, userID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	profile := &ProfileResult{
		ID:             user.ID,
		RegisterSource: user.RegisterSource,
		Avatar:         user.Avatar,
		Bio:            user.Bio,
		CreatedAt:      user.CreatedAt.UTC().Format("2006-01-02T15:04:05Z"),
	}

	if user.Phone != nil {
		profile.Phone = phone.Mask(user.Phone.Phone)
	}
	if user.SNSName != nil {
		profile.SNSName = user.SNSName.Name
	}

	return profile, nil
}
