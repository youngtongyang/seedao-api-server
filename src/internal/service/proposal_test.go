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
	return NewProposalService(db, NewSNSService(db, nil), testGovernanceCfg())
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
	assert.Equal(t, int64(123), detail.ID)
	assert.Equal(t, "guide", detail.Title)

	list, err := svc.List(10, 0, "", "")
	require.NoError(t, err)
	require.Len(t, list.Items, 1)
	assert.Equal(t, int64(123), list.Items[0].ID)
}

func TestGetProposalIDPriorityOverLegacyID(t *testing.T) {
	svc := newProposalTestDB(t)
	sns := "legacy"
	now := time.Now()
	legacy247 := int64(247)

	require.NoError(t, svc.db.Create(&model.Proposal{
		ID: 225, LegacyID: &legacy247, Type: model.ProposalTypeHallRegular, Title: "other legacy 247", Content: []byte(`{}`),
		ProposerSNS: &sns, ProposerWallet: "0x1", Status: model.ProposalStatusVotePassed,
		ResolutionRule: model.ProposalResolutionStandard, CreatedAt: now, UpdatedAt: now,
	}).Error)
	require.NoError(t, svc.db.Create(&model.Proposal{
		ID: 247, Type: model.ProposalTypeHallRegular, Title: "table id 247", Content: []byte(`{}`),
		ProposerSNS: &sns, ProposerWallet: "0x2", Status: model.ProposalStatusVotePassed,
		ResolutionRule: model.ProposalResolutionStandard, CreatedAt: now, UpdatedAt: now,
	}).Error)

	p, err := svc.getProposal(247)
	require.NoError(t, err)
	assert.Equal(t, int64(247), p.ID)
	assert.Equal(t, "table id 247", p.Title)

	detail, err := svc.GetDetail(247, false, false, 0, 0)
	require.NoError(t, err)
	assert.Equal(t, "table id 247", detail.Title)
}

func TestFormatProposalDisplayTitle(t *testing.T) {
	assert.Equal(t, "SIP-449: 专栏标题", formatProposalDisplayTitle("专栏标题", 449))
	assert.Equal(t, "专栏标题", formatProposalDisplayTitle("专栏标题", 0))
	assert.Equal(t, "SIP-449 专栏标题", formatProposalDisplayTitle("SIP-449 专栏标题", 449))
}

func TestProposalProposerSNSFromWallet(t *testing.T) {
	svc := newProposalTestDB(t)
	now := time.Now()
	wallet := "0xAbCdEf1234567890123456789012345678901234"
	walletLower := NormalizeWallet(wallet)

	t.Run("list resolves from sns_names", func(t *testing.T) {
		require.NoError(t, svc.db.Create(&model.SNSName{
			Name: "alice.seedao", UserID: 1, MintStatus: "minted",
			Wallet: &walletLower, CreatedAt: now,
		}).Error)
		require.NoError(t, svc.db.Create(&model.Proposal{
			Type: model.ProposalTypeHallRegular, Title: "from sns_names", Content: []byte(`{}`),
			ProposerWallet: wallet, Status: model.ProposalStatusVotePassed,
			ResolutionRule: model.ProposalResolutionStandard, CreatedAt: now, UpdatedAt: now,
		}).Error)

		list, err := svc.List(10, 0, "", "")
		require.NoError(t, err)
		require.Len(t, list.Items, 1)
		require.NotNil(t, list.Items[0].ProposerSNS)
		assert.Equal(t, "alice.seedao", *list.Items[0].ProposerSNS)
	})

	t.Run("detail falls back to sns_chain_registry", func(t *testing.T) {
		regWallet := "0x2222222222222222222222222222222222222222"
		require.NoError(t, svc.db.Create(&model.SNSChainRegistry{
			Wallet: regWallet, SNSName: "bob.seedao", SyncedAt: now, CreatedAt: now, UpdatedAt: now,
		}).Error)
		var p model.Proposal
		require.NoError(t, svc.db.Create(&model.Proposal{
			Type: model.ProposalTypeHallRegular, Title: "from registry", Content: []byte(`{}`),
			ProposerWallet: regWallet, Status: model.ProposalStatusVotePassed,
			ResolutionRule: model.ProposalResolutionStandard, CreatedAt: now, UpdatedAt: now,
		}).Error)
		require.NoError(t, svc.db.First(&p, "title = ?", "from registry").Error)

		detail, err := svc.GetDetail(p.ID, false, false, 0, 0)
		require.NoError(t, err)
		require.NotNil(t, detail.ProposerSNS)
		assert.Equal(t, "bob.seedao", *detail.ProposerSNS)
	})

	t.Run("keeps existing proposer_sns", func(t *testing.T) {
		existing := "existing.seedao"
		require.NoError(t, svc.db.Create(&model.Proposal{
			Type: model.ProposalTypeHallRegular, Title: "has sns", Content: []byte(`{}`),
			ProposerSNS: &existing, ProposerWallet: wallet,
			Status: model.ProposalStatusVotePassed, ResolutionRule: model.ProposalResolutionStandard,
			CreatedAt: now, UpdatedAt: now,
		}).Error)

		list, err := svc.List(10, 0, "", "")
		require.NoError(t, err)
		var item *ProposalListItem
		for i := range list.Items {
			if list.Items[i].Title == "has sns" {
				item = &list.Items[i]
				break
			}
		}
		require.NotNil(t, item)
		require.NotNil(t, item.ProposerSNS)
		assert.Equal(t, "existing.seedao", *item.ProposerSNS)
	})
}

