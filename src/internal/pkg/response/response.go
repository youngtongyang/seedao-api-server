package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/seedao/seedao-api-server/internal/pkg/errcode"
)

type Body struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Detail  string      `json:"detail,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func OK(c echo.Context, data interface{}) error {
	return WithCode(c, http.StatusOK, errcode.OK, data)
}

func WithCode(c echo.Context, httpStatus, code int, data interface{}) error {
	return c.JSON(httpStatus, Body{
		Code:    code,
		Message: errcode.Message(code),
		Data:    data,
	})
}

func Fail(c echo.Context, httpStatus, code int) error {
	return c.JSON(httpStatus, Body{
		Code:    code,
		Message: errcode.Message(code),
	})
}

func FailWithMessage(c echo.Context, httpStatus, code int, message string) error {
	return c.JSON(httpStatus, Body{
		Code:    code,
		Message: message,
	})
}

func FailWithDetail(c echo.Context, httpStatus, code int, detail string) error {
	body := Body{
		Code:    code,
		Message: errcode.Message(code),
	}
	if detail != "" {
		body.Detail = detail
	}
	return c.JSON(httpStatus, body)
}
