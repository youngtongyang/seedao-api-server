package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/seedao/seedao-api-server/internal/pkg/jwtutil"
	"github.com/seedao/seedao-api-server/internal/pkg/response"
)

const userIDKey = "user_id"

func JWT(jwtMgr *jwtutil.Manager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get("Authorization")
			if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
				return response.Fail(c, http.StatusUnauthorized, 40100)
			}

			token := strings.TrimPrefix(auth, "Bearer ")
			userID, err := jwtMgr.Parse(token)
			if err != nil {
				return response.Fail(c, http.StatusUnauthorized, 40100)
			}

			c.Set(userIDKey, userID)
			return next(c)
		}
	}
}

func UserIDFromContext(c echo.Context) (int64, bool) {
	v, ok := c.Get(userIDKey).(int64)
	return v, ok
}
