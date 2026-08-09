package legacyimport

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func (im *Importer) openSeaConfigured() bool {
	return strings.TrimSpace(im.cfg.OpenSeaAPIKey) != ""
}

func (im *Importer) waitOpenSeaRateLimit() {
	im.openSeaMu.Lock()
	defer im.openSeaMu.Unlock()
	now := time.Now()
	if now.Before(im.openSeaNext) {
		time.Sleep(im.openSeaNext.Sub(now))
	}
	// Free-tier is typically ~60 reads/min; keep ~1 req/sec.
	im.openSeaNext = time.Now().Add(1100 * time.Millisecond)
}

func (im *Importer) fetchOpenSeaImageURL(ctx context.Context, contract, tokenID string) (string, error) {
	contract = strings.ToLower(strings.TrimSpace(contract))
	tokenID = strings.TrimSpace(tokenID)
	if contract == "" || tokenID == "" {
		return "", fmt.Errorf("missing contract or token_id")
	}
	if !im.openSeaConfigured() {
		return "", fmt.Errorf("opensea_api_key is not configured")
	}

	im.waitOpenSeaRateLimit()

	chain := strings.TrimSpace(im.cfg.OpenSeaChain)
	if chain == "" {
		chain = "ethereum"
	}
	base := strings.TrimRight(strings.TrimSpace(im.cfg.OpenSeaBaseURL), "/")
	if base == "" {
		base = "https://api.opensea.io"
	}

	apiURL := fmt.Sprintf("%s/api/v2/chain/%s/contract/%s/nfts/%s",
		base, url.PathEscape(chain), url.PathEscape(contract), url.PathEscape(tokenID))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("x-api-key", strings.TrimSpace(im.cfg.OpenSeaAPIKey))

	resp, err := im.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, 4<<20))
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("http %d: %s", resp.StatusCode, truncate(string(body), 200))
	}

	var payload struct {
		NFT struct {
			ImageURL        string `json:"image_url"`
			DisplayImageURL string `json:"display_image_url"`
			OriginalImageURL string `json:"original_image_url"`
		} `json:"nft"`
	}
	if err := json.Unmarshal(body, &payload); err != nil {
		return "", err
	}
	img := firstNonEmpty(payload.NFT.ImageURL, payload.NFT.DisplayImageURL, payload.NFT.OriginalImageURL)
	if img == "" {
		return "", fmt.Errorf("empty image_url in response")
	}
	return img, nil
}
