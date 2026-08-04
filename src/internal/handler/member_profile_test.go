package handler_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/seedao/seedao-api-server/internal/model"
	"github.com/seedao/seedao-api-server/internal/testutil"
)

func TestSNSMemberProfile(t *testing.T) {
	ts := testutil.NewTestServer(t)
	now := time.Now()
	sns := "baiyu.seedao"
	scr := "100.5"
	require.NoError(t, ts.DB.Create(&model.UserProfile{
		SNSName: sns, Bio: "hello", Email: "a@b.c", TwitterProfile: "https://x.com/a",
		Mirror: "https://mirror.xyz/a", SCRAmount: &scr, CreatedAt: now, UpdatedAt: now,
	}).Error)
	require.NoError(t, ts.DB.Create(&model.UserAsset{
		SNSName: sns, AssetName: "SEE", DealtAmount: "10", ProcessingAmount: "0", Claimed: true,
		CreatedAt: now, UpdatedAt: now,
	}).Error)
	require.NoError(t, ts.DB.Create(&model.UserAsset{
		SNSName: sns, AssetName: "WANG", DealtAmount: "20", ProcessingAmount: "1", Claimed: false,
		CreatedAt: now, UpdatedAt: now,
	}).Error)
	require.NoError(t, ts.DB.Create(&model.UserRole{
		SNSName: sns, Role: "SEEDAO_MEMBER", CreatedAt: now,
	}).Error)
	require.NoError(t, ts.DB.Create(&model.UserSeed{
		SNSName: sns, TokenID: "17", ContractAddr: "0xabc", ContractType: "erc721",
		TokenAmount: "1", ImageURISrc: "ipfs://x", ImageURIOSS: "https://oss/seed.png",
		CreatedAt: now, UpdatedAt: now,
	}).Error)
	require.NoError(t, ts.DB.Create(&model.UserSBT{
		SNSName: sns, TokenID: "3", ContractAddr: "0xdef", CollectionName: "NODE_S",
		DisplayName: "S3", ImageURISrc: "https://src/a.png", ImageURIOSS: "",
		CreatedAt: now, UpdatedAt: now,
	}).Error)

	rec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/v1/sns/profile?sns_name="+sns, nil, nil)
	resp := testutil.ParseResponse(t, rec)
	data := testutil.ParseData[map[string]interface{}](t, resp.Data)

	assert.Equal(t, sns, data["sns_name"])
	assert.Equal(t, "hello", data["bio"])
	assert.Equal(t, "100.5", data["scr_amount"])
	assert.Equal(t, []interface{}{"SEEDAO_MEMBER"}, data["roles"])
	assert.Equal(t, "10", data["see"].(map[string]interface{})["dealt_amount"])
	assert.Equal(t, "20", data["wang"].(map[string]interface{})["dealt_amount"])
	seeds := data["seeds"].([]interface{})
	require.Len(t, seeds, 1)
	assert.Equal(t, "https://oss/seed.png", seeds[0].(map[string]interface{})["image_url"])
	sbts := data["sbts"].([]interface{})
	require.Len(t, sbts, 1)
	assert.Equal(t, "https://src/a.png", sbts[0].(map[string]interface{})["image_url"])
}

func TestSNSMemberProfile_Empty(t *testing.T) {
	ts := testutil.NewTestServer(t)
	rec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/v1/sns/profile?sns_name=nobody.seedao", nil, nil)
	resp := testutil.ParseResponse(t, rec)
	data := testutil.ParseData[map[string]interface{}](t, resp.Data)
	assert.Equal(t, "nobody.seedao", data["sns_name"])
	assert.Equal(t, []interface{}{}, data["roles"])
	assert.Equal(t, []interface{}{}, data["assets"])
	assert.Equal(t, []interface{}{}, data["seeds"])
	assert.Equal(t, []interface{}{}, data["sbts"])
}

func TestSNSMemberProfile_InvalidSNS(t *testing.T) {
	ts := testutil.NewTestServer(t)
	rec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/v1/sns/profile?sns_name=bad", nil, nil)
	require.Equal(t, http.StatusBadRequest, rec.Code)
}
