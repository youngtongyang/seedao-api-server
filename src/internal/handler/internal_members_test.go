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

func TestInternalMembers_List(t *testing.T) {
	ts := testutil.NewTestServer(t)
	_ = ts.RegisterAndToken(t, "+8613800138010", "memberlist1.seedao")
	_ = ts.RegisterAndToken(t, "+8613800138011", "memberlist2.seedao")

	now := time.Now()
	expires := time.Date(2028, 12, 31, 15, 59, 59, 0, time.UTC)
	require.NoError(t, ts.DB.Model(&model.User{}).Where("id = ?", 1).Updates(map[string]interface{}{
		"membership_expires_at": expires,
		"updated_at":            now,
	}).Error)

	rec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/internal/v1/members?page=1&page_size=10", nil, map[string]string{
		"X-Service-Token": "test-service-token",
	})
	resp := testutil.ParseResponse(t, rec)
	data := testutil.ParseData[map[string]interface{}](t, resp.Data)

	assert.Equal(t, float64(1), data["page"])
	assert.Equal(t, float64(10), data["page_size"])
	assert.Equal(t, float64(2), data["total"])
	assert.Equal(t, "user_id", data["sort_by"])
	assert.Equal(t, "asc", data["order"])

	items := data["items"].([]interface{})
	require.Len(t, items, 2)

	// default: user_id asc
	first := items[0].(map[string]interface{})
	second := items[1].(map[string]interface{})
	assert.Equal(t, "memberlist1.seedao", first["sns_name"])
	assert.Equal(t, "memberlist2.seedao", second["sns_name"])
	assert.Equal(t, "2028-12-31T15:59:59Z", first["membership_expires_at"])
	assert.NotEmpty(t, first["registered_at"])
}

func TestInternalMembers_ListSortDesc(t *testing.T) {
	ts := testutil.NewTestServer(t)
	_ = ts.RegisterAndToken(t, "+8613800138015", "sortdesc1.seedao")
	_ = ts.RegisterAndToken(t, "+8613800138016", "sortdesc2.seedao")

	rec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/internal/v1/members?sort_by=user_id&order=desc", nil, map[string]string{
		"X-Service-Token": "test-service-token",
	})
	data := testutil.ParseData[map[string]interface{}](t, testutil.ParseResponse(t, rec).Data)
	assert.Equal(t, "user_id", data["sort_by"])
	assert.Equal(t, "desc", data["order"])
	items := data["items"].([]interface{})
	require.Len(t, items, 2)
	assert.Equal(t, "sortdesc2.seedao", items[0].(map[string]interface{})["sns_name"])
	assert.Equal(t, "sortdesc1.seedao", items[1].(map[string]interface{})["sns_name"])
}

func TestInternalMembers_ListUnauthorized(t *testing.T) {
	ts := testutil.NewTestServer(t)

	rec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/internal/v1/members", nil, nil)
	require.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestInternalMembers_ListPagination(t *testing.T) {
	ts := testutil.NewTestServer(t)
	_ = ts.RegisterAndToken(t, "+8613800138012", "pagea.seedao")
	_ = ts.RegisterAndToken(t, "+8613800138013", "pageb.seedao")
	_ = ts.RegisterAndToken(t, "+8613800138014", "pagec.seedao")

	rec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/internal/v1/members?page=2&page_size=2", nil, map[string]string{
		"X-Service-Token": "test-service-token",
	})
	data := testutil.ParseData[map[string]interface{}](t, testutil.ParseResponse(t, rec).Data)
	assert.Equal(t, float64(3), data["total"])
	assert.Equal(t, float64(2), data["page"])
	items := data["items"].([]interface{})
	require.Len(t, items, 1)
	// default asc: ids 1,2 on page1; id 3 on page2
	assert.Equal(t, "pagec.seedao", items[0].(map[string]interface{})["sns_name"])
}
