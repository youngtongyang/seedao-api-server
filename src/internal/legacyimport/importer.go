package legacyimport

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/seedao/seedao-api-server/internal/config"
	"github.com/seedao/seedao-api-server/internal/model"
	"github.com/seedao/seedao-api-server/internal/pkg/aliyunoss"
	"github.com/seedao/seedao-api-server/internal/pkg/phone"
)

// ErrSkipNoImage means there is no OpenSea image and no usable HTTPS source.
var ErrSkipNoImage = errors.New("skip: no downloadable image")

type Importer struct {
	db         *gorm.DB
	httpClient *http.Client
	oss        aliyunoss.Uploader
	cfg        config.V1ImportConfig

	openSeaMu   sync.Mutex
	openSeaNext time.Time
}

func New(db *gorm.DB, ossUploader aliyunoss.Uploader, cfg config.V1ImportConfig) *Importer {
	if ossUploader == nil {
		ossUploader = &aliyunoss.Disabled{}
	}
	return &Importer{
		db: db,
		httpClient: &http.Client{
			Timeout: 45 * time.Second,
		},
		oss: ossUploader,
		cfg: cfg,
	}
}

type registryRow struct {
	Wallet  string `gorm:"column:wallet"`
	SNSName string `gorm:"column:sns_name"`
}

// ImportProfiles pulls 1.0 user payloads for every sns_chain_registry row.
func (im *Importer) ImportProfiles(ctx context.Context, uploadOSS bool) error {
	var rows []registryRow
	if err := im.db.Model(&model.SNSChainRegistry{}).
		Select("wallet", "sns_name").
		Find(&rows).Error; err != nil {
		return err
	}
	total := len(rows)
	log.Printf("profiles: %d registry rows", total)
	if total == 0 {
		return nil
	}

	concurrency := im.cfg.Concurrency
	if concurrency <= 0 {
		concurrency = 3
	}

	sem := make(chan struct{}, concurrency)
	var wg sync.WaitGroup
	var okCount, errCount atomic.Int64

	for i, row := range rows {
		select {
		case <-ctx.Done():
			wg.Wait()
			return ctx.Err()
		default:
		}

		wg.Add(1)
		sem <- struct{}{}
		go func(idx int, r registryRow) {
			defer wg.Done()
			defer func() { <-sem }()

			if err := im.importOneProfile(ctx, r, uploadOSS); err != nil {
				errCount.Add(1)
				log.Printf("profiles [%d/%d] %s (%s): %v", idx+1, total, r.SNSName, r.Wallet, err)
				return
			}
			okCount.Add(1)
			if (idx+1)%50 == 0 || idx+1 == total {
				log.Printf("profiles progress %d/%d ok=%d err=%d",
					idx+1, total, okCount.Load(), errCount.Load())
			}
		}(i, row)
	}
	wg.Wait()
	log.Printf("profiles done: ok=%d err=%d", okCount.Load(), errCount.Load())
	return nil
}

