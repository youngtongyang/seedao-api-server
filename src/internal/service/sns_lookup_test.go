package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/seedao/seedao-api-server/internal/model"
)

func newSNSTestService(t *testing.T) *SNSService {
	t.Helper()
	db, err := model.NewSQLiteDB(":memory:")
	require.NoError(t, err)
	require.NoError(t, model.AutoMigrate(db))
	return NewSNSService(db, nil)
}

func TestLookupSNSByWallets(t *testing.T) {
	svc := newSNSTestService(t)
	now := time.Now()
	wallet := "0x1b76b011c0a1e03896bfc8923be228e430bb9049"
	nameWallet := "0xAbCdEf1234567890123456789012345678901234"
	lower := NormalizeWallet(nameWallet)

	require.NoError(t, svc.db.Create(&model.SNSName{
		Name: "alice", UserID: 1, MintStatus: "minted", Wallet: &lower, CreatedAt: now,
	}).Error)
	require.NoError(t, svc.db.Create(&model.SNSChainRegistry{
		Wallet: wallet, SNSName: "nechi_luo.seedao", SyncedAt: now, CreatedAt: now, UpdatedAt: now,
	}).Error)

	names, err := svc.LookupSNSByWallets([]string{nameWallet, wallet, "", zeroWallet})
	require.NoError(t, err)
	assert.Equal(t, "alice", names[lower])
	assert.Equal(t, "nechi_luo.seedao", names[NormalizeWallet(wallet)])
}

func TestResolveSNSNameForDisplay(t *testing.T) {
	existing := "shawn"
	byWallet := map[string]string{NormalizeWallet("0xabc"): "resolved.seedao"}

	assert.Equal(t, &existing, ResolveSNSNameForDisplay(&existing, "0xabc", byWallet))

	var nilSNS *string
	resolved := ResolveSNSNameForDisplay(nilSNS, "0xabc", byWallet)
	require.NotNil(t, resolved)
	assert.Equal(t, "resolved.seedao", *resolved)

	missing := ResolveSNSNameForDisplay(nilSNS, "0x1b76b011c0a1e03896bfc8923be228e430bb9049", map[string]string{})
	require.NotNil(t, missing)
	assert.Equal(t, "0x1b76...9049", *missing)
}

func TestFormatWalletDisplay(t *testing.T) {
	assert.Equal(t, "0x1b76...9049", FormatWalletDisplay("0x1b76b011c0a1e03896bfc8923be228e430bb9049"))
	assert.Equal(t, "0x1234", FormatWalletDisplay("0x1234"))
	assert.True(t, IsWalletDisplayName("0x1b76...9049"))
	assert.False(t, IsWalletDisplayName("nechi_luo.seedao"))
}
