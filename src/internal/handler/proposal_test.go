package handler_test

import (
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/seedao/seedao-api-server/internal/model"
	"github.com/seedao/seedao-api-server/internal/pkg/errcode"
	"github.com/seedao/seedao-api-server/internal/testutil"
)

func TestProposalListAndGovernanceNodes(t *testing.T) {
	ts := testutil.NewTestServer(t)
	_ = ts.SeedPublishedProposal(t, "hall.seedao", model.ProposalTypeHallRegular, false)

	listRec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/v1/proposals?limit=10", nil, nil)
	listResp := testutil.ParseResponse(t, listRec)
	listData := testutil.ParseData[map[string]interface{}](t, listResp.Data)
	items := listData["items"].([]interface{})
	assert.Len(t, items, 1)

	nodesRec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/v1/governance/nodes", nil, nil)
	nodesResp := testutil.ParseResponse(t, nodesRec)
	nodesData := testutil.ParseData[map[string]interface{}](t, nodesResp.Data)
	assert.Equal(t, float64(4), nodesData["quarter_node_total"])
}

func TestProposalCreatePublishVoteFlow(t *testing.T) {
	ts := testutil.NewTestServer(t)
	token := ts.RegisterAndToken(t, "+8613900200001", "node.seedao")

	createRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/proposals", map[string]interface{}{
		"type":    model.ProposalTypeNodeConsensus,
		"title":   "Node proposal",
		"content": map[string]string{"plain_text": "body"},
	}, map[string]string{"Authorization": "Bearer " + token})
	createResp := testutil.ParseResponse(t, createRec)
	created := testutil.ParseData[map[string]interface{}](t, createResp.Data)
	id := int64(created["id"].(float64))

	pubRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/proposals/"+pid(id)+"/publish", map[string]interface{}{
		"publicity_seconds":     0,
		"vote_duration_seconds": 7200,
	}, map[string]string{"Authorization": "Bearer " + token})
	testutil.ParseResponse(t, pubRec)

	// 确保进入投票期（SQLite 时间精度可能导致边界问题）
	require.NoError(t, ts.DB.Model(&model.Proposal{}).Where("id = ?", id).Updates(map[string]interface{}{
		"publicity_end_at": time.Now().Add(-time.Minute),
		"vote_end_at":      time.Now().Add(time.Hour),
		"status":           model.ProposalStatusPublicity,
	}).Error)

	voterToken := ts.RegisterAndToken(t, "+8613900200002", "node_voter.seedao")
	voteRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/proposals/"+pid(id)+"/votes", map[string]string{
		"vote_result": "1",
	}, map[string]string{"Authorization": "Bearer " + voterToken})
	voteResp := testutil.ParseResponse(t, voteRec)
	voteData := testutil.ParseData[map[string]interface{}](t, voteResp.Data)
	assert.Equal(t, "1", voteData["vote_result"])

	meRec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/v1/proposals/"+pid(id)+"/votes/me", nil, map[string]string{
		"Authorization": "Bearer " + voterToken,
	})
	meResp := testutil.ParseResponse(t, meRec)
	meData := testutil.ParseData[map[string]interface{}](t, meResp.Data)
	assert.Equal(t, "node_voter.seedao", meData["sns_name"])

	summaryRec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/v1/proposals/"+pid(id)+"/vote-summary", nil, nil)
	summaryResp := testutil.ParseResponse(t, summaryRec)
	summary := testutil.ParseData[map[string]interface{}](t, summaryResp.Data)
	assert.Equal(t, float64(1), summary["total_votes"])

	commentRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/proposals/"+pid(id)+"/comments", map[string]string{
		"content": "support",
	}, map[string]string{"Authorization": "Bearer " + voterToken})
	commentResp := testutil.ParseResponse(t, commentRec)
	commentData := testutil.ParseData[map[string]interface{}](t, commentResp.Data)
	assert.Equal(t, "support", commentData["content"])

	detailRec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/v1/proposals/"+pid(id)+"?include_votes=true&include_comments=true", nil, nil)
	detailResp := testutil.ParseResponse(t, detailRec)
	detail := testutil.ParseData[map[string]interface{}](t, detailResp.Data)
	assert.NotNil(t, detail["votes"])
	assert.NotNil(t, detail["comments"])
}

func TestProposalCreatePermissionsEndpoint(t *testing.T) {
	ts := testutil.NewTestServer(t)

	hallToken := ts.RegisterAndToken(t, "+8613900200101", "hall.seedao")
	hallRec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/v1/governance/proposal-permissions", nil, map[string]string{
		"Authorization": "Bearer " + hallToken,
	})
	hallResp := testutil.ParseResponse(t, hallRec)
	hallData := testutil.ParseData[map[string]interface{}](t, hallResp.Data)
	assert.Equal(t, "hall.seedao", hallData["sns_name"])
	assert.Equal(t, true, hallData["can_create_any"])
	createTypes := hallData["create_types"].(map[string]interface{})
	assert.Equal(t, true, createTypes["town_hall_meeting"])
	assert.Equal(t, false, createTypes["node_consensus_special"])

	randomToken := ts.RegisterAndToken(t, "+8613900200102", "random.seedao")
	randomRec := testutil.DoJSON(t, ts.Echo, http.MethodGet, "/v1/governance/proposal-permissions", nil, map[string]string{
		"Authorization": "Bearer " + randomToken,
	})
	randomResp := testutil.ParseResponse(t, randomRec)
	randomData := testutil.ParseData[map[string]interface{}](t, randomResp.Data)
	assert.Equal(t, false, randomData["can_create_any"])
}

