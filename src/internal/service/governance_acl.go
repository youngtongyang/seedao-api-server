package service

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/seedao/seedao-api-server/internal/config"
	"github.com/seedao/seedao-api-server/internal/model"
)

type governanceACL struct {
	cfg config.GovernanceConfig
}

func newGovernanceACL(cfg config.GovernanceConfig) governanceACL {
	return governanceACL{cfg: cfg}
}

func (g governanceACL) normalize(name string) string {
	return strings.ToLower(strings.TrimSpace(name))
}

func (g governanceACL) inList(name string, list []string) bool {
	n := g.normalize(name)
	for _, item := range list {
		if g.normalize(item) == n {
			return true
		}
	}
	return false
}

func (g governanceACL) canCreateType(snsName, proposalType string) bool {
	switch proposalTypeFamily(proposalType) {
	case "hall":
		return g.inList(snsName, g.cfg.HallProposalSNS)
	case "node":
		return g.inList(snsName, g.cfg.NodeConsensusProposalSNS)
	default:
		return false
	}
}

func (g governanceACL) canVote(proposalType, snsName string) bool {
	switch proposalType {
	case model.ProposalTypeHallRegular:
		return g.inList(snsName, g.cfg.HallVoterSNS)
	case model.ProposalTypeNodeConsensus:
		return g.inList(snsName, g.cfg.NodeVoterSNS)
	default:
		// legacy imported types: node voters by default
		return g.inList(snsName, g.cfg.NodeVoterSNS)
	}
}

func defaultVoteOptionsJSON() []byte {
	opts := []map[string]string{
		{"label": "同意", "value": "1"},
		{"label": "反对", "value": "-1"},
		{"label": "弃权", "value": "0"},
	}
	b, _ := json.Marshal(opts)
	return b
}

type VoteTally struct {
	TotalVotes       int            `json:"total_votes"`
	Participation    float64        `json:"participation_rate"`
	ByResult         map[string]int `json:"by_result"`
	ApproveCount     int            `json:"approve_count"`
	RejectCount      int            `json:"reject_count"`
	AbstainCount     int            `json:"abstain_count"`
	ValidVotes       int            `json:"valid_votes"`
	ApproveRate      float64        `json:"approve_rate"`
	ResolutionRule   string         `json:"resolution_rule"`
	Resolved         bool           `json:"resolved"`
	Passed           bool           `json:"passed"`
	ResolutionReason string         `json:"resolution_reason,omitempty"`
}

func evaluateResolution(rule string, electorateTotal int, tally VoteTally) VoteTally {
	tally.ResolutionRule = rule
	if electorateTotal <= 0 {
		tally.Resolved = false
		tally.ResolutionReason = "voting electorate not configured"
		return tally
	}

	tally.Participation = float64(tally.TotalVotes) / float64(electorateTotal)
	tally.ValidVotes = tally.ApproveCount + tally.RejectCount
	if tally.ValidVotes > 0 {
		tally.ApproveRate = float64(tally.ApproveCount) / float64(tally.ValidVotes)
	}

	switch rule {
	case model.ProposalResolutionMetaRule:
		// 元规则 / 市政厅：选举人总数 2/3 以上赞成
		required := twoThirdsRequired(electorateTotal)
		tally.Resolved = true
		tally.Passed = tally.ApproveCount >= required
		if !tally.Passed {
			tally.ResolutionReason = "meta_rule requires two-thirds of electorate to approve"
		}
	default:
		// 节点共识常规：参与 > 1/2 选举人总数，且赞成 > 有效票 1/2（第三章 3.2.4）
		participationOK := tally.TotalVotes > electorateTotal/2
		approveOK := tally.ValidVotes > 0 && tally.ApproveCount > tally.ValidVotes/2
		tally.Resolved = true
		tally.Passed = participationOK && approveOK
		if !participationOK {
			tally.ResolutionReason = "participation must exceed half of electorate"
		} else if !approveOK {
			tally.ResolutionReason = "approve votes must exceed half of valid votes (excluding abstain)"
		}
	}
	return tally
}

func twoThirdsRequired(total int) int {
	return (total*2 + 2) / 3
}

func classifyVoteResult(result string) string {
	r := strings.TrimSpace(strings.ToLower(result))
	if r == "1" || strings.Contains(r, "同意") || r == "yes" || r == "approve" {
		return "approve"
	}
	if r == "-1" || strings.Contains(r, "反对") || r == "no" || r == "reject" {
		return "reject"
	}
	return "abstain"
}

func isEditableStatus(status string) bool {
	return status == model.ProposalStatusDraft
}

func isHallOrNodeCreateType(t string) bool {
	return isCreatableProposalType(t)
}

func formatTimeUTC(t time.Time) string {
	return t.UTC().Format(time.RFC3339)
}

func formatOptionalTime(t *time.Time) *string {
	if t == nil {
		return nil
	}
	s := formatTimeUTC(*t)
	return &s
}
