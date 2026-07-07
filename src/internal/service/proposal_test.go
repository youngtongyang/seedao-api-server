package service

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/seedao/seedao-api-server/internal/config"
	"github.com/seedao/seedao-api-server/internal/model"
)

func testGovernanceCfg() config.GovernanceConfig {
	return config.GovernanceConfig{
		QuarterNodeTotal:           4,
		HallProposalSNS:            []string{"hall.seedao"},
		NodeConsensusProposalSNS:   []string{"node.seedao"},
		HallVoterSNS:               []string{"hall_voter.seedao"},
		NodeVoterSNS:               []string{"node_voter.seedao"},
		DefaultPublicitySeconds:    60,
		DefaultVoteDurationSeconds: 120,
	}
}

func newProposalTestDB(t *testing.T) *ProposalService {
	t.Helper()
	db, err := model.NewSQLiteDB(":memory:")
	require.NoError(t, err)
	require.NoError(t, model.AutoMigrate(db))
	return NewProposalService(db, testGovernanceCfg())
}

func TestProposalServiceCreateUpdateClose(t *testing.T) {
	svc := newProposalTestDB(t)
	content, _ := json.Marshal(map[string]string{"plain_text": "draft"})

	created, err := svc.CreateDraft(CreateProposalInput{
		Type:     model.ProposalTypeHallRegular,
		Title:    "Hall draft",
		Content:  content,
		ActorSNS: "hall.seedao",
	})
	require.NoError(t, err)
	assert.Equal(t, model.ProposalStatusDraft, created.Status)

	title := "Updated title"
	updated, err := svc.UpdateDraft(created.ID, UpdateProposalInput{
		Title:    &title,
		ActorSNS: "hall.seedao",
	})
	require.NoError(t, err)
	assert.Equal(t, title, updated.Title)

	closed, err := svc.CloseDraft(created.ID, "hall.seedao")
	require.NoError(t, err)
	assert.Equal(t, model.ProposalStatusClosed, closed.Status)
}

func TestProposalServiceCreateForbidden(t *testing.T) {
	svc := newProposalTestDB(t)
	content, _ := json.Marshal(map[string]string{"x": "1"})
	_, err := svc.CreateDraft(CreateProposalInput{
		Type:     model.ProposalTypeHallRegular,
		Title:    "nope",
		Content:  content,
		ActorSNS: "other.seedao",
	})
	assert.ErrorIs(t, err, ErrProposalForbidden)
}

func TestProposalServicePublishAndVote(t *testing.T) {
	svc := newProposalTestDB(t)
	content, _ := json.Marshal(map[string]string{"plain_text": "go"})

	created, err := svc.CreateDraft(CreateProposalInput{
		Type:     model.ProposalTypeNodeConsensus,
		Title:    "Node draft",
		Content:  content,
		ActorSNS: "node.seedao",
	})
	require.NoError(t, err)

	pubSec := 0
	voteSec := 3600
	published, err := svc.Publish(created.ID, PublishProposalInput{
		PublicitySeconds:    &pubSec,
		VoteDurationSeconds: &voteSec,
		ActorSNS:            "node.seedao",
	})
	require.NoError(t, err)
	assert.Equal(t, model.ProposalStatusPublicity, published.Status)

	// force voting window open
	var row model.Proposal
	require.NoError(t, svc.db.First(&row, created.ID).Error)
	past := time.Now().Add(-time.Minute)
	future := time.Now().Add(time.Hour)
	require.NoError(t, svc.db.Model(&row).Updates(map[string]interface{}{
		"publicity_end_at": past,
		"vote_end_at":      future,
	}).Error)

	vote, err := svc.CastVote(CastVoteInput{
		ProposalID: created.ID,
		VoteResult: "1",
		ActorSNS:   "node_voter.seedao",
	})
	require.NoError(t, err)
	assert.Equal(t, "1", vote.VoteResult)

	_, err = svc.CastVote(CastVoteInput{
		ProposalID: created.ID,
		VoteResult: "1",
		ActorSNS:   "node_voter.seedao",
	})
	assert.ErrorIs(t, err, ErrProposalAlreadyVoted)
}

