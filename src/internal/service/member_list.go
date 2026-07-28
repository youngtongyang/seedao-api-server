package service

import (
	"context"
	"strings"

	"gorm.io/gorm"

	"github.com/seedao/seedao-api-server/internal/model"
)

const (
	defaultMemberListPageSize = 20
	maxMemberListPageSize     = 100

	MemberSortByUserID              = "user_id"
	MemberSortByRegisteredAt        = "registered_at"
	MemberSortByMembershipExpiresAt = "membership_expires_at"

	MemberSortOrderAsc  = "asc"
	MemberSortOrderDesc = "desc"
)

type MemberListService struct {
	db *gorm.DB
}

func NewMemberListService(db *gorm.DB) *MemberListService {
	return &MemberListService{db: db}
}

type RegisteredMemberItem struct {
	UserID              int64  `json:"user_id"`
	SNSName             string `json:"sns_name"`
	RegisteredAt        string `json:"registered_at"`
	MembershipExpiresAt string `json:"membership_expires_at"`
}

type RegisteredMemberListResult struct {
	Page     int                    `json:"page"`
	PageSize int                    `json:"page_size"`
	Total    int64                  `json:"total"`
	SortBy   string                 `json:"sort_by"`
	Order    string                 `json:"order"`
	Items    []RegisteredMemberItem `json:"items"`
}

type ListRegisteredMembersInput struct {
	Page     int
	PageSize int
	SortBy   string
	Order    string
}

func (s *MemberListService) ListRegisteredMembers(ctx context.Context, input ListRegisteredMembersInput) (*RegisteredMemberListResult, error) {
	_ = ctx
	page := input.Page
	pageSize := input.PageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = defaultMemberListPageSize
	}
	if pageSize > maxMemberListPageSize {
		pageSize = maxMemberListPageSize
	}

	sortBy, order, orderClause := normalizeMemberListSort(input.SortBy, input.Order)

	base := s.db.Model(&model.User{}).Where("deleted_at IS NULL")

	var total int64
	if err := base.Count(&total).Error; err != nil {
		return nil, err
	}

	var users []model.User
	offset := (page - 1) * pageSize
	if err := s.db.Preload("SNSName").
		Where("deleted_at IS NULL").
		Order(orderClause).
		Offset(offset).
		Limit(pageSize).
		Find(&users).Error; err != nil {
		return nil, err
	}

	items := make([]RegisteredMemberItem, 0, len(users))
	for i := range users {
		u := &users[i]
		item := RegisteredMemberItem{
			UserID:              u.ID,
			RegisteredAt:        u.CreatedAt.UTC().Format("2006-01-02T15:04:05Z"),
			MembershipExpiresAt: formatMembershipExpiry(u.MembershipExpiresAt),
		}
		if u.SNSName != nil {
			item.SNSName = u.SNSName.Name
		}
		items = append(items, item)
	}

	return &RegisteredMemberListResult{
		Page:     page,
		PageSize: pageSize,
		Total:    total,
		SortBy:   sortBy,
		Order:    order,
		Items:    items,
	}, nil
}

func normalizeMemberListSort(sortBy, order string) (string, string, string) {
	sortBy = strings.ToLower(strings.TrimSpace(sortBy))
	order = strings.ToLower(strings.TrimSpace(order))

	column := "id"
	switch sortBy {
	case "", MemberSortByUserID:
		sortBy = MemberSortByUserID
		column = "id"
	case MemberSortByRegisteredAt:
		column = "created_at"
	case MemberSortByMembershipExpiresAt:
		column = "membership_expires_at"
	default:
		sortBy = MemberSortByUserID
		column = "id"
	}

	if order != MemberSortOrderDesc {
		order = MemberSortOrderAsc
	}

	direction := "ASC"
	if order == MemberSortOrderDesc {
		direction = "DESC"
	}

	// Stable secondary key for pagination consistency.
	return sortBy, order, column + " " + direction + ", id ASC"
}
