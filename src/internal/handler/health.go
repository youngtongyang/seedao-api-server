package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/seedao/seedao-api-server/internal/pkg/response"
)

func Health(c echo.Context) error {
	return response.OK(c, map[string]string{"status": "ok"})
}
