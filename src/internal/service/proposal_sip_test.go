package service

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/seedao/seedao-api-server/internal/model"
)

func TestAssignSIPAllocatesOnVotingOpen(t *testing.T) {
	svc := newProposalTestDB(t)
	now := time.Now()
	pubEnd := now.Add(-time.Minute)
	voteEnd := now.Add(time.Hour)
	sns := "node.seedao"

	require.NoError(t, svc.db.Create(&model.Proposal{
		Type: model.ProposalTypeNodeConsensus, Title: "assign sip", Content: []byte(`{}`),
		ProposerSNS: &sns, ProposerWallet: "0x1", Status: model.ProposalStatusPublicity,
		PublicityEndAt: &pubEnd, VoteEndAt: &voteEnd, ResolutionRule: model.ProposalResolutionStandard,
		CreatedAt: now, UpdatedAt: now,
	}).Error)
	var p model.Proposal
	require.NoError(t, svc.db.First(&p, "title = ?", "assign sip").Error)
	assert.Equal(t, 0, p.SIP)

	sip, err := svc.AssignSIP(p.ID)
	require.NoError(t, err)
	assert.Equal(t, 1, sip)

	require.NoError(t, svc.db.First(&p, p.ID).Error)
	assert.Equal(t, 1, p.SIP)

	sip2, err := svc.AssignSIP(p.ID)
	require.NoError(t, err)
	assert.Equal(t, 1, sip2)
}

func TestMaybeAssignSIPOnVotingOpenSkipsDraft(t *testing.T) {
	svc := newProposalTestDB(t)
	now := time.Now()
	sns := "node.seedao"
	require.NoError(t, svc.db.Create(&model.Proposal{
		Type: model.ProposalTypeNodeConsensus, Title: "draft no sip", Content: []byte(`{}`),
		ProposerSNS: &sns, ProposerWallet: "0x1", Status: model.ProposalStatusDraft,
		ResolutionRule: model.ProposalResolutionStandard, CreatedAt: now, UpdatedAt: now,
	}).Error)
	var p model.Proposal
	require.NoError(t, svc.db.First(&p, "title = ?", "draft no sip").Error)

	updated, err := svc.maybeAssignSIPOnVotingOpen(&p)
	require.NoError(t, err)
	assert.Equal(t, 0, updated.SIP)
}

func TestMaybeAssignSIPOnGetDetail(t *testing.T) {
	svc := newProposalTestDB(t)
	now := time.Now()
	pubEnd := now.Add(-time.Minute)
	voteEnd := now.Add(time.Hour)
	sns := "hall.seedao"

	require.NoError(t, svc.db.Create(&model.Proposal{
		Type: model.ProposalTypeHallRegular, Title: "detail sip", Content: []byte(`{}`),
		ProposerSNS: &sns, ProposerWallet: "0x1", Status: model.ProposalStatusPublicity,
		PublicityEndAt: &pubEnd, VoteEndAt: &voteEnd, ResolutionRule: model.ProposalResolutionStandard,
		CreatedAt: now, UpdatedAt: now,
	}).Error)
	var p model.Proposal
	require.NoError(t, svc.db.First(&p, "title = ?", "detail sip").Error)

	detail, err := svc.GetDetail(p.ID, false, false, 0, 0)
	require.NoError(t, err)
	assert.Equal(t, 1, detail.SIP)
	assert.Contains(t, detail.Title, "SIP-1:")
}

func TestAssignSIPInheritsCloseProject(t *testing.T) {
	svc := newProposalTestDB(t)
	now := time.Now()
	sns := "node.seedao"
	parentID := int64(10)
	childID := int64(11)

	require.NoError(t, svc.db.Create(&model.Proposal{
		ID: parentID, Type: model.ProposalTypeNodeConsensus, Title: "parent", Content: []byte(`{}`),
		ProposerSNS: &sns, ProposerWallet: "0x1", Status: model.ProposalStatusVotePassed,
		SIP: 88, ResolutionRule: model.ProposalResolutionStandard, CreatedAt: now, UpdatedAt: now,
	}).Error)
	require.NoError(t, svc.db.Create(&model.Proposal{
		ID: childID, Type: "P2项目结项提案", Title: "close", Content: []byte(`{}`),
		ProposerSNS: &sns, ProposerWallet: "0x1", Status: model.ProposalStatusPublicity,
		AssociateProposalID: &parentID, ResolutionRule: model.ProposalResolutionStandard,
		CreatedAt: now, UpdatedAt: now,
	}).Error)

	sip, err := svc.AssignSIP(childID)
	require.NoError(t, err)
	assert.Equal(t, 88, sip)
}

func TestAssignSIPCloseProjectMissingAssociate(t *testing.T) {
	svc := newProposalTestDB(t)
	now := time.Now()
	sns := "node.seedao"
	require.NoError(t, svc.db.Create(&model.Proposal{
		Type: "P2项目结项提案", Title: "bad close", Content: []byte(`{}`),
		ProposerSNS: &sns, ProposerWallet: "0x1", Status: model.ProposalStatusPublicity,
		ResolutionRule: model.ProposalResolutionStandard, CreatedAt: now, UpdatedAt: now,
	}).Error)
	var p model.Proposal
	require.NoError(t, svc.db.First(&p, "title = ?", "bad close").Error)

	_, err := svc.AssignSIP(p.ID)
	require.Error(t, err)
	var paramErr *ProposalParamError
	require.ErrorAs(t, err, &paramErr)
}