func (im *Importer) importOneProfile(ctx context.Context, row registryRow, uploadOSS bool) error {
	wallet := phone.NormalizeWallet(row.Wallet)
	snsName := strings.TrimSpace(row.SNSName)
	if wallet == "" || snsName == "" {
		return fmt.Errorf("empty wallet or sns_name")
	}

	payload, err := im.fetchUser(ctx, wallet)
	if err != nil {
		return err
	}
	if payload == nil {
		return nil
	}

	now := time.Now()
	profile := model.UserProfile{
		SNSName:        snsName,
		Bio:            payload.Bio,
		Email:          payload.Email,
		TwitterProfile: payload.TwitterProfile,
		Mirror:         payload.Mirror,
		SCRAmount:      numericPtr(payload.SP.SCR.Amount),
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	return im.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "sns_name"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"bio", "email", "twitter_profile", "mirror", "scr_amount", "updated_at",
			}),
		}).Create(&profile).Error; err != nil {
			return err
		}

		for _, asset := range payload.Assets {
			name := strings.TrimSpace(asset.AssetName)
			if name == "" {
				continue
			}
			ua := model.UserAsset{
				SNSName:          snsName,
				AssetName:        name,
				DealtAmount:      numericOrZero(asset.DealtAmount),
				ProcessingAmount: numericOrZero(asset.ProcessingAmount),
				Claimed:          asset.Claimed,
				SourceCreateTS:   int64Ptr(asset.CreateTS),
				CreatedAt:        now,
				UpdatedAt:        now,
			}
			if err := tx.Clauses(clause.OnConflict{
				Columns: []clause.Column{{Name: "sns_name"}, {Name: "asset_name"}},
				DoUpdates: clause.AssignmentColumns([]string{
					"dealt_amount", "processing_amount", "claimed", "source_create_ts", "updated_at",
				}),
			}).Create(&ua).Error; err != nil {
				return err
			}
		}

		if err := tx.Where("sns_name = ?", snsName).Delete(&model.UserRole{}).Error; err != nil {
			return err
		}
		seenRole := map[string]struct{}{}
		for _, role := range payload.SP.Roles {
			role = strings.TrimSpace(role)
			if role == "" {
				continue
			}
			if _, ok := seenRole[role]; ok {
				continue
			}
			seenRole[role] = struct{}{}
			if err := tx.Create(&model.UserRole{
				SNSName:   snsName,
				Role:      role,
				CreatedAt: now,
			}).Error; err != nil {
				return err
			}
		}

		for _, seed := range payload.SP.Seed {
			src := strings.TrimSpace(seed.ImageURI)
			ossURL := ""
			if uploadOSS {
				var err error
				ossURL, err = im.ensureOSSURL(ctx, tx, src, path.Join("seed", strings.ToLower(seed.ContractAddr), seed.TokenID), seed.ContractAddr, seed.TokenID)
				if err != nil {
					log.Printf("oss seed %s/%s: %v", snsName, seed.TokenID, err)
				}
			}
			us := model.UserSeed{
				SNSName:      snsName,
				TokenID:      seed.TokenID,
				ContractAddr: strings.ToLower(seed.ContractAddr),
				ContractType: seed.ContractType,
				TokenAmount:  seed.TokenAmount,
				ImageURISrc:  src,
				ImageURIOSS:  ossURL,
				CreatedAt:    now,
				UpdatedAt:    now,
			}
			if err := tx.Clauses(clause.OnConflict{
				Columns: []clause.Column{{Name: "sns_name"}, {Name: "contract_addr"}, {Name: "token_id"}},
				DoUpdates: clause.Assignments(map[string]interface{}{
					"contract_type": us.ContractType,
					"token_amount":  us.TokenAmount,
					"image_uri_src": us.ImageURISrc,
					"updated_at":    now,
					"image_uri_oss": gorm.Expr(
						"CASE WHEN EXCLUDED.image_uri_oss <> '' THEN EXCLUDED.image_uri_oss ELSE user_seeds.image_uri_oss END",
					),
				}),
			}).Create(&us).Error; err != nil {
				return err
			}
		}

		for _, sbt := range payload.SP.SBT {
			src := strings.TrimSpace(sbt.ImageURI)
			ossURL := ""
			if uploadOSS {
				var err error
				ossURL, err = im.ensureOSSURL(ctx, tx, src, path.Join("sbt", strings.ToLower(sbt.ContractAddr), sbt.TokenID), sbt.ContractAddr, sbt.TokenID)
				if err != nil {
					log.Printf("oss sbt %s/%s: %v", snsName, sbt.TokenID, err)
				}
			}
			display, desc := resolveSBTNames(sbt)
			us := model.UserSBT{
				SNSName:        snsName,
				TokenID:        sbt.TokenID,
				ContractAddr:   strings.ToLower(sbt.ContractAddr),
				ContractType:   sbt.ContractType,
				TokenAmount:    sbt.TokenAmount,
				CollectionName: sbt.CollectionName,
				DisplayName:    display,
				Description:    desc,
				ImageURISrc:    src,
				ImageURIOSS:    ossURL,
				CreatedAt:      now,
				UpdatedAt:      now,
			}
			if err := tx.Clauses(clause.OnConflict{
				Columns: []clause.Column{{Name: "sns_name"}, {Name: "contract_addr"}, {Name: "token_id"}},
				DoUpdates: clause.Assignments(map[string]interface{}{
					"contract_type":   us.ContractType,
					"token_amount":    us.TokenAmount,
					"collection_name": us.CollectionName,
					"display_name":    us.DisplayName,
					"description":     us.Description,
					"image_uri_src":   us.ImageURISrc,
					"updated_at":      now,
					"image_uri_oss": gorm.Expr(
						"CASE WHEN EXCLUDED.image_uri_oss <> '' THEN EXCLUDED.image_uri_oss ELSE user_sbts.image_uri_oss END",
					),
				}),
			}).Create(&us).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// ImportRewards paginates 1.0 new_reward applications.
func (im *Importer) ImportRewards(ctx context.Context) error {
	pageSize := im.cfg.PageSize
	if pageSize <= 0 {
		pageSize = 100
	}

	walletToSNS, err := im.loadRegistryMap()
	if err != nil {
		return err
	}

	page := 1
	var imported int
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		resp, err := im.fetchRewards(ctx, page, pageSize)
		if err != nil {
			return err
		}
		if len(resp.Rows) == 0 {
			break
		}

		now := time.Now()
		for _, row := range resp.Rows {
			targetWallet := phone.NormalizeWallet(row.TargetUserWallet)
			applicantWallet := phone.NormalizeWallet(row.ApplicantWallet)
			rec := model.RewardApplication{
				ApplicationID:    row.ApplicationID,
				SeasonName:       row.SeasonName,
				EntityName:       row.EntityName,
				EntityType:       row.EntityType,
				AssetName:        row.AssetName,
				Amount:           numericPtr(row.Amount),
				BudgetSource:     row.BudgetSource,
				Status:           row.Status,
				TargetUserWallet: targetWallet,
				TargetSNSName:    snsPtr(walletToSNS[targetWallet]),
				ApplicantWallet:  applicantWallet,
				ApplicantSNSName: snsPtr(walletToSNS[applicantWallet]),
				SourceCreateTS:   int64Ptr(row.CreateTS),
				SourceUpdateTS:   int64Ptr(row.UpdateTS),
				CreatedAt:        now,
				UpdatedAt:        now,
			}
			if err := im.db.WithContext(ctx).Clauses(clause.OnConflict{
				Columns: []clause.Column{{Name: "application_id"}},
				DoUpdates: clause.AssignmentColumns([]string{
					"season_name", "entity_name", "entity_type", "asset_name", "amount",
					"budget_source", "status", "target_user_wallet", "target_sns_name",
					"applicant_wallet", "applicant_sns_name", "source_create_ts",
					"source_update_ts", "updated_at",
				}),
			}).Create(&rec).Error; err != nil {
				return fmt.Errorf("upsert application %d: %w", row.ApplicationID, err)
			}
			imported++
		}

		log.Printf("rewards page %d size=%d imported_total=%d api_total=%d",
			page, len(resp.Rows), imported, resp.Total)

		if page*pageSize >= resp.Total {
			break
		}
		page++
	}

	log.Printf("rewards done: imported=%d", imported)
	return nil
}

