package siweutil

import (
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	siwe "github.com/spruceid/siwe-go"
)

type VerifyInput struct {
	Wallet    string
	Domain    string
	Message   string
	Signature string
	Nonce     string
}

type VerifyResult struct {
	Address string
}

func Verify(input VerifyInput) (*VerifyResult, error) {
	msg, err := siwe.ParseMessage(input.Message)
	if err != nil {
		return nil, fmt.Errorf("parse siwe message: %w", err)
	}

	if !strings.EqualFold(msg.GetDomain(), input.Domain) {
		return nil, fmt.Errorf("domain mismatch")
	}

	domain := input.Domain
	var noncePtr *string
	if input.Nonce != "" {
		noncePtr = &input.Nonce
	}

	pubKey, err := msg.Verify(input.Signature, &domain, noncePtr, nil)
	if err != nil {
		return nil, fmt.Errorf("verify signature: %w", err)
	}

	addr := crypto.PubkeyToAddress(*pubKey)
	if !strings.EqualFold(addr.Hex(), input.Wallet) {
		return nil, fmt.Errorf("wallet address mismatch")
	}

	return &VerifyResult{Address: strings.ToLower(addr.Hex())}, nil
}

func BuildMessage(domain, address, nonce string, chainID int64, issuedAt time.Time) (string, error) {
	uri := "https://" + domain
	msg, err := siwe.InitMessage(domain, address, uri, nonce, map[string]interface{}{
		"chainId":  int(chainID),
		"issuedAt": issuedAt.UTC().Format(time.RFC3339),
	})
	if err != nil {
		return "", err
	}
	return msg.String(), nil
}
