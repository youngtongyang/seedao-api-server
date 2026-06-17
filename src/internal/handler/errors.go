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
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	case errors.Is(err, service.ErrInvalidSMSCode):
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidSMSCode)
	case errors.Is(err, service.ErrSNSUnavailable):
		return response.Fail(c, http.StatusBadRequest, errcode.SNSUnavailable)
	case errors.Is(err, service.ErrWalletVerifyFailed):
		return response.Fail(c, http.StatusBadRequest, errcode.WalletVerifyFailed)
	case errors.Is(err, service.ErrWalletNoSNS):
		return response.Fail(c, http.StatusBadRequest, errcode.WalletNoSNS)
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
