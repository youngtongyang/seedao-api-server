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

// SendSMS godoc
// @Summary      发送短信验证码
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body  sendSMSRequest  true  "手机号"
// @Success      200  {object}  response.Envelope
// @Failure      400  {object}  response.Envelope
// @Failure      429  {object}  response.Envelope
// @Router       /v1/auth/sms/send [post]
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

// CheckPhone godoc
// @Summary      检查手机号是否可注册
// @Tags         auth
// @Produce      json
// @Param        phone  query  string  true  "E.164 手机号"
// @Success      200  {object}  response.Envelope
// @Router       /v1/auth/phone/check [get]
func (h *AuthHandler) CheckPhone(c echo.Context) error {
	phoneNum := c.QueryParam("phone")
	if phoneNum == "" {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}

	result, err := h.auth.CheckPhone(phoneNum)
	if err != nil {
		return mapServiceError(c, err)
	}

	return response.WithCode(c, http.StatusOK, errcode.PhoneNotRegistered, result)
}

type loginRequest struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

// Login godoc
// @Summary      手机号登录
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body  loginRequest  true  "登录"
// @Success      200  {object}  response.Envelope
// @Failure      400  {object}  response.Envelope
// @Router       /v1/auth/login [post]
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

// Register godoc
// @Summary      注册（新 SNS / 认领存量 SNS）
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body  registerRequest  true  "注册"
// @Success      200  {object}  response.Envelope
// @Failure      400  {object}  response.Envelope
// @Router       /v1/auth/register [post]
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
