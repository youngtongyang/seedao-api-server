package legacyimport

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

type apiUser struct {
	ID              int64     `json:"id"`
	Bio             string    `json:"bio"`
	Email           string    `json:"email"`
	TwitterProfile  string    `json:"twitter_profile"`
	Mirror          string    `json:"mirror"`
	Assets          []apiAsset `json:"assets"`
	SP              apiSP     `json:"sp"`
}

type apiAsset struct {
	AssetName        string `json:"asset_name"`
	DealtAmount      string `json:"dealt_amount"`
	ProcessingAmount string `json:"processing_amount"`
	Claimed          bool   `json:"claimed"`
	CreateTS         int64  `json:"create_ts"`
}

type apiSP struct {
	Roles []string  `json:"roles"`
	SCR   apiSCR    `json:"scr"`
	Seed  []apiSeed `json:"seed"`
	SBT   []apiSBT  `json:"sbt"`
}

type apiSCR struct {
	Amount string `json:"amount"`
}

type apiSeed struct {
	TokenID      string `json:"token_id"`
	ContractAddr string `json:"contract_addr"`
	ContractType string `json:"contract_type"`
	ImageURI     string `json:"image_uri"`
	TokenAmount  string `json:"token_amount"`
}

type apiSBT struct {
	TokenID        string         `json:"token_id"`
	ContractAddr   string         `json:"contract_addr"`
	ContractType   string         `json:"contract_type"`
	ImageURI       string         `json:"image_uri"`
	TokenAmount    string         `json:"token_amount"`
	CollectionName string         `json:"collection_name"`
	Name           string         `json:"name"`
	Metadata       apiSBTMetadata `json:"metadata"`
}

type apiSBTMetadata struct {
	Name        string `json:"name"`
	Headline    string `json:"headline"`
	Description string `json:"description"`
}

type apiRewardsData struct {
	Page  int               `json:"page"`
	Size  int               `json:"size"`
	Total int               `json:"total"`
	Rows  []apiRewardRow    `json:"rows"`
}

type apiRewardRow struct {
	ApplicationID    int64  `json:"application_id"`
	SeasonName       string `json:"season_name"`
	EntityName       string `json:"entity_name"`
	EntityType       string `json:"entity_type"`
	AssetName        string `json:"asset_name"`
	Amount           string `json:"amount"`
	BudgetSource     string `json:"budget_source"`
	Status           string `json:"status"`
	TargetUserWallet string `json:"target_user_wallet"`
	ApplicantWallet  string `json:"applicant_wallet"`
	CreateTS         int64  `json:"create_ts"`
	UpdateTS         int64  `json:"update_ts"`
}

func resolveSBTNames(sbt apiSBT) (displayName, description string) {
	description = strings.TrimSpace(sbt.Metadata.Description)
	displayName = firstNonEmpty(
		sbt.Metadata.Name,
		sbt.Metadata.Headline,
		sbt.Metadata.Description,
		sbt.Name,
	)
	return displayName, description
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if s := strings.TrimSpace(v); s != "" {
			return s
		}
	}
	return ""
}

func numericPtr(s string) *string {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	return &s
}

func numericOrZero(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return "0"
	}
	return s
}

func int64Ptr(v int64) *int64 {
	if v == 0 {
		return nil
	}
	x := v
	return &x
}

func snsPtr(name string) *string {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil
	}
	return &name
}

func sha256Hex(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

func extFromContentType(ct string) string {
	switch strings.ToLower(strings.TrimSpace(ct)) {
	case "image/png":
		return ".png"
	case "image/jpeg", "image/jpg":
		return ".jpg"
	case "image/gif":
		return ".gif"
	case "image/webp":
		return ".webp"
	case "image/svg+xml":
		return ".svg"
	default:
		return ""
	}
}
