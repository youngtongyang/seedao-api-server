package handler_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/seedao/seedao-api-server/internal/model"
	"github.com/seedao/seedao-api-server/internal/pkg/errcode"
	"github.com/seedao/seedao-api-server/internal/testutil"
)

func TestBindings_ListUnbound(t *testing.T) {
	ts := testutil.NewTestServer(t)
	token := ts.RegisterAndToken(t, "+8613800138001", "bindtest1.seedao")

	rec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/v1/user/bindings", nil, map[string]string{
		"Authorization": "Bearer " + token,
	})
	resp := testutil.ParseResponse(t, rec)
	data := testutil.ParseData[map[string][]map[string]interface{}](t, resp.Data)
	require.Len(t, data["items"], 1)
	assert.Equal(t, "feishu", data["items"][0]["provider"])
	assert.Equal(t, false, data["items"][0]["bound"])
}

func TestBindings_BindFeishuWithToken(t *testing.T) {
	ts := testutil.NewTestServer(t)
	token := ts.RegisterAndToken(t, "+8613800138002", "bindtest2.seedao")

	createRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/internal/v1/feishu/bindings/token", map[string]string{
		"open_id":      "ou_test_user_1",
		"display_name": "Alice",
	}, map[string]string{
		"X-Service-Token": "test-service-token",
	})
	createResp := testutil.ParseResponse(t, createRec)
	createData := testutil.ParseData[map[string]interface{}](t, createResp.Data)
	bindToken := createData["bind_token"].(string)
	require.NotEmpty(t, bindToken)

	bindRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/user/bindings/feishu/token", map[string]string{
		"bind_token": bindToken,
	}, map[string]string{
		"Authorization": "Bearer " + token,
	})
	bindResp := testutil.ParseResponse(t, bindRec)
	bindData := testutil.ParseData[map[string]interface{}](t, bindResp.Data)
	assert.Equal(t, true, bindData["bound"])
	assert.Equal(t, "Alice", bindData["display_name"])
	assert.Equal(t, model.OAuthBindSourceBotToken, bindData["bind_source"])

	listRec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/v1/user/bindings", nil, map[string]string{
		"Authorization": "Bearer " + token,
	})
	listResp := testutil.ParseResponse(t, listRec)
	listData := testutil.ParseData[map[string][]map[string]interface{}](t, listResp.Data)
	assert.Equal(t, true, listData["items"][0]["bound"])
}

func TestBindings_BindFeishuTokenReuseFails(t *testing.T) {
	ts := testutil.NewTestServer(t)
	token := ts.RegisterAndToken(t, "+8613800138003", "bindtest3.seedao")

	createRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/internal/v1/feishu/bindings/token", map[string]string{
		"open_id": "ou_test_user_2",
	}, map[string]string{
		"X-Service-Token": "test-service-token",
	})
	createData := testutil.ParseData[map[string]interface{}](t, testutil.ParseResponse(t, createRec).Data)
	bindToken := createData["bind_token"].(string)

	headers := map[string]string{"Authorization": "Bearer " + token}
	rec1 := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/user/bindings/feishu/token", map[string]string{
		"bind_token": bindToken,
	}, headers)
	require.Equal(t, http.StatusOK, rec1.Code)

	rec2 := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/user/bindings/feishu/token", map[string]string{
		"bind_token": bindToken,
	}, headers)
	resp2 := testutil.ParseErrorResponse(t, rec2)
	assert.Equal(t, errcode.FeishuBindCredentialInvalid, resp2.Code)
}

func TestBindings_ConfirmFeishuBinding(t *testing.T) {
	ts := testutil.NewTestServer(t)
	token := ts.RegisterAndToken(t, "+8613800138004", "bindtest4.seedao")

	challengeRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/user/bindings/feishu/challenge", nil, map[string]string{
		"Authorization": "Bearer " + token,
	})
	challengeData := testutil.ParseData[map[string]interface{}](t, testutil.ParseResponse(t, challengeRec).Data)
	bindCode := challengeData["bind_code"].(string)

	confirmRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/internal/v1/feishu/bindings/confirm", map[string]string{
		"bind_code":    bindCode,
		"open_id":      "ou_test_user_3",
		"display_name": "Bob",
	}, map[string]string{
		"X-Service-Token": "test-service-token",
	})
	confirmResp := testutil.ParseResponse(t, confirmRec)
	confirmData := testutil.ParseData[map[string]interface{}](t, confirmResp.Data)
	assert.Equal(t, true, confirmData["bound"])
	assert.Equal(t, "bindtest4.seedao", confirmData["sns_name"])
	assert.Equal(t, model.OAuthBindSourceBotCode, confirmData["bind_source"])
}

func TestBindings_UnbindFeishu(t *testing.T) {
	ts := testutil.NewTestServer(t)
	token := ts.RegisterAndToken(t, "+8613800138005", "bindtest5.seedao")

	createRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/internal/v1/feishu/bindings/token", map[string]string{
		"open_id": "ou_unbind_me",
	}, map[string]string{"X-Service-Token": "test-service-token"})
	createData := testutil.ParseData[map[string]interface{}](t, testutil.ParseResponse(t, createRec).Data)
	testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/user/bindings/feishu/token", map[string]string{
		"bind_token": createData["bind_token"].(string),
	}, map[string]string{"Authorization": "Bearer " + token})

	unbindRec := testutil.DoJSON(t, ts.Echo, http.MethodDelete, "/v1/user/bindings/feishu", nil, map[string]string{
		"Authorization": "Bearer " + token,
	})
	unbindData := testutil.ParseData[map[string]interface{}](t, testutil.ParseResponse(t, unbindRec).Data)
	assert.Equal(t, false, unbindData["bound"])
}

func TestBindings_OpenIDTaken(t *testing.T) {
	ts := testutil.NewTestServer(t)
	token1 := ts.RegisterAndToken(t, "+8613800138006", "bindtest6.seedao")
	token2 := ts.RegisterAndToken(t, "+8613800138007", "bindtest7.seedao")

	createRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/internal/v1/feishu/bindings/token", map[string]string{
		"open_id": "ou_shared",
	}, map[string]string{"X-Service-Token": "test-service-token"})
	createData := testutil.ParseData[map[string]interface{}](t, testutil.ParseResponse(t, createRec).Data)
	bindToken := createData["bind_token"].(string)

	rec1 := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/user/bindings/feishu/token", map[string]string{
		"bind_token": bindToken,
	}, map[string]string{"Authorization": "Bearer " + token1})
	require.Equal(t, http.StatusOK, rec1.Code)

	createRec2 := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/internal/v1/feishu/bindings/token", map[string]string{
		"open_id": "ou_shared",
	}, map[string]string{"X-Service-Token": "test-service-token"})
	createData2 := testutil.ParseData[map[string]interface{}](t, testutil.ParseResponse(t, createRec2).Data)

	rec2 := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/user/bindings/feishu/token", map[string]string{
		"bind_token": createData2["bind_token"].(string),
	}, map[string]string{"Authorization": "Bearer " + token2})
	resp2 := testutil.ParseErrorResponse(t, rec2)
	assert.Equal(t, errcode.FeishuOpenIDTaken, resp2.Code)
}

func TestInternalFeishu_Unauthorized(t *testing.T) {
	ts := testutil.NewTestServer(t)

	rec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/internal/v1/feishu/bindings/token", map[string]string{
		"open_id": "ou_x",
	}, nil)
	require.Equal(t, http.StatusUnauthorized, rec.Code)
}
