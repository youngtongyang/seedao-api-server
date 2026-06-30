package handler_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/seedao/seedao-api-server/internal/pkg/errcode"
	"github.com/seedao/seedao-api-server/internal/service"
	"github.com/seedao/seedao-api-server/internal/testutil"
)

func TestHealth(t *testing.T) {
	ts := testutil.NewTestServer(t)

	rec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/health", nil, nil)
	resp := testutil.ParseResponse(t, rec)

	assert.Equal(t, 0, resp.Code)
	data := testutil.ParseData[map[string]string](t, resp.Data)
	assert.Equal(t, "ok", data["status"])
}

func TestSendSMS(t *testing.T) {
	ts := testutil.NewTestServer(t)

	rec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/auth/sms/send", map[string]string{
		"phone": "+8613800138000",
	}, nil)
	resp := testutil.ParseResponse(t, rec)

	assert.Equal(t, 0, resp.Code)
	data := testutil.ParseData[map[string]int](t, resp.Data)
	assert.Equal(t, 60, data["retry_after"])
}

func TestSendSMS_InvalidPhone(t *testing.T) {
	ts := testutil.NewTestServer(t)

	rec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/auth/sms/send", map[string]string{
		"phone": "13800138000",
	}, nil)
	require.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestSendSMS_RateLimited(t *testing.T) {
	ts := testutil.NewTestServer(t)
	phoneNum := "+8613900138999"

	rec1 := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/auth/sms/send", map[string]string{
		"phone": phoneNum,
	}, nil)
	require.Equal(t, http.StatusOK, rec1.Code)

	rec2 := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/auth/sms/send", map[string]string{
		"phone": phoneNum,
	}, nil)
	require.Equal(t, http.StatusTooManyRequests, rec2.Code)

	resp := testutil.ParseErrorResponse(t, rec2)
	assert.Equal(t, errcode.TooManyRequests, resp.Code)
	data := testutil.ParseData[map[string]int](t, resp.Data)
	assert.Greater(t, data["retry_after"], 0)
	assert.LessOrEqual(t, data["retry_after"], 60)
}

func TestPhoneCheck_NotRegistered(t *testing.T) {
	ts := testutil.NewTestServer(t)

	rec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/v1/auth/phone/check?phone=%2B8613900138999", nil, nil)
	resp := testutil.ParseResponse(t, rec)

	assert.Equal(t, errcode.PhoneNotRegistered, resp.Code)
	data := testutil.ParseData[map[string]interface{}](t, resp.Data)
	assert.Equal(t, "+8613900138999", data["phone"])
	assert.Equal(t, false, data["registered"])
}

func TestPhoneCheck_AlreadyRegistered(t *testing.T) {
	ts := testutil.NewTestServer(t)
	phoneNum := "+8613900138998"
	ts.SetSMSCode(t, phoneNum, "123456")

	reg := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/auth/register", map[string]string{
		"phone":         phoneNum,
		"code":          "123456",
		"register_type": "new_sns",
		"sns_name":      "phone_check_taken.seedao",
	}, nil)
	require.Equal(t, 0, testutil.ParseResponse(t, reg).Code)

	rec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/v1/auth/phone/check?phone=%2B8613900138998", nil, nil)
	require.Equal(t, http.StatusBadRequest, rec.Code)
	resp := testutil.ParseErrorResponse(t, rec)
	assert.Equal(t, errcode.PhoneAlreadyExists, resp.Code)
}

