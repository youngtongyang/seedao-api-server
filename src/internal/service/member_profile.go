package service

import (
	"context"
	"errors"
	"sync"

	"gorm.io/gorm"

	"github.com/seedao/seedao-api-server/internal/model"
	"github.com/seedao/seedao-api-server/internal/pkg/snsutil"
)

type MemberProfileService struct {
	db *gorm.DB
}

func NewMemberProfileService(db *gorm.DB) *MemberProfileService {
	return &MemberProfileService{db: db}
}

type MemberAssetItem struct {
	AssetName        string `json:"asset_name"`
	DealtAmount      string `json:"dealt_amount"`
	ProcessingAmount string `json:"processing_amount"`
	Claimed          bool   `json:"claimed"`
	SourceCreateTS   *int64 `json:"source_create_ts,omitempty"`
}

type MemberSeedItem struct {
	TokenID      string `json:"token_id"`
	ContractAddr string `json:"contract_addr"`
	ContractType string `json:"contract_type"`
	TokenAmount  string `json:"token_amount"`
	ImageURL     string `json:"image_url"`
}

type MemberSBTItem struct {
	TokenID        string `json:"token_id"`
	ContractAddr   string `json:"contract_addr"`
	ContractType   string `json:"contract_type"`
	TokenAmount    string `json:"token_amount"`
	CollectionName string `json:"collection_name"`
	DisplayName    string `json:"display_name"`
	Description    string `json:"description"`
	ImageURL       string `json:"image_url"`
}

// MemberProfileResult is the aggregated 1.0 community profile for one SNS.
type MemberProfileResult struct {
	SNSName        string            `json:"sns_name"`
	Bio            string            `json:"bio"`
	Email          string            `json:"email"`
	TwitterProfile string            `json:"twitter_profile"`
	Mirror         string            `json:"mirror"`
	SCRAmount      string            `json:"scr_amount,omitempty"`
	Roles          []string          `json:"roles"`
	Assets         []MemberAssetItem `json:"assets"`
	SEE            *MemberAssetItem  `json:"see,omitempty"`
	WANG           *MemberAssetItem  `json:"wang,omitempty"`
	Seeds          []MemberSeedItem  `json:"seeds"`
	SBTs           []MemberSBTItem   `json:"sbts"`
}

// GetBySNS loads profile / assets / roles / seeds / sbts concurrently by sns_name.
func (s *MemberProfileService) GetBySNS(ctx context.Context, snsName string) (*MemberProfileResult, error) {
	snsName = snsutil.Normalize(snsName)
	if err := snsutil.ValidateFormat(snsName); err != nil {
		return nil, ErrInvalidParam
	}

	var (
		profile    model.UserProfile
		profileErr error
		hasProfile bool

		assets    []model.UserAsset
		assetsErr error

		roles    []model.UserRole
		rolesErr error

		seeds    []model.UserSeed
		seedsErr error

		sbts    []model.UserSBT
		sbtsErr error
	)

	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		defer wg.Done()
		err := s.db.WithContext(ctx).Where("sns_name = ?", snsName).First(&profile).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
		if err != nil {
			profileErr = err
			return
		}
		hasProfile = true
	}()

	go func() {
		defer wg.Done()
		assetsErr = s.db.WithContext(ctx).
			Where("sns_name = ?", snsName).
			Order("asset_name ASC").
			Find(&assets).Error
	}()

	go func() {
		defer wg.Done()
		rolesErr = s.db.WithContext(ctx).
			Where("sns_name = ?", snsName).
			Order("role ASC").
			Find(&roles).Error
	}()

	go func() {
		defer wg.Done()
		seedsErr = s.db.WithContext(ctx).
			Where("sns_name = ?", snsName).
			Order("id ASC").
			Find(&seeds).Error
	}()

	go func() {
		defer wg.Done()
		sbtsErr = s.db.WithContext(ctx).
			Where("sns_name = ?", snsName).
			Order("id ASC").
			Find(&sbts).Error
	}()

	wg.Wait()

	for _, err := range []error{profileErr, assetsErr, rolesErr, seedsErr, sbtsErr} {
		if err != nil {
			return nil, err
		}
	}

	out := &MemberProfileResult{
		SNSName: snsName,
		Roles:   make([]string, 0, len(roles)),
		Assets:  make([]MemberAssetItem, 0, len(assets)),
		Seeds:   make([]MemberSeedItem, 0, len(seeds)),
		SBTs:    make([]MemberSBTItem, 0, len(sbts)),
	}
	if hasProfile {
		out.Bio = profile.Bio
		out.Email = profile.Email
		out.TwitterProfile = profile.TwitterProfile
		out.Mirror = profile.Mirror
		if profile.SCRAmount != nil {
			out.SCRAmount = *profile.SCRAmount
		}
	}

	for _, r := range roles {
		out.Roles = append(out.Roles, r.Role)
	}
	for _, a := range assets {
		item := MemberAssetItem{
			AssetName:        a.AssetName,
			DealtAmount:      a.DealtAmount,
			ProcessingAmount: a.ProcessingAmount,
			Claimed:          a.Claimed,
			SourceCreateTS:   a.SourceCreateTS,
		}
		out.Assets = append(out.Assets, item)
		switch a.AssetName {
		case "SEE":
			cp := item
			out.SEE = &cp
		case "WANG":
			cp := item
			out.WANG = &cp
		}
	}
	for _, seed := range seeds {
		out.Seeds = append(out.Seeds, MemberSeedItem{
			TokenID:      seed.TokenID,
			ContractAddr: seed.ContractAddr,
			ContractType: seed.ContractType,
			TokenAmount:  seed.TokenAmount,
			ImageURL:     preferImageURL(seed.ImageURIOSS, seed.ImageURISrc),
		})
	}
	for _, sbt := range sbts {
		out.SBTs = append(out.SBTs, MemberSBTItem{
			TokenID:        sbt.TokenID,
			ContractAddr:   sbt.ContractAddr,
			ContractType:   sbt.ContractType,
			TokenAmount:    sbt.TokenAmount,
			CollectionName: sbt.CollectionName,
			DisplayName:    sbt.DisplayName,
			Description:    sbt.Description,
			ImageURL:       preferImageURL(sbt.ImageURIOSS, sbt.ImageURISrc),
		})
	}

	return out, nil
}

func preferImageURL(oss, src string) string {
	if oss != "" {
		return oss
	}
	return src
}
