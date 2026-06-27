package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/seedao/seedao-api-server/internal/middleware"
	"github.com/seedao/seedao-api-server/internal/pkg/errcode"
	"github.com/seedao/seedao-api-server/internal/pkg/response"
	"github.com/seedao/seedao-api-server/internal/service"
)

func mapServiceError(c echo.Context, err error) error {
	switch {
	case errors.Is(err, service.ErrInvalidParam):
		return failWithReason(c, http.StatusBadRequest, errcode.InvalidParam, service.ReasonInvalidRequestBody)
	case errors.Is(err, service.ErrWalletDomainRequired):
		return failWithReason(c, http.StatusBadRequest, errcode.InvalidParam, service.ReasonDomainRequired)
	case errors.Is(err, service.ErrInvalidSMSCode):
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidSMSCode)
	case errors.Is(err, service.ErrSNSUnavailable):
		return response.Fail(c, http.StatusBadRequest, errcode.SNSUnavailable)
	case errors.Is(err, service.ErrWalletNonceMissing):
		return failWithReason(c, http.StatusBadRequest, errcode.WalletVerifyFailed, service.ReasonNonceExpiredOrMissing)
	case errors.Is(err, service.ErrWalletDomainMismatch):
		return failWithReason(c, http.StatusBadRequest, errcode.WalletVerifyFailed, service.ReasonDomainMismatch)
	case errors.Is(err, service.ErrWalletMessageInvalid):
		return failWithReason(c, http.StatusBadRequest, errcode.WalletVerifyFailed, service.ReasonSIWEMessageInvalid)
	case errors.Is(err, service.ErrWalletSignatureInvalid):
		return failWithReason(c, http.StatusBadRequest, errcode.WalletVerifyFailed, service.ReasonSignatureInvalid)
	case errors.Is(err, service.ErrWalletAddressMismatch):
		return failWithReason(c, http.StatusBadRequest, errcode.WalletVerifyFailed, service.ReasonWalletAddressMismatch)
	case errors.Is(err, service.ErrWalletVerifyFailed):
		return failWithReason(c, http.StatusBadRequest, errcode.WalletVerifyFailed, service.ReasonSignatureInvalid)
	case errors.Is(err, service.ErrWalletNoSNS):
		return failWithReason(c, http.StatusBadRequest, errcode.WalletNoSNS, service.ReasonWalletNotInRegistry)
	case errors.Is(err, service.ErrPhoneAlreadyExists):
		return response.Fail(c, http.StatusBadRequest, errcode.PhoneAlreadyExists)
	case errors.Is(err, service.ErrTooManyRequests):
		return response.Fail(c, http.StatusTooManyRequests, errcode.TooManyRequests)
	case errors.Is(err, service.ErrUserNotFound):
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidSMSCode)
	case errors.Is(err, service.ErrPaymentOrderNotFound):
		return response.Fail(c, http.StatusBadRequest, errcode.PaymentOrderNotFound)
	case errors.Is(err, service.ErrPaymentOrderForbidden):
		return response.Fail(c, http.StatusForbidden, errcode.PaymentOrderNotFound)
	case errors.Is(err, service.ErrPaymentOrderNotPending):
		return response.Fail(c, http.StatusBadRequest, errcode.PaymentOrderInvalid)
	case errors.Is(err, service.ErrPaymentOrderExpired):
		return response.Fail(c, http.StatusBadRequest, errcode.PaymentOrderInvalid)
	case errors.Is(err, service.ErrPaymentProductInvalid):
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	case errors.Is(err, service.ErrPaymentChannelInvalid):
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	case errors.Is(err, service.ErrCouponInvalid):
		return response.Fail(c, http.StatusBadRequest, errcode.CouponInvalid)
	case errors.Is(err, service.ErrCouponNotApplicable):
		return response.Fail(c, http.StatusBadRequest, errcode.CouponNotApplicable)
	case errors.Is(err, service.ErrCouponUsageExceeded):
		return response.Fail(c, http.StatusBadRequest, errcode.CouponUsageExceeded)
	case errors.Is(err, service.ErrCouponAmountTooLow):
		return response.Fail(c, http.StatusBadRequest, errcode.CouponAmountTooLow)
	case errors.Is(err, service.ErrCourseNotFound), errors.Is(err, service.ErrCourseSKUNotFound):
		return response.Fail(c, http.StatusBadRequest, errcode.CourseNotFound)
	case errors.Is(err, service.ErrCourseOffSale):
		return response.Fail(c, http.StatusBadRequest, errcode.CourseOffSale)
	case errors.Is(err, service.ErrCourseMemberDiscountRequiresLogin):
		return response.Fail(c, http.StatusUnauthorized, errcode.CourseMemberLogin)
	case errors.Is(err, service.ErrCourseOrderTokenRequired):
		return response.Fail(c, http.StatusBadRequest, errcode.CourseOrderTokenRequired)
	case errors.Is(err, service.ErrCourseOrderTokenInvalid):
		return response.Fail(c, http.StatusForbidden, errcode.CourseOrderTokenInvalid)
	case errors.Is(err, service.ErrCourseOrderAlreadyClaimed):
		return response.Fail(c, http.StatusBadRequest, errcode.CourseOrderClaimed)
	default:
		c.Logger().Error(err)
		if middleware.IsDebug(c) {
			return response.FailWithDetail(c, http.StatusInternalServerError, errcode.InternalError, err.Error())
		}
		return response.Fail(c, http.StatusInternalServerError, errcode.InternalError)
	}
}

func failWithReason(c echo.Context, httpStatus, code int, reason string) error {
	return response.FailWithDetail(c, httpStatus, code, reason)
}
