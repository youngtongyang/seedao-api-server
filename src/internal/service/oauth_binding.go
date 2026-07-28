package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/seedao/seedao-api-server/internal/config"
	"github.com/seedao/seedao-api-server/internal/model"
	"github.com/seedao/seedao-api-server/internal/pkg/bindtoken"
	"github.com/seedao/seedao-api-server/internal/pkg/redisstore"
)

var (
	ErrFeishuNotBound            = errors.New("feishu not bound")
	ErrFeishuAlreadyBound        = errors.New("feishu already bound")
	ErrFeishuOpenIDTaken         = errors.New("feishu open id taken")
	ErrFeishuBindCredentialInvalid = errors.New("feishu bind credential invalid")
)

type OAuthBindingService struct {
	db     *gorm.DB
	redis  *redisstore.Store
	feishu config.FeishuConfig
}

func NewOAuthBindingService(db *gorm.DB, redis *redisstore.Store, feishu config.FeishuConfig) *OAuthBindingService {
	return &OAuthBindingService{db: db, redis: redis, feishu: feishu}
}

type BindingItem struct {
	Provider    string `json:"provider"`
	Bound       bool   `json:"bound"`
	DisplayName string `json:"display_name,omitempty"`
	BoundAt     string `json:"bound_at,omitempty"`
	BindSource  string `json:"bind_source,omitempty"`
}

type BindingListResult struct {
	Items []BindingItem `json:"items"`
}

type FeishuBindResult struct {
	Provider    string `json:"provider"`
	Bound       bool   `json:"bound"`
	DisplayName string `json:"display_name,omitempty"`
	BindSource  string `json:"bind_source,omitempty"`
}

type FeishuChallengeResult struct {
	BindCode    string `json:"bind_code"`
	ExpiresIn   int    `json:"expires_in"`
	Instruction string `json:"instruction"`
}

type FeishuCreateTokenInput struct {
	OpenID        string
	UnionID       string
	ProviderAppID string
	DisplayName   string
}

type FeishuCreateTokenResult struct {
	BindToken string `json:"bind_token"`
	BindURL   string `json:"bind_url"`
	ExpiresIn int    `json:"expires_in"`
}

type FeishuConfirmInput struct {
	BindCode      string
	OpenID        string
	UnionID       string
	ProviderAppID string
	DisplayName   string
}

type FeishuConfirmResult struct {
	UserID     int64  `json:"user_id"`
	SNSName    string `json:"sns_name,omitempty"`
	Bound      bool   `json:"bound"`
	BindSource string `json:"bind_source"`
}

type FeishuUnbindByOpenIDInput struct {
	OpenID        string
	ProviderAppID string
}

type FeishuUnbindByOpenIDResult struct {
	UserID  int64  `json:"user_id"`
	SNSName string `json:"sns_name,omitempty"`
	Bound   bool   `json:"bound"`
}

type FeishuLookupByOpenIDInput struct {
	OpenID        string
	ProviderAppID string
}