func TestProposalVoteSNSFromWallet(t *testing.T) {
	svc := newProposalTestDB(t)
	now := time.Now()
	voterWallet := "0x1b76b011c0a1e03896bfc8923be228e430bb9049"

	require.NoError(t, svc.db.Create(&model.SNSChainRegistry{
		Wallet: voterWallet, SNSName: "nechi_luo.seedao", SyncedAt: now, CreatedAt: now, UpdatedAt: now,
	}).Error)

	var p model.Proposal
	require.NoError(t, svc.db.Create(&model.Proposal{
		Type: model.ProposalTypeHallRegular, Title: "vote sns", Content: []byte(`{}`),
		ProposerWallet: "0x1", Status: model.ProposalStatusVotePassed,
		ResolutionRule: model.ProposalResolutionStandard, CreatedAt: now, UpdatedAt: now,
	}).Error)
	require.NoError(t, svc.db.First(&p, "title = ?", "vote sns").Error)

	require.NoError(t, svc.db.Create(&model.ProposalVote{
		ProposalID: p.ID, VoterWallet: voterWallet, VoteResult: "同意", VoteWeight: 1, VotedAt: now,
	}).Error)
	existing := "baiyu"
	require.NoError(t, svc.db.Create(&model.ProposalVote{
		ProposalID: p.ID, SNSName: &existing, VoterWallet: "0x2", VoteResult: "同意", VoteWeight: 1, VotedAt: now,
	}).Error)

	votes, err := svc.ListVotes(p.ID, 10)
	require.NoError(t, err)
	require.Len(t, votes, 2)

	var resolved *ProposalVoteItem
	for i := range votes {
		if votes[i].VoterWallet == voterWallet {
			resolved = &votes[i]
			break
		}
	}
	require.NotNil(t, resolved)
	require.NotNil(t, resolved.SNSName)
	assert.Equal(t, "nechi_luo.seedao", *resolved.SNSName)
}

func TestParseCommentContent(t *testing.T) {
	assert.Equal(t, "hello", parseCommentContent(`[{"insert":"hello\n"}]`))
	assert.Equal(t, "plain text", parseCommentContent("plain text"))
	assert.Equal(t, "我不能投票嗎", parseCommentContent("我不能投票嗎"))
}

func TestProposalCommentTree(t *testing.T) {
	svc := newProposalTestDB(t)
	now := time.Now()
	sns := "shawn"
	parent516 := int64(516)
	parent533 := int64(533)

	require.NoError(t, svc.db.Create(&model.Proposal{
		Type: model.ProposalTypeHallRegular, Title: "comments", Content: []byte(`{}`),
		ProposerWallet: "0x1", Status: model.ProposalStatusVotePassed,
		ResolutionRule: model.ProposalResolutionStandard, CreatedAt: now, UpdatedAt: now,
	}).Error)
	var p model.Proposal
	require.NoError(t, svc.db.First(&p, "title = ?", "comments").Error)

	seed := []model.ProposalComment{
		{ID: 516, ProposalID: p.ID, SNSName: &sns, AuthorWallet: "0x1", Content: `[{"insert":"question\n"}]`, CommentedAt: now.Add(-3 * time.Hour)},
		{ID: 485, ProposalID: p.ID, AuthorWallet: "0x2", ParentID: &parent516, Content: `[{"insert":"answer\n"}]`, CommentedAt: now.Add(-2 * time.Hour)},
		{ID: 533, ProposalID: p.ID, AuthorWallet: "0x3", Content: "top level", CommentedAt: now.Add(-time.Hour)},
		{ID: 534, ProposalID: p.ID, AuthorWallet: "0x2", ParentID: &parent533, Content: `[{"insert":"reply\n"}]`, CommentedAt: now},
	}
	for i := range seed {
		require.NoError(t, svc.db.Create(&seed[i]).Error)
	}

	comments, err := svc.ListComments(p.ID, 50)
	require.NoError(t, err)
	require.Len(t, comments, 2)
	assert.Equal(t, int64(516), comments[0].ID)
	assert.Equal(t, "question", comments[0].Content)
	require.Len(t, comments[0].Replies, 1)
	assert.Equal(t, int64(485), comments[0].Replies[0].ID)
	assert.Equal(t, "answer", comments[0].Replies[0].Content)

	assert.Equal(t, int64(533), comments[1].ID)
	assert.Equal(t, "top level", comments[1].Content)
	require.Len(t, comments[1].Replies, 1)
	assert.Equal(t, "reply", comments[1].Replies[0].Content)
}
