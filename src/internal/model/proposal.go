package model

import "time"

const (
	ProposalTypeHallRegular   = "市政厅联席会议常规提案"
	ProposalTypeNodeConsensus = "节点共识大会常规提案"

	ProposalStatusDraft      = "draft"
	ProposalStatusClosed     = "closed"
	ProposalStatusPublicity  = "publicity"
	ProposalStatusPublished  = "published"
	ProposalStatusVotePassed = "vote_passed"
	ProposalStatusVoteFailed = "vote_failed"
	ProposalStatusExecuted   = "executed"
	ProposalStatusRejected   = "rejected"
	ProposalStatusWithdrawn  = "withdrawn"
	ProposalStatusError      = "error" // legacy import only; excluded from public list

	ProposalResolutionStandard = "standard"
	ProposalResolutionMetaRule = "meta_rule"

	ProposalVoteTypeStandard = 1
	ProposalVoteTypeMulti    = 2
)

type Proposal struct {
	ID                  int64      `gorm:"primaryKey"`
	LegacyID            *int64     `gorm:"column:legacy_id"`
	LegacyRecordID      *string    `gorm:"column:legacy_record_id;size:128"`
	SIP                 int        `gorm:"column:sip;not null;default:0"`
	AssociateProposalID *int64     `gorm:"column:associate_proposal_id"`
	Type                string     `gorm:"size:128;not null"`
	Title               string     `gorm:"size:512;not null"`
	Content             []byte     `gorm:"type:jsonb;not null"`
	ProposerSNS         *string    `gorm:"column:proposer_sns;size:64"`
	ProposerWallet      string     `gorm:"column:proposer_wallet;size:42;not null"`
	Status              string     `gorm:"size:32;not null"`
	VoteType            int16      `gorm:"column:vote_type;not null;default:0"`
	VoteOptions         []byte     `gorm:"column:vote_options;type:jsonb"`
	ResolutionRule      string     `gorm:"column:resolution_rule;size:32;not null;default:standard"`
	PublishedAt         *time.Time `gorm:"column:published_at"`
	PublicityEndAt      *time.Time `gorm:"column:publicity_end_at"`
	VoteEndAt           *time.Time `gorm:"column:vote_end_at"`
	PublicitySeconds    *int       `gorm:"column:publicity_seconds"`
	VoteDurationSeconds *int       `gorm:"column:vote_duration_seconds"`
	CreatedAt           time.Time  `gorm:"column:created_at;not null"`
	UpdatedAt           time.Time  `gorm:"column:updated_at;not null"`
}

func (Proposal) TableName() string { return "proposal" }

type ProposalVote struct {
	ID          int64     `gorm:"primaryKey"`
	ProposalID  int64     `gorm:"column:proposal_id;not null;index"`
	SNSName     *string   `gorm:"column:sns_name;size:64"`
	VoterWallet string    `gorm:"column:voter_wallet;size:42;not null"`
	VoteResult  string    `gorm:"column:vote_result;size:128;not null"`
	VoteWeight  int       `gorm:"column:vote_weight;not null;default:1"`
	VotedAt     time.Time `gorm:"column:voted_at;not null"`
	LegacyID    *int64    `gorm:"column:legacy_id;uniqueIndex"`
}

func (ProposalVote) TableName() string { return "proposal_vote" }

type ProposalComment struct {
	ID           int64     `gorm:"primaryKey"`
	ProposalID   int64     `gorm:"column:proposal_id;not null;index"`
	SNSName      *string   `gorm:"column:sns_name;size:64"`
	AuthorWallet string    `gorm:"column:author_wallet;size:42;not null"`
	ParentID     *int64    `gorm:"column:parent_id"`
	Content      string    `gorm:"type:text;not null"`
	IsReject     bool      `gorm:"column:is_reject;not null;default:false"`
	CommentedAt  time.Time `gorm:"column:commented_at;not null"`
	LegacyID     *int64    `gorm:"column:legacy_id;uniqueIndex"`
}

func (ProposalComment) TableName() string { return "proposal_comment" }
