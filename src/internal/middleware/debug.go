package middleware

import (
	"github.com/labstack/echo/v4"
)

const debugKey = "debug"

func DebugMode(enabled bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(debugKey, enabled)
			return next(c)
		}
	}
}

func IsDebug(c echo.Context) bool {
	v, _ := c.Get(debugKey).(bool)
	return v
}
