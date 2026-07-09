package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/seedao/seedao-api-server/internal/config"
	"github.com/seedao/seedao-api-server/internal/model"
)

type ProposalService struct {
	db  *gorm.DB
	sns *SNSService
	acl governanceACL
	cfg config.GovernanceConfig
}

func NewProposalService(db *gorm.DB, sns *SNSService, cfg config.GovernanceConfig) *ProposalService {
	return &ProposalService{db: db, sns: sns, acl: newGovernanceACL(cfg), cfg: cfg}
}

type ProposalListItem struct {
	ID             int64   `json:"id"`
	Type           string  `json:"type"`
	Title          string  `json:"title"`
	SIP            int     `json:"sip,omitempty"`
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
	SIP                 int             `json:"sip,omitempty"`
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
	ID           int64                 `json:"id"`
	SNSName      *string               `json:"sns_name"`
	AuthorWallet string                `json:"author_wallet"`
	ParentID     *int64                `json:"parent_id,omitempty"`
	Content      string                `json:"content"`
	IsReject     bool                  `json:"is_reject"`
	CommentedAt  string                `json:"commented_at"`
	Replies      []ProposalCommentItem `json:"replies,omitempty"`
}

type GovernanceNodesResult struct {
	QuarterNodeTotal         int      `json:"quarter_node_total"`
	HallProposalSNS          []string `json:"hall_proposal_sns"`
	NodeConsensusProposalSNS []string `json:"node_consensus_proposal_sns"`
	HallVoterSNS             []string `json:"hall_voter_sns"`
	NodeVoterSNS             []string `json:"node_voter_sns"`
}

type ProposalCreatePermissionsResult struct {
	SNSName      string          `json:"sns_name"`
	CanCreateAny bool            `json:"can_create_any"`
	CreateTypes  map[string]bool `json:"create_types"`
}

type CreateProposalInput struct {
	Type           string
	Title          string
	Content        json.RawMessage
	ResolutionRule string
	VoteType       int16
	VoteOptions    json.RawMessage
	ActorSNS       string
	ActorWallet    string
}

