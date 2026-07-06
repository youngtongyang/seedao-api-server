package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/seedao/seedao-api-server/internal/middleware"
	"github.com/seedao/seedao-api-server/internal/pkg/errcode"
	"github.com/seedao/seedao-api-server/internal/pkg/response"
	"github.com/seedao/seedao-api-server/internal/service"
)

type MembershipHandler struct {
	membership *service.MembershipService
}

func NewMembershipHandler(membership *service.MembershipService) *MembershipHandler {
	return &MembershipHandler{membership: membership}
}

type renewMembershipRequest struct {
	Days int `json:"days"`
}

// Renew godoc
// @Summary      续费会员（运维/测试）
// @Tags         user
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body  renewMembershipRequest  true  "续费天数"
// @Success      200  {object}  response.Envelope
// @Router       /v1/user/membership/renew [post]
func (h *MembershipHandler) Renew(c echo.Context) error {
	userID, ok := middleware.UserIDFromContext(c)
	if !ok {
		return response.Fail(c, http.StatusUnauthorized, 40100)
	}

	var req renewMembershipRequest
	if err := c.Bind(&req); err != nil {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}

	result, err := h.membership.Renew(userID, service.RenewMembershipInput{
		Days: req.Days,
	})
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}
