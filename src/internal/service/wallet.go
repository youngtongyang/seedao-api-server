package service

import (
	"context"
	"errors"
	"strings"
	"time"

	siwe "github.com/spruceid/siwe-go"
	"github.com/seedao/seedao-api-server/internal/config"
	"github.com/seedao/seedao-api-server/internal/pkg/phone"
	"github.com/seedao/seedao-api-server/internal/pkg/redisstore"
	"github.com/seedao/seedao-api-server/internal/pkg/siweutil"
)

const walletNonceTTL = 45 * time.Second

type WalletService struct {
	redis   *redisstore.Store
	chainID int64
}

func NewWalletService(redis *redisstore.Store, cfg config.SIWEConfig) *WalletService {
	return &WalletService{redis: redis, chainID: cfg.ChainID}
}

type WalletNonceResult struct {
	Nonce     string `json:"nonce"`
	Message   string `json:"message"`
	ExpiresAt int64  `json:"expires_at"`
}

func (s *WalletService) CreateNonce(ctx context.Context, walletAddr, domain string) (*WalletNonceResult, error) {
	walletAddr = phone.NormalizeWallet(walletAddr)
	if walletAddr == "" {
		return nil, ErrInvalidParam
	}
	domain = strings.TrimSpace(domain)
	if domain == "" {
		return nil, ErrWalletDomainRequired
	}

	nonce := siwe.GenerateNonce()
	expiresAt := time.Now().Add(walletNonceTTL)

	message, err := siweutil.BuildMessage(domain, walletAddr, nonce, s.chainID, time.Now())
	if err != nil {
		return nil, err
	}

	if err := s.redis.SetWalletNonce(ctx, walletAddr, nonce, walletNonceTTL); err != nil {
		return nil, err
	}

	return &WalletNonceResult{
		Nonce:     nonce,
		Message:   message,
		ExpiresAt: expiresAt.Unix(),
	}, nil
}

type VerifySNSInput struct {
	Wallet     string
	WalletType string
	Domain     string
	Message    string
	Signature  string
}

type VerifySNSResult struct {
	Wallet    string `json:"wallet"`
	SNSName   string `json:"sns_name"`
	Claimable bool   `json:"claimable"`
}

func (s *WalletService) VerifyWalletSignature(ctx context.Context, input VerifySNSInput) error {
	walletAddr := phone.NormalizeWallet(input.Wallet)
	if walletAddr == "" {
		return ErrInvalidParam
	}

	storedNonce, err := s.redis.GetWalletNonce(ctx, walletAddr)
	if err != nil {
		return ErrWalletNonceMissing
	}

	_, err = siweutil.Verify(siweutil.VerifyInput{
		Wallet:    walletAddr,
		Domain:    strings.TrimSpace(input.Domain),
		Message:   input.Message,
		Signature: input.Signature,
		Nonce:     storedNonce,
	})
	if err != nil {
		return mapSIWEVerifyError(err)
	}

	_ = s.redis.DeleteWalletNonce(ctx, walletAddr)
	return nil
}

func (s *WalletService) VerifySNS(ctx context.Context, sns *SNSService, input VerifySNSInput) (*VerifySNSResult, error) {
	walletAddr := phone.NormalizeWallet(input.Wallet)

	if err := s.VerifyWalletSignature(ctx, input); err != nil {
		return nil, err
	}

	registry, err := sns.getRegistryByWallet(walletAddr)
	if err != nil {
		return nil, err
	}
	if registry == nil {
		return nil, ErrWalletNoSNS
	}

	return &VerifySNSResult{
		Wallet:    walletAddr,
		SNSName:   registry.SNSName,
		Claimable: true,
	}, nil
}

func mapSIWEVerifyError(err error) error {
	switch {
	case errors.Is(err, siweutil.ErrDomainMismatch):
		return ErrWalletDomainMismatch
	case errors.Is(err, siweutil.ErrParseMessage):
		return ErrWalletMessageInvalid
	case errors.Is(err, siweutil.ErrWalletMismatch):
		return ErrWalletAddressMismatch
	case errors.Is(err, siweutil.ErrSignatureInvalid):
		return ErrWalletSignatureInvalid
	default:
		return ErrWalletVerifyFailed
	}
}
