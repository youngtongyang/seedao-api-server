package handler_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/seedao/seedao-api-server/internal/pkg/errcode"
	"github.com/seedao/seedao-api-server/internal/testutil"
)

func TestProposalCreateTownHallWirePayload(t *testing.T) {
	ts := testutil.NewTestServer(t)
	token := ts.RegisterAndToken(t, "+8613900200999", "hall.seedao")

	body := map[string]interface{}{
		"title":           "ss",
		"type":            "town_hall_meeting",
		"resolution_rule": "town_hall_standard",
		"vote_type":       1,
		"content": []int{
			123, 34, 98, 97, 99, 107, 103, 114, 111, 117, 110, 100, 34, 58, 34, 115, 115, 92, 110, 34, 44, 34, 98, 111, 100, 121, 34, 58, 34, 115, 115, 92, 110, 34, 44, 34, 114, 101, 109, 97, 114, 107, 115, 34, 58, 34, 115, 115, 92, 110, 34, 44, 34, 118, 111, 116, 101, 95, 116, 121, 112, 101, 34, 58, 49, 44, 34, 118, 111, 116, 101, 95, 111, 112, 116, 105, 111, 110, 115, 34, 58, 91, 34, 229, 144, 140, 230, 132, 143, 34, 44, 34, 229, 143, 141, 229, 175, 185, 34, 44, 34, 229, 188, 131, 230, 157, 131, 34, 93, 125,
		},
		"vote_options": []int{
			91, 34, 229, 144, 140, 230, 132, 143, 34, 44, 34, 229, 143, 141, 229, 175, 185, 34, 44, 34, 229, 188, 131, 230, 157, 131, 34, 93,
		},
	}

	rec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/proposals", body, map[string]string{
		"Authorization": "Bearer " + token,
	})
	resp := testutil.ParseResponse(t, rec)
	t.Logf("status=%d code=%d detail=%q data=%v", rec.Code, resp.Code, resp.Detail, resp.Data)
	require.Equal(t, errcode.OK, resp.Code, "expected success, got detail=%q", resp.Detail)

	data := testutil.ParseData[map[string]interface{}](t, resp.Data)
	assert.Equal(t, "meta_rule", data["resolution_rule"])
}