func TestPhoneCheck_InvalidPhone(t *testing.T) {
	ts := testutil.NewTestServer(t)

	rec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/v1/auth/phone/check?phone=13800138999", nil, nil)
	require.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestSNSCheck_Available(t *testing.T) {
	ts := testutil.NewTestServer(t)

	rec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/v1/sns/check?sns_name=alice_test_123.seedao", nil, nil)
	resp := testutil.ParseResponse(t, rec)

	data := testutil.ParseData[map[string]interface{}](t, resp.Data)
	assert.Equal(t, "alice_test_123.seedao", data["sns_name"])
	assert.Equal(t, true, data["available"])
}

func TestSNSCheck_InvalidFormatWithoutSuffix(t *testing.T) {
	ts := testutil.NewTestServer(t)

	rec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/v1/sns/check?sns_name=alice_test_123", nil, nil)
	resp := testutil.ParseResponse(t, rec)

	data := testutil.ParseData[map[string]interface{}](t, resp.Data)
	assert.Equal(t, false, data["available"])
	assert.Equal(t, "invalid_format", data["reason"])
}

func TestSNSCheck_TakenInRegistry(t *testing.T) {
	ts := testutil.NewTestServer(t)
	ts.SeedRegistry(t, "0x8c913aec7443fe2018639133398955e0e17fb0c1", "baiyu.seedao")

	rec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/v1/sns/check?sns_name=baiyu.seedao", nil, nil)
	resp := testutil.ParseResponse(t, rec)

	data := testutil.ParseData[map[string]interface{}](t, resp.Data)
	assert.Equal(t, false, data["available"])
	assert.Equal(t, "taken_in_registry", data["reason"])
}

func TestRegisterNewSNS(t *testing.T) {
	ts := testutil.NewTestServer(t)
	phoneNum := "+8613900139000"
	ts.SetSMSCode(t, phoneNum, "123456")

	rec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/auth/register", map[string]string{
		"phone":         phoneNum,
		"code":          "123456",
		"register_type": "new_sns",
		"sns_name":      "alice_test.seedao",
	}, nil)
	resp := testutil.ParseResponse(t, rec)

	data := testutil.ParseData[map[string]interface{}](t, resp.Data)
	assert.NotEmpty(t, data["token"])
	user := data["user"].(map[string]interface{})
	assert.Equal(t, "alice_test.seedao", user["sns_name"])
}

func TestRegisterNewSNS_InvalidSNSFormat(t *testing.T) {
	ts := testutil.NewTestServer(t)
	phoneNum := "+8613900139009"
	ts.SetSMSCode(t, phoneNum, "123456")

	rec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/auth/register", map[string]string{
		"phone":         phoneNum,
		"code":          "123456",
		"register_type": "new_sns",
		"sns_name":      "alice_test",
	}, nil)
	require.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestRegisterNewSNS_UnavailableSNSDoesNotConsumeCode(t *testing.T) {
	ts := testutil.NewTestServer(t)
	ts.SeedRegistry(t, "0x8c913aec7443fe2018639133398955e0e17fb0c1", "young.seedao")

	phoneNum := "+8613900139010"
	ts.SetSMSCode(t, phoneNum, "123456")

	failRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/auth/register", map[string]string{
		"phone":         phoneNum,
		"code":          "123456",
		"register_type": "new_sns",
		"sns_name":      "young.seedao",
	}, nil)
	require.Equal(t, http.StatusBadRequest, failRec.Code)

	okRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/auth/register", map[string]string{
		"phone":         phoneNum,
		"code":          "123456",
		"register_type": "new_sns",
		"sns_name":      "young_ok.seedao",
	}, nil)
	resp := testutil.ParseResponse(t, okRec)

	data := testutil.ParseData[map[string]interface{}](t, resp.Data)
	assert.NotEmpty(t, data["token"])
}

func TestLogin(t *testing.T) {
	ts := testutil.NewTestServer(t)
	phoneNum := "+8613900139001"
	ts.SetSMSCode(t, phoneNum, "123456")

	reg := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/auth/register", map[string]string{
		"phone":         phoneNum,
		"code":          "123456",
		"register_type": "new_sns",
		"sns_name":      "bob_test.seedao",
	}, nil)
	require.Equal(t, http.StatusOK, reg.Code)

	ts.SetSMSCode(t, phoneNum, "654321")
	rec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/auth/login", map[string]string{
		"phone": phoneNum,
		"code":  "654321",
	}, nil)
	resp := testutil.ParseResponse(t, rec)

	data := testutil.ParseData[map[string]interface{}](t, resp.Data)
	assert.NotEmpty(t, data["token"])
}

func TestWalletNonce(t *testing.T) {
	ts := testutil.NewTestServer(t)

	rec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/wallet/nonce", map[string]string{
		"wallet": "0x8c913aec7443fe2018639133398955e0e17fb0c1",
		"domain": "app.seedao.top",
	}, nil)
	resp := testutil.ParseResponse(t, rec)

	data := testutil.ParseData[map[string]interface{}](t, resp.Data)
	assert.NotEmpty(t, data["nonce"])
	assert.NotEmpty(t, data["message"])
}

