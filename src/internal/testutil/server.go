package testutil

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/seedao/seedao-api-server/internal/app"
	"github.com/seedao/seedao-api-server/internal/config"
	"github.com/seedao/seedao-api-server/internal/model"
	"github.com/seedao/seedao-api-server/internal/pkg/phone"
	"github.com/seedao/seedao-api-server/internal/pkg/redisstore"
)

type TestServer struct {
	Echo   *echo.Echo
	DB     *gorm.DB
	Redis  *redisstore.Store
	MiniR  *miniredis.Miniredis
	Config *config.Config
}

func NewTestServer(t *testing.T) *TestServer {
	t.Helper()

	mr, err := miniredis.Run()
	require.NoError(t, err)
	mr.RequireAuth("test-redis-password")

	db, err := model.NewSQLiteDB(":memory:")
	require.NoError(t, err)
	require.NoError(t, model.AutoMigrate(db))

	cfg := &config.Config{
		Server: config.ServerConfig{Addr: ":8080", Debug: true},
		JWT: config.JWTConfig{
			Secret:   "test-secret",
			ExpHours: 24,
		},
		Redis: config.RedisConfig{
			Addr:     mr.Addr(),
			Password: "test-redis-password",
			DB:       0,
		},
		AliyunSMS: config.AliyunSMSConfig{
			SignName:     "SeeDAO",
			TemplateCode: "SMS_TEST",
			DevMode:      true,
		},
		SIWE: config.SIWEConfig{ChainID: 1},
		WechatPay: config.WechatPayConfig{
			DevMode:            true,
			H5WapName:          "SeeDAO",
			H5WapURL:           "https://app.seedao.top",
			H5ReturnURL:        "https://app.seedao.top/pay/result",
			MembershipPriceFen: 9900,
			OrderExpireMinutes: 15,
			Description:        "SeeDAO 会员续费",
		},
		Governance: config.GovernanceConfig{
			QuarterNodeTotal:           4,
			HallProposalSNS:            []string{"hall.seedao"},
			NodeConsensusProposalSNS:   []string{"node.seedao"},
			HallVoterSNS:               []string{"hall_voter.seedao"},
			NodeVoterSNS:               []string{"node_voter.seedao"},
			DefaultPublicitySeconds:    60,
			DefaultVoteDurationSeconds: 120,
		},
	}

	redis := redisstore.New(cfg.Redis)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	require.NoError(t, redis.Ping(ctx))

	e, err := app.NewEcho(cfg, db, redis)
	require.NoError(t, err)

	t.Cleanup(func() {
		_ = redis.Close()
		mr.Close()
	})

	return &TestServer{
		Echo:   e,
		DB:     db,
		Redis:  redis,
		MiniR:  mr,
		Config: cfg,
	}
}

func (ts *TestServer) SetSMSCode(t *testing.T, phoneNum, code string) {
	t.Helper()
	ctx := context.Background()
	err := ts.Redis.SetSMSCode(ctx, phone.Hash(phoneNum), code, 5*time.Minute)
	require.NoError(t, err)
}

func (ts *TestServer) SeedCourseCatalog(t *testing.T) {
	t.Helper()
	now := time.Now()
	future := now.Add(30 * 24 * time.Hour)

	require.NoError(t, ts.DB.Create(&model.Course{
		ID:        "course_test_spring",
		Title:     "Test Course",
		Status:    model.CourseStatusOnSale,
		CreatedAt: now,
		UpdatedAt: now,
	}).Error)

	require.NoError(t, ts.DB.Create(&model.CourseSKU{
		CourseID:     "course_test_spring",
		SKUID:        "standard",
		Name:         "Standard Ticket",
		BasePriceFen: 99900,
		Status:       model.CourseStatusOnSale,
		CreatedAt:    now,
		UpdatedAt:    now,
	}).Error)

	channel := "PARTNER_X"
	require.NoError(t, ts.DB.Create(&model.CourseDiscountRule{
		CourseID:      "course_test_spring",
		SKUID:         "standard",
		RuleType:      model.CourseDiscountEarlyBird,
		DiscountType:  model.CouponDiscountFixed,
		DiscountValue: 20000,
		ValidUntil:    &future,
		Status:        model.CourseStatusOnSale,
		CreatedAt:     now,
		UpdatedAt:     now,
	}).Error)

	require.NoError(t, ts.DB.Create(&model.CourseDiscountRule{
		CourseID:      "course_test_spring",
		SKUID:         "standard",
		RuleType:      model.CourseDiscountChannel,
		DiscountType:  model.CouponDiscountFixed,
		DiscountValue: 5000,
		ChannelCode:   &channel,
		Status:        model.CourseStatusOnSale,
		CreatedAt:     now,
		UpdatedAt:     now,
	}).Error)

	require.NoError(t, ts.DB.Create(&model.CourseDiscountRule{
		CourseID:      "course_test_spring",
		SKUID:         "standard",
		RuleType:      model.CourseDiscountMember,
		DiscountType:  model.CouponDiscountFixed,
		DiscountValue: 10000,
		Status:        model.CourseStatusOnSale,
		CreatedAt:     now,
		UpdatedAt:     now,
	}).Error)
}

