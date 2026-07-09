package service

import (
	"strings"

	"github.com/seedao/seedao-api-server/internal/model"
	"github.com/seedao/seedao-api-server/internal/pkg/phone"
)

const zeroWallet = "0x0000000000000000000000000000000000000000"

// NormalizeWallet lowercases wallet addresses; empty input maps to the zero address.
func NormalizeWallet(wallet string) string {
	w := phone.NormalizeWallet(wallet)
	if w == "" {
		return zeroWallet
	}
	return w
}

// IsZeroWallet reports whether the wallet is empty or the zero address.
func IsZeroWallet(wallet string) bool {
	return NormalizeWallet(wallet) == zeroWallet
}

// SNSNameNeedsResolve reports whether sns_name should be filled from wallet lookup.
func SNSNameNeedsResolve(sns *string) bool {
	return sns == nil || strings.TrimSpace(*sns) == ""
}

// ResolveSNSNameForDisplay returns sns when set, otherwise lookup by wallet, else shortened wallet.
func ResolveSNSNameForDisplay(current *string, wallet string, byWallet map[string]string) *string {
	if !SNSNameNeedsResolve(current) {
		return current
	}
	if byWallet != nil {
		if name := byWallet[NormalizeWallet(wallet)]; name != "" {
			return &name
		}
	}
	if IsZeroWallet(wallet) {
		return current
	}
	fallback := FormatWalletDisplay(wallet)
	return &fallback
}

// FormatWalletDisplay shortens a wallet for display (first 6 + "..." + last 4).
func FormatWalletDisplay(wallet string) string {
	w := NormalizeWallet(wallet)
	if len(w) <= 10 {
		return w
	}
	return w[:6] + "..." + w[len(w)-4:]
}

// IsWalletDisplayName reports whether the value is a wallet address used for display.
func IsWalletDisplayName(name string) bool {
	return strings.HasPrefix(strings.ToLower(strings.TrimSpace(name)), "0x")
}

// LookupSNSByWallet resolves one wallet to SNS name (sns_names first, then sns_chain_registry).
func (s *SNSService) LookupSNSByWallet(wallet string) (string, error) {
	names, err := s.LookupSNSByWallets([]string{wallet})
	if err != nil {
		return "", err
	}
	if name := names[NormalizeWallet(wallet)]; name != "" {
		return name, nil
	}
	if IsZeroWallet(wallet) {
		return "", nil
	}
	return FormatWalletDisplay(wallet), nil
}

// LookupSNSByWallets batch-resolves wallets. Map keys are normalized lowercase wallets.
func (s *SNSService) LookupSNSByWallets(wallets []string) (map[string]string, error) {
	walletSet := make(map[string]struct{})
	for _, wallet := range wallets {
		if IsZeroWallet(wallet) {
			continue
		}
		walletSet[NormalizeWallet(wallet)] = struct{}{}
	}
	if len(walletSet) == 0 {
		return map[string]string{}, nil
	}

	uniqueWallets := make([]string, 0, len(walletSet))
	for w := range walletSet {
		uniqueWallets = append(uniqueWallets, w)
	}

	out := make(map[string]string, len(uniqueWallets))

	var snsRows []model.SNSName
	if err := s.db.Where("wallet IS NOT NULL AND LOWER(wallet) IN ?", uniqueWallets).Find(&snsRows).Error; err != nil {
		return nil, err
	}
	for i := range snsRows {
		if snsRows[i].Wallet == nil {
			continue
		}
		w := NormalizeWallet(*snsRows[i].Wallet)
		if snsRows[i].Name != "" {
			out[w] = snsRows[i].Name
		}
	}

	missing := make([]string, 0)
	for _, w := range uniqueWallets {
		if out[w] == "" {
			missing = append(missing, w)
		}
	}
	if len(missing) == 0 {
		return out, nil
	}

	var regRows []model.SNSChainRegistry
	if err := s.db.Where("LOWER(wallet) IN ?", missing).Find(&regRows).Error; err != nil {
		return nil, err
	}
	for i := range regRows {
		w := NormalizeWallet(regRows[i].Wallet)
		if out[w] == "" && regRows[i].SNSName != "" {
			out[w] = regRows[i].SNSName
		}
	}
	return out, nil
}
