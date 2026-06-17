package service

import "errors"

var (
	ErrInvalidParam       = errors.New("invalid parameter")
	ErrInvalidSMSCode     = errors.New("invalid sms code")
	ErrSNSUnavailable     = errors.New("sns unavailable")
	ErrWalletVerifyFailed = errors.New("wallet verify failed")
	ErrWalletNoSNS        = errors.New("wallet no sns")
	ErrPhoneAlreadyExists = errors.New("phone already exists")
	ErrUserNotFound       = errors.New("user not found")
	ErrTooManyRequests    = errors.New("too many requests")
)