func (ts *TestServer) SeedMembershipCoupon(t *testing.T, code string) {
	t.Helper()
	now := time.Now()
	future := now.Add(24 * time.Hour)
	productType := model.ProductTypeMembershipRenew
	require.NoError(t, ts.DB.Create(&model.Coupon{
		Code:           code,
		DiscountType:   model.CouponDiscountFixed,
		DiscountValue:  1000,
		ProductType:    &productType,
		MaxUsesPerUser: 1,
		ValidFrom:      now.Add(-time.Hour),
		ValidUntil:     future,
		Status:         model.CouponStatusActive,
		CreatedAt:      now,
		UpdatedAt:      now,
	}).Error)
}

func (ts *TestServer) SeedRegistry(t *testing.T, wallet, snsName string) {
	t.Helper()
	now := time.Now()
	require.NoError(t, ts.DB.Create(&model.SNSChainRegistry{
		Wallet:    wallet,
		SNSName:   snsName,
		Claimed:   false,
		SyncedAt:  now,
		CreatedAt: now,
		UpdatedAt: now,
	}).Error)
}

func (ts *TestServer) RegisterAndToken(t *testing.T, phoneNum, snsName string) string {
	t.Helper()
	ts.SetSMSCode(t, phoneNum, "123456")
	rec := DoJSON(t, ts.Echo, http.MethodPost, "/v1/auth/register", map[string]string{
		"phone":         phoneNum,
		"code":          "123456",
		"register_type": "new_sns",
		"sns_name":      snsName,
	}, nil)
	resp := ParseResponse(t, rec)
	data := ParseData[map[string]interface{}](t, resp.Data)
	return data["token"].(string)
}

func (ts *TestServer) SeedPublishedProposal(t *testing.T, proposerSNS string, proposalType string, voteOpen bool) int64 {
	t.Helper()
	now := time.Now()
	pubEnd := now.Add(-time.Minute)
	voteEnd := now.Add(2 * time.Hour)
	if !voteOpen {
		voteEnd = now.Add(-time.Minute)
	}
	sns := proposerSNS
	p := model.Proposal{
		Type:           proposalType,
		Title:          "Test Proposal",
		Content:        []byte(`{"plain_text":"hello"}`),
		ProposerSNS:    &sns,
		ProposerWallet: "0xabc",
		Status:         model.ProposalStatusPublicity,
		VoteType:       model.ProposalVoteTypeStandard,
		VoteOptions:    []byte(`[{"label":"同意","value":"1"}]`),
		ResolutionRule: model.ProposalResolutionStandard,
		PublishedAt:    &now,
		PublicityEndAt: &pubEnd,
		VoteEndAt:      &voteEnd,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	require.NoError(t, ts.DB.Create(&p).Error)
	return p.ID
}

func DoJSON(t *testing.T, e *echo.Echo, method, path string, body interface{}, headers map[string]string) *httptest.ResponseRecorder {
	t.Helper()

	var reader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		require.NoError(t, err)
		reader = bytes.NewReader(data)
	}

	req := httptest.NewRequest(method, path, reader)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	return rec
}

type APIResponse struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Detail  string          `json:"detail"`
	Data    json.RawMessage `json:"data"`
}

func ParseResponse(t *testing.T, rec *httptest.ResponseRecorder) APIResponse {
	t.Helper()
	require.Equal(t, http.StatusOK, rec.Code, rec.Body.String())
	var resp APIResponse
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	return resp
}

func ParseErrorResponse(t *testing.T, rec *httptest.ResponseRecorder) APIResponse {
	t.Helper()
	var resp APIResponse
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	require.NotEmpty(t, resp.Message)
	return resp
}

func ParseData[T any](t *testing.T, raw json.RawMessage) T {
	t.Helper()
	var out T
	require.NoError(t, json.Unmarshal(raw, &out))
	return out
}