type FeishuLookupByOpenIDResult struct {
	Bound       bool   `json:"bound"`
	UserID      int64  `json:"user_id,omitempty"`
	SNSName     string `json:"sns_name,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	BoundAt     string `json:"bound_at,omitempty"`
	BindSource  string `json:"bind_source,omitempty"`
}

func (s *OAuthBindingService) ListByUserID(ctx context.Context, userID int64) (*BindingListResult, error) {
	item := BindingItem{
		Provider: model.OAuthProviderFeishu,
		Bound:    false,
	}

	binding, err := s.getActiveFeishuBinding(userID)
	if err != nil {
		return nil, err
	}
	if binding != nil {
		item.Bound = true
		item.DisplayName = binding.DisplayName
		item.BoundAt = binding.BoundAt.UTC().Format(time.RFC3339)
		item.BindSource = binding.BindSource
	}

	return &BindingListResult{Items: []BindingItem{item}}, nil
}

func (s *OAuthBindingService) BindFeishuWithToken(ctx context.Context, userID int64, bindToken string) (*FeishuBindResult, error) {
	bindToken = strings.TrimSpace(bindToken)
	if bindToken == "" {
		return nil, ErrInvalidParam
	}

	payload, err := s.redis.ConsumeFeishuBindToken(ctx, bindToken)
	if err != nil {
		return nil, err
	}
	if payload == nil || strings.TrimSpace(payload.OpenID) == "" {
		return nil, ErrFeishuBindCredentialInvalid
	}

	appID := strings.TrimSpace(payload.ProviderAppID)
	if appID == "" {
		appID = strings.TrimSpace(s.feishu.AppID)
	}
	if appID == "" {
		return nil, ErrInvalidParam
	}

	binding, err := s.createFeishuBinding(ctx, userID, feishuBindingInput{
		ProviderAppID: appID,
		OpenID:        payload.OpenID,
		UnionID:       payload.UnionID,
		DisplayName:   payload.DisplayName,
		BindSource:    model.OAuthBindSourceBotToken,
	})
	if err != nil {
		return nil, err
	}

	return toFeishuBindResult(binding), nil
}

func (s *OAuthBindingService) CreateFeishuChallenge(ctx context.Context, userID int64) (*FeishuChallengeResult, error) {
	if _, err := s.ensureUserExists(userID); err != nil {
		return nil, err
	}
	if binding, err := s.getActiveFeishuBinding(userID); err != nil {
		return nil, err
	} else if binding != nil {
		return nil, ErrFeishuAlreadyBound
	}

	code, err := bindtoken.GenerateBindCode()
	if err != nil {
		return nil, err
	}

	ttl := time.Duration(s.feishu.BindCodeTTLSeconds) * time.Second
	if err := s.redis.SetFeishuBindCode(ctx, code, redisstore.FeishuBindCodePayload{UserID: userID}, ttl); err != nil {
		return nil, err
	}

	return &FeishuChallengeResult{
		BindCode:    code,
		ExpiresIn:   s.feishu.BindCodeTTLSeconds,
		Instruction: "在飞书对土地神发送：绑定 " + code,
	}, nil
}

func (s *OAuthBindingService) CreateFeishuBindToken(ctx context.Context, input FeishuCreateTokenInput) (*FeishuCreateTokenResult, error) {
	openID := strings.TrimSpace(input.OpenID)
	if openID == "" {
		return nil, ErrInvalidParam
	}

	appID := strings.TrimSpace(input.ProviderAppID)
	if appID == "" {
		appID = strings.TrimSpace(s.feishu.AppID)
	}
	if appID == "" {
		return nil, ErrInvalidParam
	}

	token, err := bindtoken.GenerateBindToken()
	if err != nil {
		return nil, err
	}

	ttl := time.Duration(s.feishu.BindTokenTTLSeconds) * time.Second
	payload := redisstore.FeishuBindTokenPayload{
		OpenID:        openID,
		UnionID:       strings.TrimSpace(input.UnionID),
		ProviderAppID: appID,
		DisplayName:   strings.TrimSpace(input.DisplayName),
	}
	if err := s.redis.SetFeishuBindToken(ctx, token, payload, ttl); err != nil {
		return nil, err
	}

	bindURL := strings.ReplaceAll(s.feishu.BindURLTemplate, "{token}", token)
	return &FeishuCreateTokenResult{
		BindToken: token,
		BindURL:   bindURL,
		ExpiresIn: s.feishu.BindTokenTTLSeconds,
	}, nil
}

func (s *OAuthBindingService) ConfirmFeishuBinding(ctx context.Context, input FeishuConfirmInput) (*FeishuConfirmResult, error) {
	code := strings.TrimSpace(input.BindCode)
	openID := strings.TrimSpace(input.OpenID)
	if code == "" || openID == "" {
		return nil, ErrInvalidParam
	}

	payload, err := s.redis.ConsumeFeishuBindCode(ctx, code)
	if err != nil {
		return nil, err
	}
	if payload == nil || payload.UserID == 0 {
		return nil, ErrFeishuBindCredentialInvalid
	}

	appID := strings.TrimSpace(input.ProviderAppID)
	if appID == "" {
		appID = strings.TrimSpace(s.feishu.AppID)
	}
	if appID == "" {
		return nil, ErrInvalidParam
	}

	binding, err := s.createFeishuBinding(ctx, payload.UserID, feishuBindingInput{
		ProviderAppID: appID,
		OpenID:        openID,
		UnionID:       strings.TrimSpace(input.UnionID),
		DisplayName:   strings.TrimSpace(input.DisplayName),
		BindSource:    model.OAuthBindSourceBotCode,
	})
	if err != nil {
		return nil, err
	}

	result := &FeishuConfirmResult{
		UserID:     binding.UserID,
		Bound:      true,
		BindSource: binding.BindSource,
	}

	var sns model.SNSName
	if err := s.db.Where("user_id = ?", binding.UserID).First(&sns).Error; err == nil {
		result.SNSName = sns.Name
	}

	return result, nil
}

func (s *OAuthBindingService) UnbindFeishu(ctx context.Context, userID int64) (*FeishuBindResult, error) {
	_ = ctx
	binding, err := s.getActiveFeishuBinding(userID)
	if err != nil {
		return nil, err
	}
	if binding == nil {
		return nil, ErrFeishuNotBound
	}

	if err := s.db.Delete(binding).Error; err != nil {
		return nil, err
	}

	return &FeishuBindResult{
		Provider: model.OAuthProviderFeishu,
		Bound:    false,
	}, nil
}

func (s *OAuthBindingService) UnbindFeishuByOpenID(ctx context.Context, input FeishuUnbindByOpenIDInput) (*FeishuUnbindByOpenIDResult, error) {
	_ = ctx
	openID := strings.TrimSpace(input.OpenID)
	if openID == "" {
		return nil, ErrInvalidParam
	}

	appID := strings.TrimSpace(input.ProviderAppID)
	if appID == "" {
		appID = strings.TrimSpace(s.feishu.AppID)
	}
	if appID == "" {
		return nil, ErrInvalidParam
	}

	binding, err := s.getActiveFeishuBindingByOpenID(appID, openID)
	if err != nil {
		return nil, err
	}
	if binding == nil {
		return nil, ErrFeishuNotBound
	}

	result := &FeishuUnbindByOpenIDResult{
		UserID: binding.UserID,
		Bound:  false,
	}

	var sns model.SNSName
	if err := s.db.Where("user_id = ?", binding.UserID).First(&sns).Error; err == nil {
		result.SNSName = sns.Name
	}

	if err := s.db.Delete(binding).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (s *OAuthBindingService) LookupFeishuByOpenID(ctx context.Context, input FeishuLookupByOpenIDInput) (*FeishuLookupByOpenIDResult, error) {
	_ = ctx
	openID := strings.TrimSpace(input.OpenID)
	if openID == "" {
		return nil, ErrInvalidParam
	}

	appID := strings.TrimSpace(input.ProviderAppID)
	if appID == "" {
		appID = strings.TrimSpace(s.feishu.AppID)
	}
	if appID == "" {
		return nil, ErrInvalidParam
	}

	binding, err := s.getActiveFeishuBindingByOpenID(appID, openID)
	if err != nil {
		return nil, err
	}
	if binding == nil {
		return &FeishuLookupByOpenIDResult{Bound: false}, nil
	}

	result := &FeishuLookupByOpenIDResult{
		Bound:       true,
		UserID:      binding.UserID,
		DisplayName: binding.DisplayName,
		BoundAt:     binding.BoundAt.UTC().Format(time.RFC3339),
		BindSource:  binding.BindSource,
	}

	var sns model.SNSName
	if err := s.db.Where("user_id = ?", binding.UserID).First(&sns).Error; err == nil {
		result.SNSName = sns.Name
	}

	return result, nil
}

type feishuBindingInput struct {
	ProviderAppID string
	OpenID        string
	UnionID       string
	DisplayName   string
	BindSource    string
}

func (s *OAuthBindingService) createFeishuBinding(ctx context.Context, userID int64, input feishuBindingInput) (*model.UserOAuthBinding, error) {
	_ = ctx
	if _, err := s.ensureUserExists(userID); err != nil {
		return nil, err
	}

	if existing, err := s.getActiveFeishuBinding(userID); err != nil {
		return nil, err
	} else if existing != nil {
		return nil, ErrFeishuAlreadyBound
	}

	var taken model.UserOAuthBinding
	err := s.db.Where(
		"provider = ? AND provider_app_id = ? AND open_id = ?",
		model.OAuthProviderFeishu, input.ProviderAppID, input.OpenID,
	).First(&taken).Error
	if err == nil {
		if taken.UserID != userID {
			return nil, ErrFeishuOpenIDTaken
		}
		return &taken, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	now := time.Now()
	binding := model.UserOAuthBinding{
		UserID:        userID,
		Provider:      model.OAuthProviderFeishu,
		ProviderAppID: input.ProviderAppID,
		OpenID:        input.OpenID,
		DisplayName:   input.DisplayName,
		BindSource:    input.BindSource,
		BoundAt:       now,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
	if input.UnionID != "" {
		unionID := input.UnionID
		binding.UnionID = &unionID
	}

	if err := s.db.Create(&binding).Error; err != nil {
		return nil, err
	}
	return &binding, nil
}

func (s *OAuthBindingService) getActiveFeishuBinding(userID int64) (*model.UserOAuthBinding, error) {
	var binding model.UserOAuthBinding
	err := s.db.Where("user_id = ? AND provider = ?", userID, model.OAuthProviderFeishu).First(&binding).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &binding, nil
}

func (s *OAuthBindingService) getActiveFeishuBindingByOpenID(providerAppID, openID string) (*model.UserOAuthBinding, error) {
	var binding model.UserOAuthBinding
	err := s.db.Where(
		"provider = ? AND provider_app_id = ? AND open_id = ?",
		model.OAuthProviderFeishu, providerAppID, openID,
	).First(&binding).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &binding, nil
}

func (s *OAuthBindingService) ensureUserExists(userID int64) (*model.User, error) {
	var user model.User
	err := s.db.First(&user, userID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func toFeishuBindResult(binding *model.UserOAuthBinding) *FeishuBindResult {
	return &FeishuBindResult{
		Provider:    model.OAuthProviderFeishu,
		Bound:       true,
		DisplayName: binding.DisplayName,
		BindSource:  binding.BindSource,
	}
}
