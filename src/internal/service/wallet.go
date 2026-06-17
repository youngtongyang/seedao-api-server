package service

import (
	"context"
	"time"

	"github.com/google/uuid"
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

	nonce := uuid.New().String()
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

func (s *WalletService) VerifySNS(ctx context.Context, sns *SNSService, input VerifySNSInput) (*VerifySNSResult, error) {
	walletAddr := phone.NormalizeWallet(input.Wallet)

	storedNonce, err := s.redis.GetWalletNonce(ctx, walletAddr)
	if err != nil {
		return nil, ErrWalletVerifyFailed
	}

	_, err = siweutil.Verify(siweutil.VerifyInput{
		Wallet:    walletAddr,
		Domain:    input.Domain,
		Message:   input.Message,
		Signature: input.Signature,
		Nonce:     storedNonce,
	})
	if err != nil {
		return nil, ErrWalletVerifyFailed
	}

	_ = s.redis.DeleteWalletNonce(ctx, walletAddr)

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
