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

// Check godoc
// @Summary      检查 SNS 名称是否可用
// @Tags         sns
// @Produce      json
// @Param        sns_name  query  string  true  "SNS 名称"
// @Success      200  {object}  response.Envelope
// @Router       /v1/sns/check [get]
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
