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