// BackfillOSS uploads pending seed/sbt images and fills image_uri_oss.
// When opensea_api_key is set, images are resolved via OpenSea by contract+token_id (not IPFS first).
func (im *Importer) BackfillOSS(ctx context.Context, limit int) error {
	if !im.oss.Configured() {
		return fmt.Errorf("aliyun_oss is not configured")
	}
	if limit <= 0 {
		limit = 500
	}
	if im.openSeaConfigured() {
		log.Printf("oss backfill: using OpenSea API (chain=%s)", im.cfg.OpenSeaChain)
	} else {
		log.Printf("oss backfill: opensea_api_key empty, falling back to image_uri_src")
	}

	pendingWhere := "image_uri_oss = '' AND contract_addr <> '' AND token_id <> ''"
	// Prefer rows that already have an HTTPS source (faster fallback when OpenSea has no image).
	orderHTTPSFirst := "CASE WHEN image_uri_src LIKE 'https://%' OR image_uri_src LIKE 'http://%' THEN 0 ELSE 1 END"

	var seeds []model.UserSeed
	if err := im.db.WithContext(ctx).
		Where(pendingWhere).
		Order(orderHTTPSFirst).
		Order("id ASC").
		Limit(limit).
		Find(&seeds).Error; err != nil {
		return err
	}
	var seedOK, seedErr, seedSkip int
	for _, seed := range seeds {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		ossURL, err := im.ensureOSSURL(ctx, im.db, seed.ImageURISrc, path.Join("seed", seed.ContractAddr, seed.TokenID), seed.ContractAddr, seed.TokenID)
		if errors.Is(err, ErrSkipNoImage) {
			seedSkip++
			log.Printf("oss seed id=%d token=%s: %v", seed.ID, seed.TokenID, err)
			continue
		}
		if err != nil {
			seedErr++
			log.Printf("oss seed id=%d token=%s: %v", seed.ID, seed.TokenID, err)
			continue
		}
		if err := im.db.Model(&model.UserSeed{}).Where("id = ?", seed.ID).
			Updates(map[string]interface{}{"image_uri_oss": ossURL, "updated_at": time.Now()}).Error; err != nil {
			return err
		}
		seedOK++
	}

	var sbts []model.UserSBT
	if err := im.db.WithContext(ctx).
		Where(pendingWhere).
		Order(orderHTTPSFirst).
		Order("id ASC").
		Limit(limit).
		Find(&sbts).Error; err != nil {
		return err
	}
	var sbtOK, sbtErr, sbtSkip int
	for _, sbt := range sbts {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		ossURL, err := im.ensureOSSURL(ctx, im.db, sbt.ImageURISrc, path.Join("sbt", sbt.ContractAddr, sbt.TokenID), sbt.ContractAddr, sbt.TokenID)
		if errors.Is(err, ErrSkipNoImage) {
			sbtSkip++
			log.Printf("oss sbt id=%d token=%s: %v", sbt.ID, sbt.TokenID, err)
			continue
		}
		if err != nil {
			sbtErr++
			log.Printf("oss sbt id=%d token=%s: %v", sbt.ID, sbt.TokenID, err)
			continue
		}
		if err := im.db.Model(&model.UserSBT{}).Where("id = ?", sbt.ID).
			Updates(map[string]interface{}{"image_uri_oss": ossURL, "updated_at": time.Now()}).Error; err != nil {
			return err
		}
		sbtOK++
	}

	var seedPending, sbtPending, ossCached int64
	_ = im.db.Model(&model.UserSeed{}).Where(pendingWhere).Count(&seedPending).Error
	_ = im.db.Model(&model.UserSBT{}).Where(pendingWhere).Count(&sbtPending).Error
	_ = im.db.Model(&model.OSSObject{}).Count(&ossCached).Error

	log.Printf("oss backfill done: seed ok=%d skip=%d err=%d tried=%d | sbt ok=%d skip=%d err=%d tried=%d | pending seed=%d sbt=%d | oss_objects=%d",
		seedOK, seedSkip, seedErr, len(seeds), sbtOK, sbtSkip, sbtErr, len(sbts), seedPending, sbtPending, ossCached)
	return nil
}

