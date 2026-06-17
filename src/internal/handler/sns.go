package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/seedao/seedao-api-server/internal/pkg/errcode"
	"github.com/seedao/seedao-api-server/internal/pkg/response"
	"github.com/seedao/seedao-api-server/internal/service"
)

type SNSHandler struct {
	sns *service.SNSService
}

func NewSNSHandler(sns *service.SNSService) *SNSHandler {
	return &SNSHandler{sns: sns}
}

func (h *SNSHandler) Check(c echo.Context) error {
	snsName := c.QueryParam("sns_name")
	if snsName == "" {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}

	result, err := h.sns.Check(c.Request().Context(), snsName)
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}
