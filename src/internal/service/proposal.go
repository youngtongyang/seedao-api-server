package service

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/seedao/seedao-api-server/internal/config"
	"github.com/seedao/seedao-api-server/internal/model"
)

type ProposalService struct {
	db  *gorm.DB
	acl governanceACL
	cfg config.GovernanceConfig
}

func NewProposalService(db *gorm.DB, cfg config.GovernanceConfig) *ProposalService {
	return &ProposalService{db: db, acl: newGovernanceACL(cfg), cfg: cfg}
}

type ProposalListItem struct {
	ID             int64   `json:"id"`
	Type           string  `json:"type"`
	Title          string  `json:"title"`
	Status         string  `json:"status"`
	ProposerSNS    *string `json:"proposer_sns"`
	PublishedAt    *string `json:"published_at,omitempty"`
	PublicityEndAt *string `json:"publicity_end_at,omitempty"`
	VoteEndAt      *string `json:"vote_end_at,omitempty"`
	CreatedAt      string  `json:"created_at"`
	VoteCount      int     `json:"vote_count"`
	CommentCount   int     `json:"comment_count"`
}

type ProposalListResult struct {
	Items      []ProposalListItem `json:"items"`
	NextBefore *int64             `json:"next_before,omitempty"`
}

type ProposalDetail struct {
	ProposalResponse
	VoteSummary VoteTally             `json:"vote_summary"`
	Votes       []ProposalVoteItem    `json:"votes,omitempty"`
	Comments    []ProposalCommentItem `json:"comments,omitempty"`
}

type ProposalResponse struct {
	ID                  int64           `json:"id"`
	Type                string          `json:"type"`
	Title               string          `json:"title"`
	Content             json.RawMessage `json:"content"`
	ProposerSNS         *string         `json:"proposer_sns"`
	ProposerWallet      string          `json:"proposer_wallet"`
	Status              string          `json:"status"`
	VoteType            int16           `json:"vote_type"`
	VoteOptions         json.RawMessage `json:"vote_options,omitempty"`
	ResolutionRule      string          `json:"resolution_rule"`
	PublishedAt         *string         `json:"published_at,omitempty"`
	PublicityEndAt      *string         `json:"publicity_end_at,omitempty"`
	VoteEndAt           *string         `json:"vote_end_at,omitempty"`
	PublicitySeconds    *int            `json:"publicity_seconds,omitempty"`
	VoteDurationSeconds *int            `json:"vote_duration_seconds,omitempty"`
	CreatedAt           string          `json:"created_at"`
	UpdatedAt           string          `json:"updated_at"`
}

type ProposalVoteItem struct {
	ID          int64   `json:"id"`
	SNSName     *string `json:"sns_name"`
	VoterWallet string  `json:"voter_wallet"`
	VoteResult  string  `json:"vote_result"`
	VoteWeight  int     `json:"vote_weight"`
	VotedAt     string  `json:"voted_at"`
}

type ProposalCommentItem struct {
	ID           int64   `json:"id"`
	SNSName      *string `json:"sns_name"`
	AuthorWallet string  `json:"author_wallet"`
	ParentID     *int64  `json:"parent_id"`
	Content      string  `json:"content"`
	IsReject     bool    `json:"is_reject"`
	CommentedAt  string  `json:"commented_at"`
}

type GovernanceNodesResult struct {
	QuarterNodeTotal         int      `json:"quarter_node_total"`
	HallProposalSNS          []string `json:"hall_proposal_sns"`
	NodeConsensusProposalSNS []string `json:"node_consensus_proposal_sns"`
	HallVoterSNS             []string `json:"hall_voter_sns"`
	NodeVoterSNS             []string `json:"node_voter_sns"`
}

type CreateProposalInput struct {
	Type           string
	Title          string
	Content        json.RawMessage
	ResolutionRule string
	ActorSNS       string
	ActorWallet    string
}

