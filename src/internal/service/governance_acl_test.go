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
