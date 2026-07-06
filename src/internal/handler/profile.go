package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/seedao/seedao-api-server/internal/middleware"
	"github.com/seedao/seedao-api-server/internal/pkg/response"
	"github.com/seedao/seedao-api-server/internal/service"
)

type ProfileHandler struct {
	profile *service.ProfileService
}

func NewProfileHandler(profile *service.ProfileService) *ProfileHandler {
	return &ProfileHandler{profile: profile}
}

// Get godoc
// @Summary      获取当前用户资料
// @Tags         user
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  response.Envelope
// @Failure      401  {object}  response.Envelope
// @Router       /v1/user/profile [get]
func (h *ProfileHandler) Get(c echo.Context) error {
	userID, ok := middleware.UserIDFromContext(c)
	if !ok {
		return response.Fail(c, http.StatusUnauthorized, 40100)
	}

	result, err := h.profile.GetByUserID(userID)
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}