func TestWalletVerifySNS_InvalidSignature(t *testing.T) {
	ts := testutil.NewTestServer(t)
	ts.SeedRegistry(t, "0x8c913aec7443fe2018639133398955e0e17fb0c1", "baiyu.seedao")

	wallet := "0x8c913aec7443fe2018639133398955e0e17fb0c1"
	nonceRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/wallet/nonce", map[string]string{
		"wallet": wallet,
		"domain": "app.seedao.top",
	}, nil)
	nonceResp := testutil.ParseResponse(t, nonceRec)
	nonceData := testutil.ParseData[map[string]interface{}](t, nonceResp.Data)

	rec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/wallet/verify-sns", map[string]string{
		"wallet":      wallet,
		"wallet_type": "EOA",
		"domain":      "app.seedao.top",
		"message":     nonceData["message"].(string),
		"signature":   "0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
	}, nil)
	require.Equal(t, http.StatusBadRequest, rec.Code)
	resp := testutil.ParseErrorResponse(t, rec)
	assert.Equal(t, errcode.WalletVerifyFailed, resp.Code)
	assert.Equal(t, service.ReasonSignatureInvalid, resp.Detail)
}

func TestWalletVerifySNS_NoNonce(t *testing.T) {
	ts := testutil.NewTestServer(t)
	wallet := "0x8c913aec7443fe2018639133398955e0e17fb0c1"

	rec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/wallet/verify-sns", map[string]string{
		"wallet":      wallet,
		"wallet_type": "EOA",
		"domain":      "app.seedao.top",
		"message":     "app.seedao.top wants you to sign in",
		"signature":   "0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
	}, nil)
	require.Equal(t, http.StatusBadRequest, rec.Code)
	resp := testutil.ParseErrorResponse(t, rec)
	assert.Equal(t, errcode.WalletVerifyFailed, resp.Code)
	assert.Equal(t, service.ReasonNonceExpiredOrMissing, resp.Detail)
}

func TestProfile(t *testing.T) {
	ts := testutil.NewTestServer(t)
	phoneNum := "+8613900139002"
	ts.SetSMSCode(t, phoneNum, "123456")

	reg := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/auth/register", map[string]string{
		"phone":         phoneNum,
		"code":          "123456",
		"register_type": "new_sns",
		"sns_name":      "carol_test.seedao",
	}, nil)
	regResp := testutil.ParseResponse(t, reg)
	regData := testutil.ParseData[map[string]interface{}](t, regResp.Data)
	token := regData["token"].(string)

	rec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/v1/user/profile", nil, map[string]string{
		"Authorization": "Bearer " + token,
	})
	resp := testutil.ParseResponse(t, rec)

	data := testutil.ParseData[map[string]interface{}](t, resp.Data)
	assert.Equal(t, "carol_test.seedao", data["sns_name"])
	assert.Contains(t, data["phone"].(string), "****")
}

func TestProfile_Unauthorized(t *testing.T) {
	ts := testutil.NewTestServer(t)

	rec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/v1/user/profile", nil, nil)
	require.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestStatsUserCount(t *testing.T) {
	ts := testutil.NewTestServer(t)
	ts.SeedRegistry(t, "0x8c913aec7443fe2018639133398955e0e17fb0c1", "legacy_one.seedao")
	ts.SeedRegistry(t, "0x1111111111111111111111111111111111111111", "legacy_two.seedao")

	phoneNum := "+8613900139011"
	ts.SetSMSCode(t, phoneNum, "123456")
	reg := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/auth/register", map[string]string{
		"phone":         phoneNum,
		"code":          "123456",
		"register_type": "new_sns",
		"sns_name":      "stats_test.seedao",
	}, nil)
	require.Equal(t, http.StatusOK, reg.Code)

	rec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/v1/stats/user-count", nil, nil)
	resp := testutil.ParseResponse(t, rec)

	data := testutil.ParseData[map[string]int64](t, resp.Data)
	assert.Equal(t, int64(3), data["total_user_count"])
}

func TestRegisterNewSNS_DefaultMembership(t *testing.T) {
	ts := testutil.NewTestServer(t)
	phoneNum := "+8613900139012"
	ts.SetSMSCode(t, phoneNum, "123456")

	rec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/auth/register", map[string]string{
		"phone":         phoneNum,
		"code":          "123456",
		"register_type": "new_sns",
		"sns_name":      "member_test.seedao",
	}, nil)
	resp := testutil.ParseResponse(t, rec)

	data := testutil.ParseData[map[string]interface{}](t, resp.Data)
	user := data["user"].(map[string]interface{})
	assert.NotEmpty(t, user["membership_expires_at"])
}