type UpdateProposalInput struct {
	Title    *string
	Content  json.RawMessage
	ActorSNS string
}

type PublishProposalInput struct {
	PublicitySeconds    *int
	VoteDurationSeconds *int
	ActorSNS            string
}

type CastVoteInput struct {
	ProposalID  int64
	VoteResult  string
	ActorSNS    string
	ActorWallet string
}

type CreateCommentInput struct {
	ProposalID  int64
	Content     string
	ParentID    *int64
	IsReject    bool
	ActorSNS    string
	ActorWallet string
}

func (s *ProposalService) GovernanceNodes() *GovernanceNodesResult {
	return &GovernanceNodesResult{
		QuarterNodeTotal:         s.cfg.QuarterNodeTotal,
		HallProposalSNS:          append([]string(nil), s.cfg.HallProposalSNS...),
		NodeConsensusProposalSNS: append([]string(nil), s.cfg.NodeConsensusProposalSNS...),
		HallVoterSNS:             append([]string(nil), s.cfg.HallVoterSNS...),
		NodeVoterSNS:             append([]string(nil), s.cfg.NodeVoterSNS...),
	}
}

func (s *ProposalService) List(limit int, beforeID int64, proposalType, status string) (*ProposalListResult, error) {
	if limit <= 0 {
		limit = 20
	}
	if limit > 50 {
		limit = 50
	}

	q := s.db.Model(&model.Proposal{})
	if proposalType != "" {
		q = q.Where("type = ?", proposalType)
	}
	if status != "" {
		q = q.Where("status = ?", status)
	}
	if beforeID > 0 {
		q = q.Where("id < ?", beforeID)
	}

	var rows []model.Proposal
	if err := q.Order("id DESC").Limit(limit + 1).Find(&rows).Error; err != nil {
		return nil, err
	}

	var nextBefore *int64
	if len(rows) > limit {
		id := rows[limit].ID
		nextBefore = &id
		rows = rows[:limit]
	}

	items := make([]ProposalListItem, 0, len(rows))
	for _, p := range rows {
		item := ProposalListItem{
			ID:          p.ID,
			Type:        p.Type,
			Title:       p.Title,
			Status:      p.Status,
			ProposerSNS: p.ProposerSNS,
			CreatedAt:   formatTimeUTC(p.CreatedAt),
		}
		item.PublishedAt = formatOptionalTime(p.PublishedAt)
		item.PublicityEndAt = formatOptionalTime(p.PublicityEndAt)
		item.VoteEndAt = formatOptionalTime(p.VoteEndAt)

		var voteCount, commentCount int64
		_ = s.db.Model(&model.ProposalVote{}).Where("proposal_id = ?", p.ID).Count(&voteCount).Error
		_ = s.db.Model(&model.ProposalComment{}).Where("proposal_id = ?", p.ID).Count(&commentCount).Error
		item.VoteCount = int(voteCount)
		item.CommentCount = int(commentCount)
		items = append(items, item)
	}

	return &ProposalListResult{Items: items, NextBefore: nextBefore}, nil
}

func (s *ProposalService) GetDetail(id int64, includeVotes, includeComments bool, voteLimit, commentLimit int) (*ProposalDetail, error) {
	p, err := s.getProposal(id)
	if err != nil {
		return nil, err
	}

	tally, err := s.tallyVotes(p)
	if err != nil {
		return nil, err
	}

	detail := &ProposalDetail{
		ProposalResponse: toProposalResponse(p),
		VoteSummary:      tally,
	}
	if includeVotes {
		detail.Votes, err = s.listVotes(id, voteLimit)
		if err != nil {
			return nil, err
		}
	}
	if includeComments {
		detail.Comments, err = s.listComments(id, commentLimit)
		if err != nil {
			return nil, err
		}
	}
	return detail, nil
}

