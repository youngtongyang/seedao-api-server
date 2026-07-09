package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/seedao/seedao-api-server/internal/config"
	"github.com/seedao/seedao-api-server/internal/model"
)

func TestEvaluateResolutionStandard(t *testing.T) {
	tally := VoteTally{
		TotalVotes:   12,
		ApproveCount: 8,
		RejectCount:  3,
		AbstainCount: 1,
		ByResult:     map[string]int{},
	}
	out := evaluateResolution(model.ProposalResolutionStandard, 20, tally)
	assert.True(t, out.Resolved)
	assert.True(t, out.Passed)
}

func TestEvaluateResolutionStandardFailsParticipation(t *testing.T) {
	tally := VoteTally{TotalVotes: 8, ApproveCount: 7, RejectCount: 1, ByResult: map[string]int{}}
	out := evaluateResolution(model.ProposalResolutionStandard, 20, tally)
	assert.True(t, out.Resolved)
	assert.False(t, out.Passed)
}

func TestEvaluateResolutionMetaRule(t *testing.T) {
	tally := VoteTally{TotalVotes: 15, ApproveCount: 14, ByResult: map[string]int{}}
	out := evaluateResolution(model.ProposalResolutionMetaRule, 20, tally)
	assert.True(t, out.Resolved)
	assert.True(t, out.Passed)
}

func TestEvaluateResolutionMetaRuleFails(t *testing.T) {
	tally := VoteTally{TotalVotes: 15, ApproveCount: 12, ByResult: map[string]int{}}
	out := evaluateResolution(model.ProposalResolutionMetaRule, 20, tally)
	assert.True(t, out.Resolved)
	assert.False(t, out.Passed)
}

func TestEvaluateResolutionMetaRuleHallFourMembers(t *testing.T) {
	tally := VoteTally{TotalVotes: 4, ApproveCount: 3, ByResult: map[string]int{}}
	out := evaluateResolution(model.ProposalResolutionMetaRule, 4, tally)
	assert.True(t, out.Resolved)
	assert.True(t, out.Passed)

	tally.ApproveCount = 2
	out = evaluateResolution(model.ProposalResolutionMetaRule, 4, tally)
	assert.False(t, out.Passed)
}

func TestTwoThirdsRequired(t *testing.T) {
	assert.Equal(t, 9, twoThirdsRequired(13))
	assert.Equal(t, 3, twoThirdsRequired(4))
	assert.Equal(t, 14, twoThirdsRequired(20))
}

func TestGovernanceACLCanCreate(t *testing.T) {
	acl := newGovernanceACL(configWithGovernance())
	assert.True(t, acl.canCreateType("baiyu.seedao", model.ProposalTypeHallRegular))
	assert.False(t, acl.canCreateType("other.seedao", model.ProposalTypeHallRegular))
}

func configWithGovernance() config.GovernanceConfig {
	return config.GovernanceConfig{
		HallProposalSNS:          []string{"baiyu.seedao"},
		NodeConsensusProposalSNS: []string{"node.seedao"},
		NodeVoterSNS:             []string{"voter.seedao"},
	}
}