func TestProposalUpdateDraftEmptyBody(t *testing.T) {
	ts := testutil.NewTestServer(t)
	token := ts.RegisterAndToken(t, "+8613900200201", "hall.seedao")

	createRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/proposals", map[string]interface{}{
		"type":    model.ProposalTypeHallRegular,
		"title":   "draft",
		"content": map[string]string{"background": "x"},
	}, map[string]string{"Authorization": "Bearer " + token})
	created := testutil.ParseData[map[string]interface{}](t, testutil.ParseResponse(t, createRec).Data)
	id := int64(created["id"].(float64))

	patchRec := testutil.DoJSON(t, ts.Echo, http.MethodPatch, "/v1/proposals/"+pid(id), map[string]interface{}{}, map[string]string{
		"Authorization": "Bearer " + token,
	})
	resp := testutil.ParseErrorResponse(t, patchRec)
	assert.Equal(t, errcode.InvalidParam, resp.Code)
	assert.Contains(t, resp.Detail, "no updatable fields provided")
}

func TestProposalCreateForbidden(t *testing.T) {
	ts := testutil.NewTestServer(t)
	token := ts.RegisterAndToken(t, "+8613900200003", "random.seedao")

	rec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/proposals", map[string]interface{}{
		"type":    model.ProposalTypeHallRegular,
		"title":   "nope",
		"content": map[string]string{"plain_text": "x"},
	}, map[string]string{"Authorization": "Bearer " + token})
	resp := testutil.ParseErrorResponse(t, rec)
	assert.Equal(t, errcode.ProposalForbidden, resp.Code)
}

func TestProposalVoteNotVoter(t *testing.T) {
	ts := testutil.NewTestServer(t)
	id := ts.SeedPublishedProposal(t, "hall.seedao", model.ProposalTypeHallRegular, true)
	token := ts.RegisterAndToken(t, "+8613900200004", "random.seedao")

	rec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/proposals/"+pid(id)+"/votes", map[string]string{
		"vote_result": "1",
	}, map[string]string{"Authorization": "Bearer " + token})
	resp := testutil.ParseErrorResponse(t, rec)
	assert.Equal(t, errcode.ProposalNotVoter, resp.Code)
}

func TestProposalCloseDraft(t *testing.T) {
	ts := testutil.NewTestServer(t)
	token := ts.RegisterAndToken(t, "+8613900200005", "hall.seedao")

	createRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/proposals", map[string]interface{}{
		"type":    model.ProposalTypeHallRegular,
		"title":   "to close",
		"content": map[string]string{"plain_text": "x"},
	}, map[string]string{"Authorization": "Bearer " + token})
	created := testutil.ParseData[map[string]interface{}](t, testutil.ParseResponse(t, createRec).Data)
	id := int64(created["id"].(float64))

	closeRec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/proposals/"+pid(id)+"/close", nil, map[string]string{
		"Authorization": "Bearer " + token,
	})
	closeResp := testutil.ParseResponse(t, closeRec)
	closeData := testutil.ParseData[map[string]interface{}](t, closeResp.Data)
	assert.Equal(t, model.ProposalStatusClosed, closeData["status"])
}

func TestProposalFinalize(t *testing.T) {
	ts := testutil.NewTestServer(t)
	id := ts.SeedPublishedProposal(t, "node.seedao", model.ProposalTypeNodeConsensus, false)

	now := time.Now()
	for _, sns := range []string{"node_voter.seedao", "node_voter2.seedao", "node_voter3.seedao"} {
		name := sns
		require.NoError(t, ts.DB.Create(&model.ProposalVote{
			ProposalID:  id,
			SNSName:     &name,
			VoterWallet: "0x1",
			VoteResult:  "1",
			VoteWeight:  1,
			VotedAt:     now,
		}).Error)
	}

	token := ts.RegisterAndToken(t, "+8613900200006", "node.seedao")
	rec := testutil.DoJSON(t, ts.Echo, http.MethodPost, "/v1/proposals/"+pid(id)+"/finalize", nil, map[string]string{
		"Authorization": "Bearer " + token,
	})
	resp := testutil.ParseResponse(t, rec)
	data := testutil.ParseData[map[string]interface{}](t, resp.Data)
	assert.Equal(t, model.ProposalStatusVotePassed, data["status"])
}

func pid(id int64) string {
	return strconv.FormatInt(id, 10)
}
