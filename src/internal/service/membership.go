package service

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/seedao/seedao-api-server/internal/model"
)

func DefaultMembershipExpiry(from time.Time) time.Time {
	return from.AddDate(0, 0, model.DefaultMembershipDays)
}

func ExtendMembershipExpiry(current *time.Time, now time.Time, days int) time.Time {
	base := now
	if current != nil && current.After(now) {
		base = *current
	}
	return base.AddDate(0, 0, days)
}

func formatMembershipExpiry(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.UTC().Format("2006-01-02T15:04:05Z")
}

type MembershipService struct {
	db *gorm.DB
}

func NewMembershipService(db *gorm.DB) *MembershipService {
	return &MembershipService{db: db}
}

type RenewMembershipInput struct {
	Days int
}

type RenewMembershipResult struct {
	MembershipExpiresAt string `json:"membership_expires_at"`
	DaysAdded           int    `json:"days_added"`
}

func (s *MembershipService) Renew(userID int64, input RenewMembershipInput) (*RenewMembershipResult, error) {
	days := input.Days
	if days <= 0 {
		days = model.DefaultMembershipDays
	}

	var user model.User
	err := s.db.First(&user, userID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	now := time.Now()
	expiresAt := ExtendMembershipExpiry(user.MembershipExpiresAt, now, days)
	user.MembershipExpiresAt = &expiresAt
	user.UpdatedAt = now

	if err := s.db.Save(&user).Error; err != nil {
		return nil, err
	}

	return &RenewMembershipResult{
		MembershipExpiresAt: formatMembershipExpiry(&expiresAt),
		DaysAdded:           days,
	}, nil
}
