package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/seedao/seedao-api-server/internal/pkg/response"
)

// Health godoc
// @Summary      健康检查
// @Tags         health
// @Produce      json
// @Success      200  {object}  response.Envelope
// @Router       /health [get]
func Health(c echo.Context) error {
	return response.OK(c, map[string]string{"status": "ok"})
}