func (im *Importer) loadRegistryMap() (map[string]string, error) {
	var rows []registryRow
	if err := im.db.Model(&model.SNSChainRegistry{}).
		Select("wallet", "sns_name").
		Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make(map[string]string, len(rows))
	for _, r := range rows {
		w := phone.NormalizeWallet(r.Wallet)
		if w == "" || r.SNSName == "" {
			continue
		}
		out[w] = strings.TrimSpace(r.SNSName)
	}
	return out, nil
}

func (im *Importer) fetchUser(ctx context.Context, wallet string) (*apiUser, error) {
	u := fmt.Sprintf("%s/v1/user/users?wallets=%s",
		strings.TrimRight(im.cfg.BaseURL, "/"), url.QueryEscape(wallet))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := im.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, 8<<20))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http %d: %s", resp.StatusCode, truncate(string(body), 200))
	}

	var envelope struct {
		Code int             `json:"code"`
		Msg  string          `json:"msg"`
		Data json.RawMessage `json:"data"`
	}
	if err := json.Unmarshal(body, &envelope); err != nil {
		return nil, err
	}
	if envelope.Code != 200 && envelope.Code != 0 {
		return nil, fmt.Errorf("api code=%d msg=%s", envelope.Code, envelope.Msg)
	}
	if len(envelope.Data) == 0 || string(envelope.Data) == "null" {
		return nil, nil
	}

	var users []apiUser
	if err := json.Unmarshal(envelope.Data, &users); err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, nil
	}
	return &users[0], nil
}