func TestMembershipRenew(t *testing.T) {
	ts := testutil.NewTestServer(t)
	phoneNum := "+8613900139013"
	ts.SetSMSCode(t, phoneNum, "123456")

	reg := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/auth/register", map[string]string{
		"phone":         phoneNum,
		"code":          "123456",
		"register_type": "new_sns",
		"sns_name":      "renew_test.seedao",
	}, nil)
	regResp := testutil.ParseResponse(t, reg)
	regData := testutil.ParseData[map[string]interface{}](t, regResp.Data)
	token := regData["token"].(string)

	rec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/user/membership/renew", map[string]int{
		"days": 30,
	}, map[string]string{
		"Authorization": "Bearer " + token,
	})
	resp := testutil.ParseResponse(t, rec)

	data := testutil.ParseData[map[string]interface{}](t, resp.Data)
	assert.Equal(t, float64(30), data["days_added"])
	assert.NotEmpty(t, data["membership_expires_at"])
}

func TestMembershipRenew_DefaultDays(t *testing.T) {
	ts := testutil.NewTestServer(t)
	phoneNum := "+8613900139014"
	ts.SetSMSCode(t, phoneNum, "123456")

	reg := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/auth/register", map[string]string{
		"phone":         phoneNum,
		"code":          "123456",
		"register_type": "new_sns",
		"sns_name":      "renew_default.seedao",
	}, nil)
	regResp := testutil.ParseResponse(t, reg)
	regData := testutil.ParseData[map[string]interface{}](t, regResp.Data)
	token := regData["token"].(string)

	rec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/user/membership/renew", map[string]string{}, map[string]string{
		"Authorization": "Bearer " + token,
	})
	resp := testutil.ParseResponse(t, rec)

	data := testutil.ParseData[map[string]interface{}](t, resp.Data)
	assert.Equal(t, float64(365), data["days_added"])
}

func registerAndGetToken(t *testing.T, ts *testutil.TestServer, phoneNum, snsName string) string {
	t.Helper()
	ts.SetSMSCode(t, phoneNum, "123456")
	rec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/auth/register", map[string]string{
		"phone":         phoneNum,
		"code":          "123456",
		"register_type": "new_sns",
		"sns_name":      snsName,
	}, nil)
	resp := testutil.ParseResponse(t, rec)
	data := testutil.ParseData[map[string]interface{}](t, resp.Data)
	return data["token"].(string)
}

func TestPaymentCreateOrder(t *testing.T) {
	ts := testutil.NewTestServer(t)
	token := registerAndGetToken(t, ts, "+8613900139015", "pay_create.seedao")

	rec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/payments/orders", map[string]interface{}{
		"product_type": "membership_renew",
		"days":         365,
	}, map[string]string{
		"Authorization": "Bearer " + token,
	})
	resp := testutil.ParseResponse(t, rec)

	data := testutil.ParseData[map[string]interface{}](t, resp.Data)
	assert.NotEmpty(t, data["order_no"])
	assert.Equal(t, "pending", data["status"])
	assert.NotEmpty(t, data["code_url"])
	assert.Equal(t, float64(9900), data["amount"])
}

func TestPaymentCreateOrderH5(t *testing.T) {
	ts := testutil.NewTestServer(t)
	token := registerAndGetToken(t, ts, "+8613900139018", "pay_h5.seedao")

	rec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/payments/orders", map[string]interface{}{
		"product_type": "membership_renew",
		"pay_channel":  "wechat_h5",
	}, map[string]string{
		"Authorization": "Bearer " + token,
	})
	resp := testutil.ParseResponse(t, rec)

	data := testutil.ParseData[map[string]interface{}](t, resp.Data)
	assert.NotEmpty(t, data["order_no"])
	assert.Equal(t, "wechat_h5", data["pay_channel"])
	assert.Equal(t, "pending", data["status"])
	assert.NotEmpty(t, data["h5_url"])
	assert.Empty(t, data["code_url"])
}

