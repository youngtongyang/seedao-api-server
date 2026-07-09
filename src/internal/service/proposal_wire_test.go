package service

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNormalizeCreateProposalType(t *testing.T) {
	t.Run("frontend alias", func(t *testing.T) {
		got, err := normalizeCreateProposalType("node_consensus_rule")
		require.NoError(t, err)
		assert.Equal(t, "节点共识大会规则提案", got)
	})
	t.Run("unsupported", func(t *testing.T) {
		_, err := normalizeCreateProposalType("unknown_type")
		require.Error(t, err)
		var paramErr *ProposalParamError
		require.ErrorAs(t, err, &paramErr)
		assert.Contains(t, paramErr.Reason, "unsupported proposal type")
	})
}

func TestNormalizeResolutionRule(t *testing.T) {
	got, err := normalizeResolutionRule("standard")
	require.NoError(t, err)
	assert.Equal(t, "standard", got)

	got, err = normalizeResolutionRule("town_hall_standard")
	require.NoError(t, err)
	assert.Equal(t, "standard", got)

	_, err = normalizeResolutionRule("bad_rule")
	require.Error(t, err)
	var paramErr *ProposalParamError
	require.ErrorAs(t, err, &paramErr)
	assert.Contains(t, paramErr.Reason, "unsupported resolution_rule")
}

func TestNormalizeProposalContentWireBytes(t *testing.T) {
	obj := map[string]interface{}{
		"background": "这是一个测试提案\n",
		"body":       "正文内容\n",
	}
	expected, err := json.Marshal(obj)
	require.NoError(t, err)

	wire := make([]int, len(expected))
	for i, b := range expected {
		wire[i] = int(b)
	}
	raw, err := json.Marshal(wire)
	require.NoError(t, err)

	got, err := NormalizeProposalContent(raw)
	require.NoError(t, err)
	assert.JSONEq(t, string(expected), string(got))
}

func TestNormalizeVoteOptionsWireBytes(t *testing.T) {
	raw := json.RawMessage(`[91,34,229,144,140,230,132,143,34,44,34,229,143,141,229,175,185,34,44,34,229,188,131,230,157,131,34,93]`)
	got, err := NormalizeVoteOptions(raw)
	require.NoError(t, err)

	var opts []map[string]string
	require.NoError(t, json.Unmarshal(got, &opts))
	require.Len(t, opts, 3)
	assert.Equal(t, "同意", opts[0]["label"])
	assert.Equal(t, "1", opts[0]["value"])
	assert.Equal(t, "反对", opts[1]["label"])
	assert.Equal(t, "-1", opts[1]["value"])
	assert.Equal(t, "弃权", opts[2]["label"])
	assert.Equal(t, "0", opts[2]["value"])
}

func TestCreateDraftFrontendPayload(t *testing.T) {
	svc := newProposalTestDB(t)

	contentObj := map[string]interface{}{
		"background":   "这是一个测试提案\n",
		"body":         "正文内容\n",
		"remarks":      "备注\n",
		"vote_type":    1,
		"vote_options": []string{"同意", "反对", "弃权"},
	}
	contentJSON, err := json.Marshal(contentObj)
	require.NoError(t, err)
	contentWire := make([]int, len(contentJSON))
	for i, b := range contentJSON {
		contentWire[i] = int(b)
	}
	contentRaw, err := json.Marshal(contentWire)
	require.NoError(t, err)

	voteOptionsRaw := json.RawMessage(`[91,34,229,144,140,230,132,143,34,44,34,229,143,141,229,175,185,34,44,34,229,188,131,230,157,131,34,93]`)

	created, err := svc.CreateDraft(CreateProposalInput{
		Type:           "node_consensus_rule",
		Title:          "测试新建提案",
		Content:        contentRaw,
		ResolutionRule: "standard",
		VoteType:       1,
		VoteOptions:    voteOptionsRaw,
		ActorSNS:       "node.seedao",
	})
	require.NoError(t, err)
	assert.Equal(t, "节点共识大会规则提案", created.Type)
	assert.Equal(t, int16(1), created.VoteType)
	require.NotEmpty(t, created.VoteOptions)

	var opts []map[string]string
	require.NoError(t, json.Unmarshal(created.VoteOptions, &opts))
	require.Len(t, opts, 3)
}

func TestResolveVoteTypeFromContent(t *testing.T) {
	content, err := json.Marshal(map[string]interface{}{
		"background": "bg",
		"vote_type":  2,
	})
	require.NoError(t, err)

	assert.Equal(t, int16(2), resolveVoteType(0, content))
	assert.Equal(t, int16(2), resolveVoteType(2, content))
	assert.Equal(t, int16(1), resolveVoteType(1, content))
}

func TestCreateDraftMultiVoteFromContent(t *testing.T) {
	svc := newProposalTestDB(t)
	contentObj := map[string]interface{}{
		"background": "bg",
		"body":       "body",
		"vote_type":  2,
		"vote_options": []string{"A", "B", "C"},
	}
	contentJSON, err := json.Marshal(contentObj)
	require.NoError(t, err)

	created, err := svc.CreateDraft(CreateProposalInput{
		Type:           "node_consensus_rule",
		Title:          "multi vote",
		Content:        contentJSON,
		ResolutionRule: "standard",
		VoteOptions:    json.RawMessage(`["A","B","C"]`),
		ActorSNS:       "node.seedao",
	})
	require.NoError(t, err)
	assert.Equal(t, int16(2), created.VoteType)
}
