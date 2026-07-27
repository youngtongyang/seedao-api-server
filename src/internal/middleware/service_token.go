package middleware

import (
	"crypto/subtle"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/seedao/seedao-api-server/internal/pkg/response"
)

func ServiceToken(expected string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if expected == "" {
				return response.Fail(c, http.StatusUnauthorized, 40100)
			}
			token := c.Request().Header.Get("X-Service-Token")
			if subtle.ConstantTimeCompare([]byte(token), []byte(expected)) != 1 {
				return response.Fail(c, http.StatusUnauthorized, 40100)
			}
			return next(c)
		}
	}
}