type UpdateProposalInput struct {
	Title       *string
	Content     json.RawMessage
	VoteType    *int16
	VoteOptions json.RawMessage
	ActorSNS    string
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

func (s *ProposalService) CreatePermissions(snsName string) *ProposalCreatePermissionsResult {
	createTypes := make(map[string]bool, len(creatableProposalTypeAPIKeys))
	canCreateAny := false
	for _, apiType := range creatableProposalTypeAPIKeys {
		normalized, err := normalizeCreateProposalType(apiType)
		allowed := err == nil && snsName != "" && s.acl.canCreateType(snsName, normalized)
		createTypes[apiType] = allowed
		if allowed {
			canCreateAny = true
		}
	}
	return &ProposalCreatePermissionsResult{
		SNSName:      snsName,
		CanCreateAny: canCreateAny,
		CreateTypes:  createTypes,
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
	} else {
		q = q.Where("status <> ?", model.ProposalStatusError)
	}
	if beforeID > 0 {
		anchor, err := s.getProposal(beforeID)
		if err != nil {
			return &ProposalListResult{Items: []ProposalListItem{}}, nil
		}
		q = q.Where("(created_at < ?) OR (created_at = ? AND id < ?)", anchor.CreatedAt, anchor.CreatedAt, anchor.ID)
	}

	var rows []model.Proposal
	if err := q.Order("created_at DESC, id DESC").Limit(limit + 1).Find(&rows).Error; err != nil {
		return nil, err
	}

	var nextBefore *int64
	if len(rows) > limit {
		nextID := rows[limit].ID
		nextBefore = &nextID
		rows = rows[:limit]
	}

	items := make([]ProposalListItem, 0, len(rows))
	snsByWallet, err := s.resolveProposerSNSMap(rows)
	if err != nil {
		return nil, err
	}
	for i := range rows {
		p := rows[i]
		if updated, err := s.maybeAssignSIPOnVotingOpen(&p); err == nil && updated != nil {
			p = *updated
			rows[i] = p
		}
		item := ProposalListItem{
			ID:          p.ID,
			Type:        p.Type,
			Title:       formatProposalDisplayTitle(p.Title, p.SIP),
			SIP:         p.SIP,
			Status:      p.Status,
			ProposerSNS: proposerSNSForDisplay(&p, snsByWallet),
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
	p, err = s.maybeAssignSIPOnVotingOpen(p)
	if err != nil {
		return nil, err
	}

	tally, err := s.tallyVotes(p)
	if err != nil {
		return nil, err
	}

	detail := &ProposalDetail{
		ProposalResponse: s.toProposalResponse(p, nil),
		VoteSummary:      tally,
	}
	if includeVotes {
		detail.Votes, err = s.listVotes(p.ID, voteLimit)
		if err != nil {
			return nil, err
		}
	}
	if includeComments {
		detail.Comments, err = s.listComments(p.ID, commentLimit)
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
	p, err := s.getProposal(proposalID)
	if err != nil {
		return nil, err
	}
	return s.listVotes(p.ID, limit)
}

func (s *ProposalService) GetMyVote(proposalID int64, actorSNS string) (*ProposalVoteItem, error) {
	if actorSNS == "" {
		return nil, ErrProposalSNSRequired
	}
	p, err := s.getProposal(proposalID)
	if err != nil {
		return nil, err
	}
	var row model.ProposalVote
	err = s.db.Where("proposal_id = ? AND LOWER(sns_name) = ?", p.ID, strings.ToLower(actorSNS)).First(&row).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrProposalVoteNotFound
	}
	if err != nil {
		return nil, err
	}
	item := toVoteItem(&row, nil)
	if SNSNameNeedsResolve(row.SNSName) && !IsZeroWallet(row.VoterWallet) {
		snsByWallet, err := s.sns.LookupSNSByWallets([]string{row.VoterWallet})
		if err != nil {
			return nil, err
		}
		item = toVoteItem(&row, snsByWallet)
	}
	return &item, nil
}

func (s *ProposalService) ListComments(proposalID int64, limit int) ([]ProposalCommentItem, error) {
	p, err := s.getProposal(proposalID)
	if err != nil {
		return nil, err
	}
	return s.listComments(p.ID, limit)
}

func (s *ProposalService) CreateDraft(input CreateProposalInput) (*ProposalResponse, error) {
	proposalType, err := normalizeCreateProposalType(input.Type)
	if err != nil {
		return nil, err
	}
	content, err := NormalizeProposalContent(input.Content)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(input.Title) == "" {
		return nil, &ProposalParamError{Reason: "proposal title is required"}
	}

	ruleRaw := strings.TrimSpace(input.ResolutionRule)
	if ruleRaw == "" {
		ruleRaw = defaultResolutionRule(proposalType)
	}
	rule, err := normalizeResolutionRule(ruleRaw)
	if err != nil {
		return nil, err
	}

	if input.ActorSNS == "" {
		return nil, ErrProposalSNSRequired
	}
	if !s.acl.canCreateType(input.ActorSNS, proposalType) {
		return nil, &ProposalForbiddenError{
			Reason: fmt.Sprintf("sns %q is not allowed to create proposal type %q; check governance hall_proposal_sns / node_consensus_proposal_sns", input.ActorSNS, proposalType),
		}
	}

	voteOptions, err := NormalizeVoteOptions(input.VoteOptions)
	if err != nil {
		return nil, err
	}
	voteType := resolveVoteType(input.VoteType, content)

	now := time.Now()
	sns := input.ActorSNS
	wallet := NormalizeWallet(input.ActorWallet)

	p := model.Proposal{
		Type:           proposalType,
		Title:          strings.TrimSpace(input.Title),
		Content:        []byte(content),
		ProposerSNS:    &sns,
		ProposerWallet: wallet,
		Status:         model.ProposalStatusDraft,
		VoteType:       voteType,
		VoteOptions:    voteOptions,
		ResolutionRule: rule,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	if err := s.db.Create(&p).Error; err != nil {
		return nil, err
	}
	resp := s.toProposalResponse(&p, nil)
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
	var normalizedContent json.RawMessage
	if len(input.Content) > 0 {
		content, err := NormalizeProposalContent(input.Content)
		if err != nil {
			return nil, err
		}
		normalizedContent = content
		updates["content"] = []byte(content)
	}
	if input.VoteType != nil {
		contentForVoteType := normalizedContent
		if len(contentForVoteType) == 0 {
			contentForVoteType = json.RawMessage(p.Content)
		}
		updates["vote_type"] = resolveVoteType(*input.VoteType, contentForVoteType)
	} else if len(normalizedContent) > 0 {
		updates["vote_type"] = resolveVoteType(0, normalizedContent)
	}
	if len(input.VoteOptions) > 0 {
		voteOptions, err := NormalizeVoteOptions(input.VoteOptions)
		if err != nil {
			return nil, err
		}
		updates["vote_options"] = voteOptions
	}
	if len(updates) == 1 {
		return nil, &ProposalParamError{Reason: "no updatable fields provided; send title, content, vote_type, or vote_options"}
	}
	if err := s.db.Model(p).Updates(updates).Error; err != nil {
		return nil, err
	}
	updated, err := s.getProposal(id)
	if err != nil {
		return nil, err
	}
	resp := s.toProposalResponse(updated, nil)
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
	resp := s.toProposalResponse(p, nil)
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
	resp := s.toProposalResponse(updated, nil)
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
	p, err = s.maybeAssignSIPOnVotingOpen(p)
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
		VoterWallet: NormalizeWallet(input.ActorWallet),
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
	item := toVoteItem(&v, nil)
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
		AuthorWallet: NormalizeWallet(input.ActorWallet),
		ParentID:     input.ParentID,
		Content:      strings.TrimSpace(input.Content),
		IsReject:     input.IsReject,
		CommentedAt:  now,
	}
	if err := s.db.Create(&c).Error; err != nil {
		return nil, err
	}
	item := toCommentItem(&c, nil)
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
		ProposalResponse: s.toProposalResponse(p, nil),
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
	if err := s.db.First(&p, id).Error; err == nil {
		return &p, nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err := s.db.Where("legacy_id = ?", id).First(&p).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProposalNotFound
		}
		return nil, err
	}
	return &p, nil
}

func formatProposalDisplayTitle(title string, sip int) string {
	label := normalizeSIPLabel(sip)
	if label == "" {
		return title
	}
	if strings.HasPrefix(strings.TrimSpace(title), label) {
		return title
	}
	return label + ": " + title
}

func normalizeSIPLabel(sip int) string {
	if sip <= 0 {
		return ""
	}
	return fmt.Sprintf("SIP-%d", sip)
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
	return evaluateResolution(p.ResolutionRule, s.electorateTotal(p), tally), nil
}

func (s *ProposalService) electorateTotal(p *model.Proposal) int {
	switch p.Type {
	case model.ProposalTypeHallRegular:
		if n := len(s.cfg.HallVoterSNS); n > 0 {
			return n
		}
	}
	return s.cfg.QuarterNodeTotal
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
	snsByWallet, err := s.resolveVoteSNSMap(rows)
	if err != nil {
		return nil, err
	}
	out := make([]ProposalVoteItem, 0, len(rows))
	for i := range rows {
		out = append(out, toVoteItem(&rows[i], snsByWallet))
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
	q := s.db.Where("proposal_id = ?", proposalID).Order("commented_at ASC, id ASC")
	if limit > 0 {
		q = q.Limit(limit)
	}
	if err := q.Find(&rows).Error; err != nil {
		return nil, err
	}
	snsByWallet, err := s.resolveCommentSNSMap(rows)
	if err != nil {
		return nil, err
	}
	return buildCommentTree(rows, snsByWallet), nil
}

func (s *ProposalService) toProposalResponse(p *model.Proposal, snsByWallet map[string]string) ProposalResponse {
	if snsByWallet == nil && SNSNameNeedsResolve(p.ProposerSNS) && !IsZeroWallet(p.ProposerWallet) {
		var err error
		snsByWallet, err = s.sns.LookupSNSByWallets([]string{p.ProposerWallet})
		if err != nil {
			snsByWallet = map[string]string{}
		}
	}
	resp := ProposalResponse{
		ID:             p.ID,
		Type:           p.Type,
		Title:          formatProposalDisplayTitle(p.Title, p.SIP),
		SIP:            p.SIP,
		Content:        json.RawMessage(p.Content),
		ProposerSNS:    proposerSNSForDisplay(p, snsByWallet),
		ProposerWallet: p.ProposerWallet,
		Status:         p.Status,
		VoteType:       resolveVoteType(p.VoteType, json.RawMessage(p.Content)),
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

func (s *ProposalService) resolveProposerSNSMap(proposals []model.Proposal) (map[string]string, error) {
	wallets := make([]string, 0)
	for i := range proposals {
		p := &proposals[i]
		if !SNSNameNeedsResolve(p.ProposerSNS) || IsZeroWallet(p.ProposerWallet) {
			continue
		}
		wallets = append(wallets, p.ProposerWallet)
	}
	return s.sns.LookupSNSByWallets(wallets)
}

func (s *ProposalService) resolveVoteSNSMap(votes []model.ProposalVote) (map[string]string, error) {
	wallets := make([]string, 0)
	for i := range votes {
		v := &votes[i]
		if !SNSNameNeedsResolve(v.SNSName) || IsZeroWallet(v.VoterWallet) {
			continue
		}
		wallets = append(wallets, v.VoterWallet)
	}
	return s.sns.LookupSNSByWallets(wallets)
}

func (s *ProposalService) resolveCommentSNSMap(comments []model.ProposalComment) (map[string]string, error) {
	wallets := make([]string, 0)
	for i := range comments {
		c := &comments[i]
		if !SNSNameNeedsResolve(c.SNSName) || IsZeroWallet(c.AuthorWallet) {
			continue
		}
		wallets = append(wallets, c.AuthorWallet)
	}
	return s.sns.LookupSNSByWallets(wallets)
}

func proposerSNSForDisplay(p *model.Proposal, snsByWallet map[string]string) *string {
	return ResolveSNSNameForDisplay(p.ProposerSNS, p.ProposerWallet, snsByWallet)
}

func toVoteItem(v *model.ProposalVote, snsByWallet map[string]string) ProposalVoteItem {
	return ProposalVoteItem{
		ID:          v.ID,
		SNSName:     voteSNSForDisplay(v, snsByWallet),
		VoterWallet: v.VoterWallet,
		VoteResult:  v.VoteResult,
		VoteWeight:  v.VoteWeight,
		VotedAt:     formatTimeUTC(v.VotedAt),
	}
}

func voteSNSForDisplay(v *model.ProposalVote, snsByWallet map[string]string) *string {
	return ResolveSNSNameForDisplay(v.SNSName, v.VoterWallet, snsByWallet)
}

func buildCommentTree(rows []model.ProposalComment, snsByWallet map[string]string) []ProposalCommentItem {
	if len(rows) == 0 {
		return []ProposalCommentItem{}
	}

	ids := make(map[int64]struct{}, len(rows))
	for i := range rows {
		ids[rows[i].ID] = struct{}{}
	}

	byParent := make(map[int64][]ProposalCommentItem)
	for i := range rows {
		item := toCommentItem(&rows[i], snsByWallet)
		parentID := int64(0)
		if rows[i].ParentID != nil && *rows[i].ParentID > 0 {
			if _, ok := ids[*rows[i].ParentID]; ok {
				parentID = *rows[i].ParentID
			}
		}
		byParent[parentID] = append(byParent[parentID], item)
	}

	var attach func(parentID int64) []ProposalCommentItem
	attach = func(parentID int64) []ProposalCommentItem {
		items := byParent[parentID]
		sortCommentItems(items)
		for i := range items {
			items[i].Replies = attach(items[i].ID)
		}
		return items
	}
	return attach(0)
}

func sortCommentItems(items []ProposalCommentItem) {
	sort.Slice(items, func(i, j int) bool {
		if items[i].CommentedAt != items[j].CommentedAt {
			return items[i].CommentedAt < items[j].CommentedAt
		}
		return items[i].ID < items[j].ID
	})
}

func parseCommentContent(raw string) string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return raw
	}
	if !strings.HasPrefix(raw, "[") && !strings.HasPrefix(raw, "{") {
		return raw
	}

	var ops []struct {
		Insert json.RawMessage `json:"insert"`
	}
	if err := json.Unmarshal([]byte(raw), &ops); err != nil {
		return raw
	}
	if len(ops) == 0 {
		return raw
	}

	var b strings.Builder
	for _, op := range ops {
		if len(op.Insert) == 0 {
			continue
		}
		var text string
		if err := json.Unmarshal(op.Insert, &text); err == nil {
			b.WriteString(text)
		}
	}
	if b.Len() == 0 {
		return raw
	}
	return strings.TrimSpace(b.String())
}

func toCommentItem(c *model.ProposalComment, snsByWallet map[string]string) ProposalCommentItem {
	return ProposalCommentItem{
		ID:           c.ID,
		SNSName:      commentSNSForDisplay(c, snsByWallet),
		AuthorWallet: c.AuthorWallet,
		ParentID:     c.ParentID,
		Content:      parseCommentContent(c.Content),
		IsReject:     c.IsReject,
		CommentedAt:  formatTimeUTC(c.CommentedAt),
	}
}

func commentSNSForDisplay(c *model.ProposalComment, snsByWallet map[string]string) *string {
	return ResolveSNSNameForDisplay(c.SNSName, c.AuthorWallet, snsByWallet)
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