func (s *ProposalService) GetVoteSummary(proposalID int64) (*VoteTally, error) {
	p, err := s.getProposal(proposalID)
	if err != nil {
		return nil, err
	}
	tally, err := s.tallyVotes(p)
	if err != nil {
		return nil, err
	}
	return &tally, nil
}

func (s *ProposalService) ListVotes(proposalID int64, limit int) ([]ProposalVoteItem, error) {
	if _, err := s.getProposal(proposalID); err != nil {
		return nil, err
	}
	return s.listVotes(proposalID, limit)
}

func (s *ProposalService) GetMyVote(proposalID int64, actorSNS string) (*ProposalVoteItem, error) {
	if actorSNS == "" {
		return nil, ErrProposalSNSRequired
	}
	if _, err := s.getProposal(proposalID); err != nil {
		return nil, err
	}
	var row model.ProposalVote
	err := s.db.Where("proposal_id = ? AND LOWER(sns_name) = ?", proposalID, strings.ToLower(actorSNS)).First(&row).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrProposalVoteNotFound
	}
	if err != nil {
		return nil, err
	}
	item := toVoteItem(&row)
	return &item, nil
}

func (s *ProposalService) ListComments(proposalID int64, limit int) ([]ProposalCommentItem, error) {
	if _, err := s.getProposal(proposalID); err != nil {
		return nil, err
	}
	return s.listComments(proposalID, limit)
}