func TestPaymentCreateOrderInvalidChannel(t *testing.T) {
	ts := testutil.NewTestServer(t)
	token := registerAndGetToken(t, ts, "+8613900139019", "pay_bad.seedao")

	rec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/payments/orders", map[string]interface{}{
		"pay_channel": "wechat_jsapi",
	}, map[string]string{
		"Authorization": "Bearer " + token,
	})
	require.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestPaymentSimulatePayDevH5(t *testing.T) {
	ts := testutil.NewTestServer(t)
	token := registerAndGetToken(t, ts, "+8613900139020", "pay_h5_sim.seedao")

	createRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/payments/orders", map[string]interface{}{
		"pay_channel": "wechat_h5",
	}, map[string]string{
		"Authorization": "Bearer " + token,
	})
	createResp := testutil.ParseResponse(t, createRec)
	createData := testutil.ParseData[map[string]interface{}](t, createResp.Data)
	orderNo := createData["order_no"].(string)

	simRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/payments/dev/simulate-pay", map[string]string{
		"order_no": orderNo,
	}, map[string]string{
		"Authorization": "Bearer " + token,
	})
	simResp := testutil.ParseResponse(t, simRec)
	simData := testutil.ParseData[map[string]interface{}](t, simResp.Data)
	assert.Equal(t, "paid", simData["status"])
	assert.Equal(t, "wechat_h5", simData["pay_channel"])
}

func TestPaymentSimulatePayDev(t *testing.T) {
	ts := testutil.NewTestServer(t)
	token := registerAndGetToken(t, ts, "+8613900139016", "pay_sim.seedao")

	createRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/payments/orders", map[string]interface{}{
		"product_type": "membership_renew",
	}, map[string]string{
		"Authorization": "Bearer " + token,
	})
	createResp := testutil.ParseResponse(t, createRec)
	createData := testutil.ParseData[map[string]interface{}](t, createResp.Data)
	orderNo := createData["order_no"].(string)

	simRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/payments/dev/simulate-pay", map[string]string{
		"order_no": orderNo,
	}, map[string]string{
		"Authorization": "Bearer " + token,
	})
	simResp := testutil.ParseResponse(t, simRec)
	simData := testutil.ParseData[map[string]interface{}](t, simResp.Data)
	assert.Equal(t, "paid", simData["status"])
	assert.NotEmpty(t, simData["membership_expires_at"])

	getRec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/v1/payments/orders/"+orderNo, nil, map[string]string{
		"Authorization": "Bearer " + token,
	})
	getResp := testutil.ParseResponse(t, getRec)
	getData := testutil.ParseData[map[string]interface{}](t, getResp.Data)
	assert.Equal(t, "paid", getData["status"])
	assert.NotEmpty(t, getData["membership_expires_at"])
}

func TestPaymentCloseOrder(t *testing.T) {
	ts := testutil.NewTestServer(t)
	token := registerAndGetToken(t, ts, "+8613900139017", "pay_close.seedao")

	createRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/payments/orders", map[string]interface{}{
		"product_type": "membership_renew",
	}, map[string]string{
		"Authorization": "Bearer " + token,
	})
	createResp := testutil.ParseResponse(t, createRec)
	createData := testutil.ParseData[map[string]interface{}](t, createResp.Data)
	orderNo := createData["order_no"].(string)

	closeRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/payments/orders/"+orderNo+"/close", nil, map[string]string{
		"Authorization": "Bearer " + token,
	})
	closeResp := testutil.ParseResponse(t, closeRec)
	closeData := testutil.ParseData[map[string]interface{}](t, closeResp.Data)
	assert.Equal(t, "closed", closeData["status"])
}

func TestCouponPreviewWithoutJWT(t *testing.T) {
	ts := testutil.NewTestServer(t)
	ts.SeedMembershipCoupon(t, "SAVE10")

	rec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/payments/coupons/preview", map[string]interface{}{
		"coupon_code": "SAVE10",
		"days":        365,
	}, nil)
	resp := testutil.ParseResponse(t, rec)
	data := testutil.ParseData[map[string]interface{}](t, resp.Data)
	assert.Equal(t, float64(9900), data["original_amount"])
	assert.Equal(t, float64(8900), data["final_amount"])
}

func TestCouponPreviewWithJWT(t *testing.T) {
	ts := testutil.NewTestServer(t)
	ts.SeedMembershipCoupon(t, "SAVE10")
	token := registerAndGetToken(t, ts, "+8613900139021", "coupon_preview.seedao")

	rec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/payments/coupons/preview", map[string]interface{}{
		"coupon_code": "SAVE10",
	}, map[string]string{
		"Authorization": "Bearer " + token,
	})
	resp := testutil.ParseResponse(t, rec)
	data := testutil.ParseData[map[string]interface{}](t, resp.Data)
	assert.Equal(t, float64(1000), data["discount_amount"])
}

