package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/seedao/seedao-api-server/internal/pkg/errcode"
	"github.com/seedao/seedao-api-server/internal/pkg/response"
	"github.com/seedao/seedao-api-server/internal/service"
)

type AuthHandler struct {
	auth *service.AuthService
	sms  *service.SMSService
}

func NewAuthHandler(auth *service.AuthService, sms *service.SMSService) *AuthHandler {
	return &AuthHandler{auth: auth, sms: sms}
}

type sendSMSRequest struct {
	Phone string `json:"phone"`
}

func (h *AuthHandler) SendSMS(c echo.Context) error {
	var req sendSMSRequest
	if err := c.Bind(&req); err != nil {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}

	result, err := h.sms.Send(c.Request().Context(), req.Phone, c.RealIP())
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

type loginRequest struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req loginRequest
	if err := c.Bind(&req); err != nil {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}

	result, err := h.auth.Login(c.Request().Context(), req.Phone, req.Code)
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

type registerRequest struct {
	Phone        string `json:"phone"`
	Code         string `json:"code"`
	RegisterType string `json:"register_type"`
	SNSName      string `json:"sns_name"`
	Wallet       string `json:"wallet"`
	WalletType   string `json:"wallet_type"`
	Domain       string `json:"domain"`
	Message      string `json:"message"`
	Signature    string `json:"signature"`
}

func (h *AuthHandler) Register(c echo.Context) error {
	var req registerRequest
	if err := c.Bind(&req); err != nil {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}

	var result *service.AuthResult
	var err error

	switch req.RegisterType {
	case "new_sns":
		result, err = h.auth.RegisterNewSNS(c.Request().Context(), service.RegisterNewSNSInput{
			Phone:   req.Phone,
			Code:    req.Code,
			SNSName: req.SNSName,
		})
	case "existing_sns":
		result, err = h.auth.RegisterExistingSNS(c.Request().Context(), service.RegisterExistingSNSInput{
			Phone:      req.Phone,
			Code:       req.Code,
			Wallet:     req.Wallet,
			WalletType: req.WalletType,
			Domain:     req.Domain,
			Message:    req.Message,
			Signature:  req.Signature,
		})
	default:
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}

	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}