func TestProposalServiceCommentAndFinalize(t *testing.T) {
	svc := newProposalTestDB(t)
	now := time.Now()
	pubEnd := now.Add(-2 * time.Hour)
	voteEnd := now.Add(-time.Minute)
	sns := "node.seedao"
	p := model.Proposal{
		Type:           model.ProposalTypeNodeConsensus,
		Title:          "Finalize me",
		Content:        []byte(`{"plain_text":"x"}`),
		ProposerSNS:    &sns,
		ProposerWallet: "0x1",
		Status:         model.ProposalStatusPublicity,
		ResolutionRule: model.ProposalResolutionStandard,
		PublishedAt:    &now,
		PublicityEndAt: &pubEnd,
		VoteEndAt:      &voteEnd,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	require.NoError(t, svc.db.Create(&p).Error)

	voter := "node_voter.seedao"
	require.NoError(t, svc.db.Create(&model.ProposalVote{
		ProposalID:  p.ID,
		SNSName:     &voter,
		VoterWallet: "0x2",
		VoteResult:  "1",
		VoteWeight:  1,
		VotedAt:     now,
	}).Error)
	for _, extra := range []string{"node_voter2.seedao", "node_voter3.seedao"} {
		name := extra
		require.NoError(t, svc.db.Create(&model.ProposalVote{
			ProposalID:  p.ID,
			SNSName:     &name,
			VoterWallet: "0x3",
			VoteResult:  "1",
			VoteWeight:  1,
			VotedAt:     now,
		}).Error)
	}

	comment, err := svc.CreateComment(CreateCommentInput{
		ProposalID: p.ID,
		Content:    "looks good",
		ActorSNS:   "node_voter.seedao",
	})
	require.NoError(t, err)
	assert.Equal(t, "looks good", comment.Content)

	detail, err := svc.FinalizeVote(p.ID, "node.seedao")
	require.NoError(t, err)
	assert.Equal(t, model.ProposalStatusVotePassed, detail.Status)
	assert.True(t, detail.VoteSummary.Passed)
}

func TestProposalServiceListAndSummary(t *testing.T) {
	svc := newProposalTestDB(t)
	now := time.Now()
	sns := "legacy"
	require.NoError(t, svc.db.Create(&model.Proposal{
		Type:           model.ProposalTypeHallRegular,
		Title:          "Listed",
		Content:        []byte(`{}`),
		ProposerSNS:    &sns,
		ProposerWallet: "0x1",
		Status:         model.ProposalStatusVotePassed,
		ResolutionRule: model.ProposalResolutionStandard,
		CreatedAt:      now,
		UpdatedAt:      now,
	}).Error)

	list, err := svc.List(10, 0, "", "")
	require.NoError(t, err)
	assert.Len(t, list.Items, 1)

	summary, err := svc.GetVoteSummary(list.Items[0].ID)
	require.NoError(t, err)
	assert.Equal(t, 0, summary.TotalVotes)
}

func TestProposalListOrderByCreatedAt(t *testing.T) {
	svc := newProposalTestDB(t)
	sns := "legacy"
	older := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	newer := time.Date(2026, 6, 1, 0, 0, 0, 0, time.UTC)

	require.NoError(t, svc.db.Create(&model.Proposal{
		ID: 100, Type: model.ProposalTypeHallRegular, Title: "older", Content: []byte(`{}`),
		ProposerSNS: &sns, ProposerWallet: "0x1", Status: model.ProposalStatusVotePassed,
		ResolutionRule: model.ProposalResolutionStandard, CreatedAt: older, UpdatedAt: older,
	}).Error)
	require.NoError(t, svc.db.Create(&model.Proposal{
		ID: 1, Type: model.ProposalTypeHallRegular, Title: "newer", Content: []byte(`{}`),
		ProposerSNS: &sns, ProposerWallet: "0x1", Status: model.ProposalStatusVotePassed,
		ResolutionRule: model.ProposalResolutionStandard, CreatedAt: newer, UpdatedAt: newer,
	}).Error)

	list, err := svc.List(10, 0, "", "")
	require.NoError(t, err)
	require.Len(t, list.Items, 2)
	assert.Equal(t, "newer", list.Items[0].Title)
	assert.Equal(t, "older", list.Items[1].Title)

	page2, err := svc.List(10, list.Items[0].ID, "", "")
	require.NoError(t, err)
	require.Len(t, page2.Items, 1)
	assert.Equal(t, "older", page2.Items[0].Title)
}

func TestProposalListExcludesErrorStatus(t *testing.T) {
	svc := newProposalTestDB(t)
	sns := "legacy"
	now := time.Now()

	for _, p := range []model.Proposal{
		{Type: model.ProposalTypeHallRegular, Title: "ok", Content: []byte(`{}`), ProposerSNS: &sns, ProposerWallet: "0x1", Status: model.ProposalStatusVotePassed, ResolutionRule: model.ProposalResolutionStandard, CreatedAt: now, UpdatedAt: now},
		{Type: model.ProposalTypeHallRegular, Title: "bad", Content: []byte(`{}`), ProposerSNS: &sns, ProposerWallet: "0x1", Status: model.ProposalStatusError, ResolutionRule: model.ProposalResolutionStandard, CreatedAt: now.Add(-time.Hour), UpdatedAt: now},
	} {
		require.NoError(t, svc.db.Create(&p).Error)
	}

	list, err := svc.List(10, 0, "", "")
	require.NoError(t, err)
	require.Len(t, list.Items, 1)
	assert.Equal(t, "ok", list.Items[0].Title)
}

func TestGetProposalByLegacyID(t *testing.T) {
	svc := newProposalTestDB(t)
	sns := "legacy"
	legacy := int64(449)
	now := time.Now()
	require.NoError(t, svc.db.Create(&model.Proposal{
		ID: 123, LegacyID: &legacy, Type: model.ProposalTypeHallRegular, Title: "guide", Content: []byte(`{}`),
		ProposerSNS: &sns, ProposerWallet: "0x1", Status: model.ProposalStatusVoteFailed,
		ResolutionRule: model.ProposalResolutionStandard, CreatedAt: now, UpdatedAt: now,
	}).Error)

	p, err := svc.getProposal(449)
	require.NoError(t, err)
	assert.Equal(t, int64(123), p.ID)

	detail, err := svc.GetDetail(449, false, false, 50, 50)
	require.NoError(t, err)
	assert.Equal(t, int64(449), detail.ID)
	assert.Equal(t, "guide", detail.Title)

	list, err := svc.List(10, 0, "", "")
	require.NoError(t, err)
	require.Len(t, list.Items, 1)
	assert.Equal(t, int64(449), list.Items[0].ID)
}

func TestFormatProposalDisplayTitle(t *testing.T) {
	num := "449"
	assert.Equal(t, "SIP-449 专栏标题", formatProposalDisplayTitle("专栏标题", &num))
	label := "SIP-449"
	assert.Equal(t, "SIP-449 专栏标题", formatProposalDisplayTitle("专栏标题", &label))
	assert.Equal(t, "专栏标题", formatProposalDisplayTitle("专栏标题", nil))
	assert.Equal(t, "SIP-449 专栏标题", formatProposalDisplayTitle("SIP-449 专栏标题", &num))
}
