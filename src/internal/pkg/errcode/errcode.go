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
	CouponInvalid        = 40010
	CouponNotApplicable  = 40011
	CouponUsageExceeded  = 40012
	CouponAmountTooLow   = 40013
	CourseNotFound       = 40014
	CourseOffSale        = 40015
	CourseMemberLogin    = 40016
	CourseOrderTokenRequired = 40017
	CourseOrderTokenInvalid  = 40018
	CourseOrderClaimed       = 40019
	ProposalNotFound         = 40020
	ProposalForbidden        = 40021
	ProposalInvalidStatus    = 40022
	ProposalVoteClosed       = 40023
	ProposalAlreadyVoted     = 40024
	ProposalNotVoter         = 40025
	ProposalSNSRequired      = 40026
	ProposalVoteNotFound     = 40027
	FeishuNotBound           = 40040
	FeishuAlreadyBound       = 40041
	FeishuOpenIDTaken        = 40042
	FeishuBindCredentialInvalid = 40043
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
	CouponInvalid:        "coupon invalid",
	CouponNotApplicable:  "coupon not applicable",
	CouponUsageExceeded:  "coupon usage exceeded",
	CouponAmountTooLow:   "coupon amount too low",
	CourseNotFound:       "course not found",
	CourseOffSale:        "course off sale",
	CourseMemberLogin:    "member discount requires login",
	CourseOrderTokenRequired: "order token required",
	CourseOrderTokenInvalid:  "order token invalid",
	CourseOrderClaimed:       "order already claimed",
	ProposalNotFound:         "proposal not found",
	ProposalForbidden:        "proposal operation forbidden",
	ProposalInvalidStatus:    "proposal status invalid",
	ProposalVoteClosed:       "proposal voting closed",
	ProposalAlreadyVoted:     "already voted on this proposal",
	ProposalNotVoter:         "not eligible to vote",
	ProposalSNSRequired:      "sns required for governance action",
	ProposalVoteNotFound:     "vote not found for this proposal",
	FeishuNotBound:           "feishu not bound",
	FeishuAlreadyBound:       "feishu already bound",
	FeishuOpenIDTaken:        "feishu account already bound to another user",
	FeishuBindCredentialInvalid: "feishu bind token or code invalid or expired",
	InternalError:      "internal server error",
}

func Message(code int) string {
	if msg, ok := messages[code]; ok {
		return msg
	}
	return "unknown error"
}
