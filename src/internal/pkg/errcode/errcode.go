package errcode

const (
	OK                  = 0
	InvalidParam        = 40001
	InvalidSMSCode      = 40002
	SNSUnavailable      = 40003
	WalletVerifyFailed  = 40004
	WalletNoSNS         = 40005
	PhoneAlreadyExists = 40006
	TooManyRequests     = 42901
	PaymentOrderNotFound = 40007
	PaymentOrderInvalid  = 40008
	PhoneNotRegistered   = 40009
	InternalError       = 50000
)

var messages = map[int]string{
	OK:                 "ok",
	InvalidParam:       "invalid parameter",
	InvalidSMSCode:     "invalid or expired verification code",
	SNSUnavailable:     "sns name unavailable",
	WalletVerifyFailed: "wallet signature verification failed",
	WalletNoSNS:        "wallet has no claimable sns",
	PhoneAlreadyExists: "phone already registered",
	TooManyRequests:    "too many requests",
	PaymentOrderNotFound: "payment order not found",
	PaymentOrderInvalid:  "payment order invalid",
	PhoneNotRegistered:   "phone not registered",
	InternalError:      "internal server error",
}

func Message(code int) string {
	if msg, ok := messages[code]; ok {
		return msg
	}
	return "unknown error"
}
