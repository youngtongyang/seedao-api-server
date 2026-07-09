package service

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/seedao/seedao-api-server/internal/model"
)

// ProposalParamError carries a human-readable validation reason for API clients.
type ProposalParamError struct {
	Reason string
}

func (e *ProposalParamError) Error() string {
	if e.Reason != "" {
		return e.Reason
	}
	return "invalid proposal parameter"
}

func (e *ProposalParamError) Is(target error) bool {
	return target == ErrInvalidParam
}

// ProposalForbiddenError carries a human-readable ACL denial reason.
type ProposalForbiddenError struct {
	Reason string
}

func (e *ProposalForbiddenError) Error() string {
	if e.Reason != "" {
		return e.Reason
	}
	return ErrProposalForbidden.Error()
}

func (e *ProposalForbiddenError) Is(target error) bool {
	return target == ErrProposalForbidden
}

var creatableProposalTypeAPIKeys = []string{
	"town_hall_meeting",
	"node_consensus_special",
	"node_consensus_rule",
}

var createProposalTypeAliases = map[string]string{
	"node_consensus_rule":    "节点共识大会规则提案",
	"node_consensus_special": "节点共识大会特别提案",
	"town_hall_meeting":      model.ProposalTypeHallRegular,
	model.ProposalTypeHallRegular:   model.ProposalTypeHallRegular,
	model.ProposalTypeNodeConsensus: model.ProposalTypeNodeConsensus,
	"节点共识大会规则提案":                  "节点共识大会规则提案",
	"节点共识大会特别提案":                  "节点共识大会特别提案",
	"节点共识大会提案":                    "节点共识大会规则提案",
}

var creatableProposalTypes = map[string]struct{}{
	model.ProposalTypeHallRegular:   {},
	model.ProposalTypeNodeConsensus: {},
	"节点共识大会规则提案":                  {},
	"节点共识大会特别提案":                  {},
}

var resolutionRuleAliases = map[string]string{
	"":                              model.ProposalResolutionStandard,
	model.ProposalResolutionStandard: model.ProposalResolutionStandard,
	model.ProposalResolutionMetaRule: model.ProposalResolutionMetaRule,
	"town_hall_standard":            model.ProposalResolutionMetaRule,
	"node_consensus_special":        model.ProposalResolutionStandard,
	"p1_standard":                   model.ProposalResolutionStandard,
	"p2_standard":                   model.ProposalResolutionStandard,
	"p3_standard":                   model.ProposalResolutionStandard,
}

func normalizeCreateProposalType(raw string) (string, error) {
	t := strings.TrimSpace(raw)
	if t == "" {
		return "", &ProposalParamError{Reason: "proposal type is required"}
	}
	if mapped, ok := createProposalTypeAliases[t]; ok {
		t = mapped
	}
	if _, ok := creatableProposalTypes[t]; !ok {
		return "", &ProposalParamError{
			Reason: fmt.Sprintf("unsupported proposal type %q; allowed: node_consensus_rule, node_consensus_special, town_hall_meeting", raw),
		}
	}
	return t, nil
}

func defaultResolutionRule(proposalType string) string {
	switch proposalType {
	case model.ProposalTypeHallRegular, "节点共识大会规则提案":
		return model.ProposalResolutionMetaRule
	default:
		return model.ProposalResolutionStandard
	}
}

func normalizeResolutionRule(raw string) (string, error) {
	key := strings.TrimSpace(strings.ToLower(raw))
	if mapped, ok := resolutionRuleAliases[key]; ok {
		return mapped, nil
	}
	if mapped, ok := resolutionRuleAliases[raw]; ok {
		return mapped, nil
	}
	return "", &ProposalParamError{
		Reason: fmt.Sprintf("unsupported resolution_rule %q; allowed: standard, meta_rule (aliases: town_hall_standard, node_consensus_special)", raw),
	}
}

