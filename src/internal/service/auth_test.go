package service

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/seedao/seedao-api-server/internal/config"
	"github.com/seedao/seedao-api-server/internal/model"
	"github.com/seedao/seedao-api-server/internal/pkg/jwtutil"
	"github.com/seedao/seedao-api-server/internal/pkg/phone"
	"github.com/seedao/seedao-api-server/internal/pkg/redisstore"
)

func newAuthTestService(t *testing.T) (*AuthService, *gorm.DB, *redisstore.Store) {
	t.Helper()

	mr, err := miniredis.Run()
	require.NoError(t, err)
	mr.RequireAuth("test-redis-password")
	t.Cleanup(mr.Close)

	db, err := model.NewSQLiteDB(":memory:")
	require.NoError(t, err)
	require.NoError(t, model.AutoMigrate(db))

	redis := redisstore.New(config.RedisConfig{
		Addr:     mr.Addr(),
		Password: "test-redis-password",
		DB:       0,
	})
	t.Cleanup(func() { _ = redis.Close() })

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	require.NoError(t, redis.Ping(ctx))

	sms := NewSMSService(db, redis, nil)
	sns := NewSNSService(db, redis)
	wallet := NewWalletService(redis, config.SIWEConfig{ChainID: 1})
	jwtMgr := jwtutil.NewManager("test-secret", 24)
	return NewAuthService(db, sms, sns, wallet, jwtMgr), db, redis
}

func signPersonalMessage(t *testing.T, message string, key *ecdsa.PrivateKey) string {
	t.Helper()
	prefixed := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256Hash([]byte(prefixed))
	sig, err := crypto.Sign(hash.Bytes(), key)
	require.NoError(t, err)
	sig[64] += 27
	return "0x" + hex.EncodeToString(sig)
}

func TestRegisterExistingSNS_GrantsDefaultMembership(t *testing.T) {
	auth, db, redis := newAuthTestService(t)

	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	wallet := crypto.PubkeyToAddress(key.PublicKey).Hex()
	snsName := "legacy_claim.seedao"
	domain := "app.seedao.top"
	phoneNum := "+8613900199001"
	code := "123456"

	now := time.Now()
	require.NoError(t, db.Create(&model.SNSChainRegistry{
		Wallet:    phone.NormalizeWallet(wallet),
		SNSName:   snsName,
		Claimed:   false,
		SyncedAt:  now,
		CreatedAt: now,
		UpdatedAt: now,
	}).Error)

	ctx := context.Background()
	require.NoError(t, redis.SetSMSCode(ctx, phone.Hash(phoneNum), code, 5*time.Minute))

	nonceRes, err := auth.wallet.CreateNonce(ctx, wallet, domain)
	require.NoError(t, err)
	sig := signPersonalMessage(t, nonceRes.Message, key)

	before := time.Now()
	result, err := auth.RegisterExistingSNS(ctx, RegisterExistingSNSInput{
		Phone:      phoneNum,
		Code:       code,
		Wallet:     wallet,
		WalletType: "evm",
		Domain:     domain,
		Message:    nonceRes.Message,
		Signature:  sig,
	})
	require.NoError(t, err)
	after := time.Now()

	require.NotEmpty(t, result.User.MembershipExpiresAt)
	assert.Equal(t, model.RegisterSourceExistingSNS, result.User.RegisterSource)
	assert.Equal(t, snsName, result.User.SNSName)

	expires, err := time.Parse(time.RFC3339, result.User.MembershipExpiresAt)
	require.NoError(t, err)

	minExp := before.AddDate(0, 0, model.DefaultMembershipDays).Add(-2 * time.Second)
	maxExp := after.AddDate(0, 0, model.DefaultMembershipDays).Add(2 * time.Second)
	assert.True(t, !expires.Before(minExp) && !expires.After(maxExp), "expires=%s range=[%s,%s]", expires, minExp, maxExp)

	var user model.User
	require.NoError(t, db.First(&user, result.User.ID).Error)
	require.NotNil(t, user.MembershipExpiresAt)
	assert.WithinDuration(t, expires, user.MembershipExpiresAt.UTC(), time.Second)
}

func TestRegisterNewSNS_NoMembershipUntilPaid(t *testing.T) {
	auth, db, redis := newAuthTestService(t)
	phoneNum := "+8613900199002"
	code := "123456"
	ctx := context.Background()
	require.NoError(t, redis.SetSMSCode(ctx, phone.Hash(phoneNum), code, 5*time.Minute))

	result, err := auth.RegisterNewSNS(ctx, RegisterNewSNSInput{
		Phone:   phoneNum,
		Code:    code,
		SNSName: "new_paid_only.seedao",
	})
	require.NoError(t, err)
	assert.Empty(t, result.User.MembershipExpiresAt)

	var user model.User
	require.NoError(t, db.First(&user, result.User.ID).Error)
	assert.Nil(t, user.MembershipExpiresAt)
}
