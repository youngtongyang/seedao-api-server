package service

import "errors"

var (
	ErrInvalidParam       = errors.New("invalid parameter")
	ErrInvalidSMSCode     = errors.New("invalid sms code")
	ErrSNSUnavailable     = errors.New("sns unavailable")
	ErrWalletVerifyFailed = errors.New("wallet verify failed")
	ErrWalletNoSNS        = errors.New("wallet no sns")
	ErrWalletNonceMissing = errors.New("wallet nonce missing")
	ErrWalletDomainMismatch = errors.New("wallet domain mismatch")
	ErrWalletMessageInvalid = errors.New("wallet message invalid")
	ErrWalletSignatureInvalid = errors.New("wallet signature invalid")
	ErrWalletAddressMismatch  = errors.New("wallet address mismatch")
	ErrWalletDomainRequired   = errors.New("wallet domain required")
	ErrPhoneAlreadyExists  = errors.New("phone already exists")
	ErrPhoneNotRegistered  = errors.New("phone not registered")
	ErrUserNotFound        = errors.New("user not found")
	ErrTooManyRequests = errors.New("too many requests")
)

// RateLimitError carries retry_after seconds for 429 responses.
type RateLimitError struct {
	RetryAfter int
}

func (e *RateLimitError) Error() string {
	return ErrTooManyRequests.Error()
}

func (e *RateLimitError) Is(target error) bool {
	return target == ErrTooManyRequests
}

// Client-facing detail strings for wallet/SIWE errors (stable for frontend).
const (
	ReasonInvalidRequestBody   = "invalid_request_body"
	ReasonWalletRequired       = "wallet_required"
	ReasonDomainRequired       = "domain_required"
	ReasonNonceExpiredOrMissing = "nonce_expired_or_missing"
	ReasonDomainMismatch       = "domain_mismatch"
	ReasonSIWEMessageInvalid   = "siwe_message_invalid"
	ReasonSignatureInvalid     = "signature_invalid"
	ReasonWalletAddressMismatch = "wallet_address_mismatch"
	ReasonWalletNotInRegistry  = "wallet_not_in_registry"
)