func (im *Importer) fetchRewards(ctx context.Context, page, size int) (*apiRewardsData, error) {
	u := fmt.Sprintf("%s/v1/applications/?page=%d&size=%d&sort_field=create_ts&sort_order=desc&type=new_reward",
		strings.TrimRight(im.cfg.BaseURL, "/"), page, size)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := im.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, 16<<20))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http %d: %s", resp.StatusCode, truncate(string(body), 200))
	}

	var envelope struct {
		Code int            `json:"code"`
		Msg  string         `json:"msg"`
		Data apiRewardsData `json:"data"`
	}
	if err := json.Unmarshal(body, &envelope); err != nil {
		return nil, err
	}
	if envelope.Code != 200 && envelope.Code != 0 {
		return nil, fmt.Errorf("api code=%d msg=%s", envelope.Code, envelope.Msg)
	}
	return &envelope.Data, nil
}

func (im *Importer) ensureOSSURL(ctx context.Context, db *gorm.DB, srcURL, objectKey, contract, tokenID string) (string, error) {
	downloadURL, cacheSrc, err := im.resolveDownloadURL(ctx, srcURL, contract, tokenID)
	if err != nil {
		return "", err
	}
	hash := sha256Hex(cacheSrc)

	var cached model.OSSObject
	err = db.WithContext(ctx).Where("src_url_hash = ?", hash).First(&cached).Error
	if err == nil && cached.OSSURL != "" {
		return cached.OSSURL, nil
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", err
	}

	body, contentType, err := im.downloadImage(ctx, downloadURL)
	if err != nil {
		return "", err
	}
	if ext := extFromContentType(contentType); ext != "" && path.Ext(objectKey) == "" {
		objectKey = objectKey + ext
	}

	ossURL, err := im.oss.Upload(objectKey, body, contentType)
	if err != nil {
		return "", err
	}

	obj := model.OSSObject{
		SrcURLHash:  hash,
		SrcURL:      cacheSrc,
		OSSURL:      ossURL,
		ContentType: contentType,
		CreatedAt:   time.Now(),
	}
	_ = db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "src_url_hash"}},
		DoNothing: true,
	}).Create(&obj).Error

	return ossURL, nil
}

// resolveDownloadURL prefers OpenSea when API key is configured.
// If OpenSea has no image, falls back to HTTPS image_uri_src only (skips empty/IPFS).
func (im *Importer) resolveDownloadURL(ctx context.Context, srcURL, contract, tokenID string) (downloadURL, cacheSrc string, err error) {
	contract = strings.ToLower(strings.TrimSpace(contract))
	tokenID = strings.TrimSpace(tokenID)

	if im.openSeaConfigured() && contract != "" && tokenID != "" {
		img, osErr := im.fetchOpenSeaImageURL(ctx, contract, tokenID)
		if osErr == nil && strings.TrimSpace(img) != "" {
			chain := strings.TrimSpace(im.cfg.OpenSeaChain)
			if chain == "" {
				chain = "ethereum"
			}
			cacheSrc = fmt.Sprintf("opensea:%s:%s:%s", chain, contract, tokenID)
			return img, cacheSrc, nil
		}
		if osErr != nil {
			log.Printf("opensea miss %s/%s: %v", contract, tokenID, osErr)
		} else {
			log.Printf("opensea miss %s/%s: empty image_url", contract, tokenID)
		}
		if httpsSrc := httpImageSrc(srcURL); httpsSrc != "" {
			log.Printf("fallback https src for %s/%s", contract, tokenID)
			return httpsSrc, httpsSrc, nil
		}
		return "", "", ErrSkipNoImage
	}

	normalized := im.normalizeImageURL(srcURL)
	if normalized == "" {
		return "", "", ErrSkipNoImage
	}
	return normalized, normalized, nil
}