// NormalizeWireJSON decodes frontend integer[] wire bytes or passes through valid JSON.
func NormalizeWireJSON(raw json.RawMessage, fieldName string) (json.RawMessage, error) {
	if len(raw) == 0 {
		return nil, &ProposalParamError{Reason: fmt.Sprintf("%s is required", fieldName)}
	}
	var wireBytes []int
	if err := json.Unmarshal(raw, &wireBytes); err == nil && len(wireBytes) > 0 {
		buf := make([]byte, len(wireBytes))
		for i, n := range wireBytes {
			if n < 0 || n > 255 {
				return nil, &ProposalParamError{Reason: fmt.Sprintf("%s wire bytes out of range", fieldName)}
			}
			buf[i] = byte(n)
		}
		if !json.Valid(buf) {
			return nil, &ProposalParamError{Reason: fmt.Sprintf("%s wire bytes are not valid JSON", fieldName)}
		}
		return json.RawMessage(buf), nil
	}
	if json.Valid(raw) {
		return raw, nil
	}
	return nil, &ProposalParamError{
		Reason: fmt.Sprintf("%s must be valid JSON or integer[] wire bytes", fieldName),
	}
}

// NormalizeProposalContent accepts JSON object content or frontend integer[] wire bytes.
func NormalizeProposalContent(raw json.RawMessage) (json.RawMessage, error) {
	if len(raw) == 0 {
		return nil, &ProposalParamError{Reason: "proposal content is required"}
	}
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(raw, &obj); err == nil && len(obj) > 0 {
		return raw, nil
	}
	return NormalizeWireJSON(raw, "proposal content")
}

// NormalizeVoteOptions accepts label/value objects or a string[] from the frontend wire codec.
func NormalizeVoteOptions(raw json.RawMessage) ([]byte, error) {
	if len(raw) == 0 {
		return defaultVoteOptionsJSON(), nil
	}
	normalized, err := NormalizeWireJSON(raw, "vote_options")
	if err != nil {
		return nil, err
	}
	var labels []string
	if err := json.Unmarshal(normalized, &labels); err == nil && len(labels) > 0 {
		opts := make([]map[string]string, len(labels))
		for i, label := range labels {
			opts[i] = map[string]string{
				"label": label,
				"value": mapVoteLabelToValue(label, i),
			}
		}
		b, _ := json.Marshal(opts)
		return b, nil
	}
	var objs []map[string]string
	if err := json.Unmarshal(normalized, &objs); err == nil && len(objs) > 0 {
		b, _ := json.Marshal(objs)
		return b, nil
	}
	return nil, &ProposalParamError{
		Reason: "vote_options must be a JSON array of labels or {label,value} objects",
	}
}

func mapVoteLabelToValue(label string, index int) string {
	switch strings.TrimSpace(label) {
	case "同意":
		return "1"
	case "反对":
		return "-1"
	case "弃权":
		return "0"
	default:
		switch index {
		case 0:
			return "1"
		case 1:
			return "-1"
		default:
			return "0"
		}
	}
}

// resolveVoteType prefers the explicit API field, then vote_type embedded in content JSON.
func resolveVoteType(explicit int16, content json.RawMessage) int16 {
	if explicit == model.ProposalVoteTypeStandard || explicit == model.ProposalVoteTypeMulti {
		return explicit
	}
	var payload struct {
		VoteType int16 `json:"vote_type"`
	}
	if err := json.Unmarshal(content, &payload); err == nil {
		if payload.VoteType == model.ProposalVoteTypeMulti {
			return model.ProposalVoteTypeMulti
		}
		if payload.VoteType == model.ProposalVoteTypeStandard {
			return model.ProposalVoteTypeStandard
		}
	}
	return model.ProposalVoteTypeStandard
}

func proposalTypeFamily(proposalType string) string {
	switch proposalType {
	case model.ProposalTypeHallRegular:
		return "hall"
	case model.ProposalTypeNodeConsensus, "节点共识大会规则提案", "节点共识大会特别提案":
		return "node"
	default:
		return ""
	}
}

func isCreatableProposalType(t string) bool {
	_, ok := creatableProposalTypes[t]
	return ok
}
