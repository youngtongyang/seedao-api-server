package siweutil

import "errors"

var (
	ErrParseMessage    = errors.New("parse siwe message")
	ErrDomainMismatch  = errors.New("domain mismatch")
	ErrSignatureInvalid = errors.New("verify signature")
	ErrWalletMismatch  = errors.New("wallet address mismatch")
)