func (s *ProposalService) CreateDraft(input CreateProposalInput) (*ProposalResponse, error) {
	if !isHallOrNodeCreateType(input.Type) {
		return nil, ErrInvalidParam
	}
	if strings.TrimSpace(input.Title) == "" || len(input.Content) == 0 {
		return nil, ErrInvalidParam
	}
	if input.ActorSNS == "" || !s.acl.canCreateType(input.ActorSNS, input.Type) {
		return nil, ErrProposalForbidden
	}

	rule := input.ResolutionRule
	if rule == "" {
		rule = model.ProposalResolutionStandard
	}
	if rule != model.ProposalResolutionStandard && rule != model.ProposalResolutionMetaRule {
		return nil, ErrInvalidParam
	}

	now := time.Now()
	sns := input.ActorSNS
	wallet := normalizeWallet(input.ActorWallet)

	p := model.Proposal{
		Type:           input.Type,
		Title:          strings.TrimSpace(input.Title),
		Content:        []byte(input.Content),
		ProposerSNS:    &sns,
		ProposerWallet: wallet,
		Status:         model.ProposalStatusDraft,
		VoteType:       model.ProposalVoteTypeStandard,
		VoteOptions:    defaultVoteOptionsJSON(),
		ResolutionRule: rule,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	if err := s.db.Create(&p).Error; err != nil {
		return nil, err
	}
	resp := toProposalResponse(&p)
	return &resp, nil
}

func (s *ProposalService) UpdateDraft(id int64, input UpdateProposalInput) (*ProposalResponse, error) {
	p, err := s.getProposal(id)
	if err != nil {
		return nil, err
	}
	if !isEditableStatus(p.Status) {
		return nil, ErrProposalInvalidStatus
	}
	if !s.isProposer(p, input.ActorSNS) {
		return nil, ErrProposalForbidden
	}

	updates := map[string]interface{}{"updated_at": time.Now()}
	if input.Title != nil && strings.TrimSpace(*input.Title) != "" {
		updates["title"] = strings.TrimSpace(*input.Title)
	}
	if len(input.Content) > 0 {
		updates["content"] = []byte(input.Content)
	}
	if len(updates) == 1 {
		return nil, ErrInvalidParam
	}
	if err := s.db.Model(p).Updates(updates).Error; err != nil {
		return nil, err
	}
	updated, err := s.getProposal(id)
	if err != nil {
		return nil, err
	}
	resp := toProposalResponse(updated)
	return &resp, nil
}

func (s *ProposalService) CloseDraft(id int64, actorSNS string) (*ProposalResponse, error) {
	p, err := s.getProposal(id)
	if err != nil {
		return nil, err
	}
	if !isEditableStatus(p.Status) {
		return nil, ErrProposalInvalidStatus
	}
	if !s.isProposer(p, actorSNS) {
		return nil, ErrProposalForbidden
	}
	if err := s.db.Model(p).Updates(map[string]interface{}{
		"status":     model.ProposalStatusClosed,
		"updated_at": time.Now(),
	}).Error; err != nil {
		return nil, err
	}
	p.Status = model.ProposalStatusClosed
	resp := toProposalResponse(p)
	return &resp, nil
}

func (s *ProposalService) Publish(id int64, input PublishProposalInput) (*ProposalResponse, error) {
	p, err := s.getProposal(id)
	if err != nil {
		return nil, err
	}
	if !isEditableStatus(p.Status) {
		return nil, ErrProposalInvalidStatus
	}
	if !s.isProposer(p, input.ActorSNS) {
		return nil, ErrProposalForbidden
	}

	pubSec := s.cfg.DefaultPublicitySeconds
	voteSec := s.cfg.DefaultVoteDurationSeconds
	if input.PublicitySeconds != nil && *input.PublicitySeconds > 0 {
		pubSec = *input.PublicitySeconds
	}
	if input.VoteDurationSeconds != nil && *input.VoteDurationSeconds > 0 {
		voteSec = *input.VoteDurationSeconds
	}

	now := time.Now()
	pubEnd := now.Add(time.Duration(pubSec) * time.Second)
	voteEnd := pubEnd.Add(time.Duration(voteSec) * time.Second)

	updates := map[string]interface{}{
		"status":                model.ProposalStatusPublicity,
		"published_at":          now,
		"publicity_end_at":      pubEnd,
		"vote_end_at":           voteEnd,
		"publicity_seconds":     pubSec,
		"vote_duration_seconds": voteSec,
		"updated_at":            now,
	}
	if err := s.db.Model(p).Updates(updates).Error; err != nil {
		return nil, err
	}
	updated, err := s.getProposal(id)
	if err != nil {
		return nil, err
	}
	resp := toProposalResponse(updated)
	return &resp, nil
}

func (s *ProposalService) CastVote(input CastVoteInput) (*ProposalVoteItem, error) {
	if strings.TrimSpace(input.VoteResult) == "" {
		return nil, ErrInvalidParam
	}
	if input.ActorSNS == "" {
		return nil, ErrProposalSNSRequired
	}

	p, err := s.getProposal(input.ProposalID)
	if err != nil {
		return nil, err
	}
	if !s.acl.canVote(p.Type, input.ActorSNS) {
		return nil, ErrProposalNotVoter
	}
	if !s.isVotingOpen(p) {
		return nil, ErrProposalVoteClosed
	}

	var existing int64
	if err := s.db.Model(&model.ProposalVote{}).
		Where("proposal_id = ? AND LOWER(sns_name) = ?", p.ID, strings.ToLower(input.ActorSNS)).
		Count(&existing).Error; err != nil {
		return nil, err
	}
	if existing > 0 {
		return nil, ErrProposalAlreadyVoted
	}

	sns := input.ActorSNS
	now := time.Now()
	v := model.ProposalVote{
		ProposalID:  p.ID,
		SNSName:     &sns,
		VoterWallet: normalizeWallet(input.ActorWallet),
		VoteResult:  strings.TrimSpace(input.VoteResult),
		VoteWeight:  1,
		VotedAt:     now,
	}
	if err := s.db.Create(&v).Error; err != nil {
		if isUniqueViolation(err) {
			return nil, ErrProposalAlreadyVoted
		}
		return nil, err
	}
	item := toVoteItem(&v)
	return &item, nil
}

func (s *ProposalService) CreateComment(input CreateCommentInput) (*ProposalCommentItem, error) {
	if strings.TrimSpace(input.Content) == "" {
		return nil, ErrInvalidParam
	}
	if input.ActorSNS == "" {
		return nil, ErrProposalSNSRequired
	}

	p, err := s.getProposal(input.ProposalID)
	if err != nil {
		return nil, err
	}
	if !isCommentableStatus(p.Status) {
		return nil, ErrProposalInvalidStatus
	}
	if input.ParentID != nil {
		var parent model.ProposalComment
		if err := s.db.Where("id = ? AND proposal_id = ?", *input.ParentID, p.ID).First(&parent).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrInvalidParam
			}
			return nil, err
		}
	}

	sns := input.ActorSNS
	now := time.Now()
	c := model.ProposalComment{
		ProposalID:   p.ID,
		SNSName:      &sns,
		AuthorWallet: normalizeWallet(input.ActorWallet),
		ParentID:     input.ParentID,
		Content:      strings.TrimSpace(input.Content),
		IsReject:     input.IsReject,
		CommentedAt:  now,
	}
	if err := s.db.Create(&c).Error; err != nil {
		return nil, err
	}
	item := toCommentItem(&c)
	return &item, nil
}

