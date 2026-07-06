package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/seedao/seedao-api-server/internal/pkg/response"
	"github.com/seedao/seedao-api-server/internal/service"
)

type StatsHandler struct {
	stats *service.StatsService
}

func NewStatsHandler(stats *service.StatsService) *StatsHandler {
	return &StatsHandler{stats: stats}
}

// UserCount godoc
// @Summary      获取用户总数
// @Tags         stats
// @Produce      json
// @Success      200  {object}  response.Envelope
// @Router       /v1/stats/user-count [get]
func (h *StatsHandler) UserCount(c echo.Context) error {
	result, err := h.stats.GetUserCount()
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}