func httpImageSrc(src string) string {
	src = strings.TrimSpace(src)
	lower := strings.ToLower(src)
	if strings.HasPrefix(lower, "https://") || strings.HasPrefix(lower, "http://") {
		return src
	}
	return ""
}

func (im *Importer) normalizeImageURL(raw string) string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return ""
	}
	if cidPath := extractIPFSPath(raw); cidPath != "" {
		return im.primaryGateway() + cidPath
	}
	lower := strings.ToLower(raw)
	if strings.HasPrefix(lower, "http://") || strings.HasPrefix(lower, "https://") {
		return raw
	}
	return ""
}

func (im *Importer) primaryGateway() string {
	gw := strings.TrimSpace(im.cfg.IPFSGateway)
	if gw == "" {
		gw = "https://dweb.link/ipfs/"
	}
	return strings.TrimRight(gw, "/") + "/"
}

func (im *Importer) ipfsGateways() []string {
	seen := map[string]struct{}{}
	var out []string
	add := func(gw string) {
		gw = strings.TrimRight(strings.TrimSpace(gw), "/") + "/"
		if gw == "/" {
			return
		}
		if _, ok := seen[gw]; ok {
			return
		}
		seen[gw] = struct{}{}
		out = append(out, gw)
	}
	add(im.primaryGateway())
	add("https://dweb.link/ipfs/")
	add("https://gateway.pinata.cloud/ipfs/")
	add("https://ipfs.io/ipfs/")
	return out
}

// extractIPFSPath returns "<cid>[/path]" for ipfs://, bare CID, or */ipfs/<cid> URLs.
func extractIPFSPath(raw string) string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return ""
	}
	lower := strings.ToLower(raw)
	switch {
	case strings.HasPrefix(lower, "ipfs://"):
		rest := strings.TrimPrefix(raw, "ipfs://")
		rest = strings.TrimPrefix(rest, "ipfs/")
		return strings.TrimLeft(rest, "/")
	case strings.Contains(lower, "/ipfs/"):
		idx := strings.Index(lower, "/ipfs/")
		return strings.TrimLeft(raw[idx+len("/ipfs/"):], "/")
	case strings.HasPrefix(lower, "qm") || strings.HasPrefix(lower, "bafy") || strings.HasPrefix(lower, "bafk"):
		return strings.TrimLeft(raw, "/")
	default:
		if !strings.Contains(raw, "://") && len(raw) > 10 {
			return strings.TrimLeft(raw, "/")
		}
		return ""
	}
}

func (im *Importer) candidateImageURLs(normalized string) []string {
	if cidPath := extractIPFSPath(normalized); cidPath != "" {
		gateways := im.ipfsGateways()
		out := make([]string, 0, len(gateways))
		for _, gw := range gateways {
			out = append(out, gw+cidPath)
		}
		return out
	}
	return []string{normalized}
}

func (im *Importer) downloadImage(ctx context.Context, normalized string) ([]byte, string, error) {
	var lastErr error
	for _, u := range im.candidateImageURLs(normalized) {
		reqCtx, cancel := context.WithTimeout(ctx, 25*time.Second)
		body, ct, err := im.downloadBytes(reqCtx, u)
		cancel()
		if err == nil {
			return body, ct, nil
		}
		lastErr = err
		log.Printf("download retry next gateway: %s: %v", u, err)
	}
	if lastErr == nil {
		lastErr = fmt.Errorf("no download url")
	}
	return nil, "", lastErr
}

func (im *Importer) downloadBytes(ctx context.Context, rawURL string) ([]byte, string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return nil, "", err
	}
	resp, err := im.httpClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("download http %d", resp.StatusCode)
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, 12<<20))
	if err != nil {
		return nil, "", err
	}
	ct := resp.Header.Get("Content-Type")
	if i := strings.Index(ct, ";"); i >= 0 {
		ct = strings.TrimSpace(ct[:i])
	}
	return body, ct, nil
}