func TestCoursePreviewAnonymousShowsMemberPrice(t *testing.T) {
	ts := testutil.NewTestServer(t)
	ts.SeedCourseCatalog(t)

	rec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/courses/checkout/preview", map[string]interface{}{
		"course_id":    "course_test_spring",
		"sku_id":       "standard",
		"channel_code": "PARTNER_X",
	}, nil)
	resp := testutil.ParseResponse(t, rec)
	data := testutil.ParseData[map[string]interface{}](t, resp.Data)

	assert.Equal(t, float64(99900), data["original_amount"])
	assert.Equal(t, float64(74900), data["final_amount"])
	assert.Equal(t, float64(64900), data["amount_if_member"])

	member := data["member_discount"].(map[string]interface{})
	assert.Equal(t, false, member["applied"])
	assert.Equal(t, true, member["requires_login"])
	assert.Contains(t, member["hint"], "登录后可享")
}

func TestCourseAnonymousOrderAndSimulatePay(t *testing.T) {
	ts := testutil.NewTestServer(t)
	ts.SeedCourseCatalog(t)

	createRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/courses/orders", map[string]interface{}{
		"course_id":    "course_test_spring",
		"sku_id":       "standard",
		"channel_code": "PARTNER_X",
		"pay_channel":  "wechat_native",
	}, nil)
	createResp := testutil.ParseResponse(t, createRec)
	createData := testutil.ParseData[map[string]interface{}](t, createResp.Data)

	orderNo := createData["order_no"].(string)
	orderToken := createData["order_token"].(string)
	assert.NotEmpty(t, orderToken)
	assert.Equal(t, float64(74900), createData["amount"])

	simRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/courses/dev/orders/"+orderNo+"/simulate-pay", map[string]string{
		"order_token": orderToken,
	}, nil)
	simResp := testutil.ParseResponse(t, simRec)
	simData := testutil.ParseData[map[string]interface{}](t, simResp.Data)
	assert.Equal(t, "paid", simData["status"])
	assert.Equal(t, "active", simData["enrollment_status"])

	getRec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/v1/courses/orders/"+orderNo+"?token="+orderToken, nil, nil)
	getResp := testutil.ParseResponse(t, getRec)
	getData := testutil.ParseData[map[string]interface{}](t, getResp.Data)
	assert.Equal(t, "paid", getData["status"])
}

func TestCourseClaimOrder(t *testing.T) {
	ts := testutil.NewTestServer(t)
	ts.SeedCourseCatalog(t)

	createRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/courses/orders", map[string]interface{}{
		"course_id":    "course_test_spring",
		"sku_id":       "standard",
		"channel_code": "PARTNER_X",
	}, nil)
	createResp := testutil.ParseResponse(t, createRec)
	createData := testutil.ParseData[map[string]interface{}](t, createResp.Data)
	orderNo := createData["order_no"].(string)
	orderToken := createData["order_token"].(string)

	token := registerAndGetToken(t, ts, "+8613900139022", "course_claim.seedao")

	claimRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/courses/orders/claim", map[string]string{
		"order_no":     orderNo,
		"order_token":  orderToken,
	}, map[string]string{
		"Authorization": "Bearer " + token,
	})
	claimResp := testutil.ParseResponse(t, claimRec)
	assert.Equal(t, orderNo, testutil.ParseData[map[string]interface{}](t, claimResp.Data)["order_no"])

	getRec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/v1/courses/orders/"+orderNo, nil, map[string]string{
		"Authorization": "Bearer " + token,
	})
	getResp := testutil.ParseResponse(t, getRec)
	assert.Equal(t, "pending", testutil.ParseData[map[string]interface{}](t, getResp.Data)["status"])
}

func TestPaymentCreateOrderRejectsCourseProduct(t *testing.T) {
	ts := testutil.NewTestServer(t)
	token := registerAndGetToken(t, ts, "+8613900139023", "pay_reject.seedao")

	rec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/payments/orders", map[string]interface{}{
		"product_type": "course_purchase",
	}, map[string]string{
		"Authorization": "Bearer " + token,
	})
	require.Equal(t, http.StatusBadRequest, rec.Code)
}