func (s *ProposalService) FinalizeVote(id int64, actorSNS string) (*ProposalDetail, error) {
	p, err := s.getProposal(id)
	if err != nil {
		return nil, err
	}
	if !s.canFinalize(p, actorSNS) {
		return nil, ErrProposalForbidden
	}
	if p.VoteEndAt == nil || time.Now().Before(*p.VoteEndAt) {
		return nil, ErrProposalVoteClosed
	}
	if p.Status != model.ProposalStatusPublicity && p.Status != model.ProposalStatusPublished {
		return nil, ErrProposalInvalidStatus
	}

	tally, err := s.tallyVotes(p)
	if err != nil {
		return nil, err
	}
	newStatus := model.ProposalStatusVoteFailed
	if tally.Passed {
		newStatus = model.ProposalStatusVotePassed
	}
	if err := s.db.Model(p).Updates(map[string]interface{}{
		"status":     newStatus,
		"updated_at": time.Now(),
	}).Error; err != nil {
		return nil, err
	}
	p.Status = newStatus
	return &ProposalDetail{
		ProposalResponse: toProposalResponse(p),
		VoteSummary:      tally,
	}, nil
}

func (s *ProposalService) canFinalize(p *model.Proposal, actorSNS string) bool {
	if s.isProposer(p, actorSNS) {
		return true
	}
	return s.acl.canCreateType(actorSNS, p.Type)
}

func (s *ProposalService) getProposal(id int64) (*model.Proposal, error) {
	var p model.Proposal
	if err := s.db.First(&p, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProposalNotFound
		}
		return nil, err
	}
	return &p, nil
}

func (s *ProposalService) isProposer(p *model.Proposal, sns string) bool {
	if p.ProposerSNS == nil {
		return false
	}
	return s.acl.normalize(*p.ProposerSNS) == s.acl.normalize(sns)
}

func (s *ProposalService) isVotingOpen(p *model.Proposal) bool {
	if p.PublicityEndAt == nil || p.VoteEndAt == nil {
		return false
	}
	now := time.Now()
	if now.Before(*p.PublicityEndAt) || now.After(*p.VoteEndAt) {
		return false
	}
	switch p.Status {
	case model.ProposalStatusPublicity, model.ProposalStatusPublished:
		return true
	default:
		return false
	}
}

func (s *ProposalService) tallyVotes(p *model.Proposal) (VoteTally, error) {
	var votes []model.ProposalVote
	if err := s.db.Where("proposal_id = ?", p.ID).Find(&votes).Error; err != nil {
		return VoteTally{}, err
	}

	tally := VoteTally{ByResult: map[string]int{}}
	for _, v := range votes {
		tally.TotalVotes++
		tally.ByResult[v.VoteResult]++
		switch classifyVoteResult(v.VoteResult) {
		case "approve":
			tally.ApproveCount++
		case "reject":
			tally.RejectCount++
		default:
			tally.AbstainCount++
		}
	}
	return evaluateResolution(p.ResolutionRule, s.cfg.QuarterNodeTotal, tally), nil
}

