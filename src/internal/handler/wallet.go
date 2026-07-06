package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/seedao/seedao-api-server/internal/pkg/errcode"
	"github.com/seedao/seedao-api-server/internal/pkg/response"
	"github.com/seedao/seedao-api-server/internal/service"
)

type WalletHandler struct {
	wallet *service.WalletService
	sns    *service.SNSService
}

func NewWalletHandler(wallet *service.WalletService, sns *service.SNSService) *WalletHandler {
	return &WalletHandler{wallet: wallet, sns: sns}
}

type walletNonceRequest struct {
	Wallet string `json:"wallet"`
	Domain string `json:"domain"`
}

// Nonce godoc
// @Summary      获取钱包 SIWE nonce
// @Tags         wallet
// @Accept       json
// @Produce      json
// @Param        body  body  walletNonceRequest  true  "钱包与域名"
// @Success      200  {object}  response.Envelope
// @Router       /v1/wallet/nonce [post]
func (h *WalletHandler) Nonce(c echo.Context) error {
	var req walletNonceRequest
	if err := c.Bind(&req); err != nil {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}

	result, err := h.wallet.CreateNonce(c.Request().Context(), req.Wallet, req.Domain)
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

type verifySNSRequest struct {
	Wallet     string `json:"wallet"`
	WalletType string `json:"wallet_type"`
	Domain     string `json:"domain"`
	Message    string `json:"message"`
	Signature  string `json:"signature"`
}

// VerifySNS godoc
// @Summary      验证钱包签名并查询可认领 SNS
// @Tags         wallet
// @Accept       json
// @Produce      json
// @Param        body  body  verifySNSRequest  true  "SIWE 验签"
// @Success      200  {object}  response.Envelope
// @Router       /v1/wallet/verify-sns [post]
func (h *WalletHandler) VerifySNS(c echo.Context) error {
	var req verifySNSRequest
	if err := c.Bind(&req); err != nil {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}

	result, err := h.wallet.VerifySNS(c.Request().Context(), h.sns, service.VerifySNSInput{
		Wallet:     req.Wallet,
		WalletType: req.WalletType,
		Domain:     req.Domain,
		Message:    req.Message,
		Signature:  req.Signature,
	})
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}
