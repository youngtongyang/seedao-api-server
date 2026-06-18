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