func (s *ProposalService) listVotes(proposalID int64, limit int) ([]ProposalVoteItem, error) {
	if limit <= 0 {
		limit = 50
	}
	if limit > 200 {
		limit = 200
	}
	var rows []model.ProposalVote
	if err := s.db.Where("proposal_id = ?", proposalID).Order("id DESC").Limit(limit).Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]ProposalVoteItem, 0, len(rows))
	for i := range rows {
		out = append(out, toVoteItem(&rows[i]))
	}
	return out, nil
}

func (s *ProposalService) listComments(proposalID int64, limit int) ([]ProposalCommentItem, error) {
	if limit <= 0 {
		limit = 50
	}
	if limit > 200 {
		limit = 200
	}
	var rows []model.ProposalComment
	if err := s.db.Where("proposal_id = ?", proposalID).Order("id ASC").Limit(limit).Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]ProposalCommentItem, 0, len(rows))
	for i := range rows {
		out = append(out, toCommentItem(&rows[i]))
	}
	return out, nil
}

func toProposalResponse(p *model.Proposal) ProposalResponse {
	resp := ProposalResponse{
		ID:             p.ID,
		Type:           p.Type,
		Title:          p.Title,
		Content:        json.RawMessage(p.Content),
		ProposerSNS:    p.ProposerSNS,
		ProposerWallet: p.ProposerWallet,
		Status:         p.Status,
		VoteType:       p.VoteType,
		ResolutionRule: p.ResolutionRule,
		CreatedAt:      formatTimeUTC(p.CreatedAt),
		UpdatedAt:      formatTimeUTC(p.UpdatedAt),
	}
	if len(p.VoteOptions) > 0 {
		resp.VoteOptions = json.RawMessage(p.VoteOptions)
	}
	resp.PublishedAt = formatOptionalTime(p.PublishedAt)
	resp.PublicityEndAt = formatOptionalTime(p.PublicityEndAt)
	resp.VoteEndAt = formatOptionalTime(p.VoteEndAt)
	resp.PublicitySeconds = p.PublicitySeconds
	resp.VoteDurationSeconds = p.VoteDurationSeconds
	return resp
}

func toVoteItem(v *model.ProposalVote) ProposalVoteItem {
	return ProposalVoteItem{
		ID:          v.ID,
		SNSName:     v.SNSName,
		VoterWallet: v.VoterWallet,
		VoteResult:  v.VoteResult,
		VoteWeight:  v.VoteWeight,
		VotedAt:     formatTimeUTC(v.VotedAt),
	}
}

func toCommentItem(c *model.ProposalComment) ProposalCommentItem {
	return ProposalCommentItem{
		ID:           c.ID,
		SNSName:      c.SNSName,
		AuthorWallet: c.AuthorWallet,
		ParentID:     c.ParentID,
		Content:      c.Content,
		IsReject:     c.IsReject,
		CommentedAt:  formatTimeUTC(c.CommentedAt),
	}
}

func normalizeWallet(wallet string) string {
	w := strings.ToLower(strings.TrimSpace(wallet))
	if w == "" {
		return "0x0000000000000000000000000000000000000000"
	}
	return w
}

func isCommentableStatus(status string) bool {
	switch status {
	case model.ProposalStatusDraft, model.ProposalStatusClosed:
		return false
	default:
		return true
	}
}

func isUniqueViolation(err error) bool {
	if err == nil {
		return false
	}
	s := strings.ToLower(err.Error())
	return strings.Contains(s, "duplicate") || strings.Contains(s, "unique constraint")
}
