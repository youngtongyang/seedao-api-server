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

// LookupByWallet godoc
// @Summary      根据钱包地址查询 SNS 名称
// @Tags         sns
// @Produce      json
// @Param        wallet  query  string  true  "钱包地址"
// @Success      200  {object}  response.Envelope{data=service.WalletSNSLookupResult}
// @Router       /v1/sns/by-wallet [get]
func (h *SNSHandler) LookupByWallet(c echo.Context) error {
	wallet := c.QueryParam("wallet")
	if wallet == "" {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}
	result, err := h.sns.LookupWallet(c.Request().Context(), wallet)
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}
