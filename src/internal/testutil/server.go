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

func (ts *TestServer) SeedCoupon(t *testing.T, code string, discountFen int) {
	t.Helper()
	now := time.Now()
	productType := model.ProductTypeMembershipRenew
	require.NoError(t, ts.DB.Create(&model.Coupon{
		Code:           code,
		DiscountType:   model.CouponDiscountFixed,
		DiscountValue:  discountFen,
		ProductType:    &productType,
		MaxUsesPerUser: 1,
		ValidFrom:      now.Add(-time.Hour),
		ValidUntil:     now.Add(24 * time.Hour),
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
