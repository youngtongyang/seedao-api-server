package handler

import (
	"github.com/labstack/echo/v4"

	"github.com/seedao/seedao-api-server/internal/pkg/response"
	"github.com/seedao/seedao-api-server/internal/service"
)

type InternalFeishuHandler struct {
	bindings *service.OAuthBindingService
}

func NewInternalFeishuHandler(bindings *service.OAuthBindingService) *InternalFeishuHandler {
	return &InternalFeishuHandler{bindings: bindings}
}

type internalCreateFeishuTokenRequest struct {
	OpenID        string `json:"open_id"`
	UnionID       string `json:"union_id"`
	ProviderAppID string `json:"provider_app_id"`
	DisplayName   string `json:"display_name"`
}

type internalConfirmFeishuRequest struct {
	BindCode      string `json:"bind_code"`
	OpenID        string `json:"open_id"`
	UnionID       string `json:"union_id"`
	ProviderAppID string `json:"provider_app_id"`
	DisplayName   string `json:"display_name"`
}

type internalUnbindFeishuRequest struct {
	OpenID        string `json:"open_id"`
	ProviderAppID string `json:"provider_app_id"`
}

// CreateBindToken godoc
// @Summary      土地神创建飞书 bind_token
// @Tags         internal
// @Accept       json
// @Produce      json
// @Param        X-Service-Token  header  string  true  "service token"
// @Param        body  body  internalCreateFeishuTokenRequest  true  "open id"
// @Success      200   {object}  response.Envelope
// @Failure      400   {object}  response.Envelope
// @Failure      401   {object}  response.Envelope
// @Router       /internal/v1/feishu/bindings/token [post]
func (h *InternalFeishuHandler) CreateBindToken(c echo.Context) error {
	var req internalCreateFeishuTokenRequest
	if err := c.Bind(&req); err != nil {
		return mapServiceError(c, service.ErrInvalidParam)
	}

	result, err := h.bindings.CreateFeishuBindToken(c.Request().Context(), service.FeishuCreateTokenInput{
		OpenID:        req.OpenID,
		UnionID:       req.UnionID,
		ProviderAppID: req.ProviderAppID,
		DisplayName:   req.DisplayName,
	})
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

// ConfirmBinding godoc
// @Summary      土地神确认飞书 bind_code 绑定
// @Tags         internal
// @Accept       json
// @Produce      json
// @Param        X-Service-Token  header  string  true  "service token"
// @Param        body  body  internalConfirmFeishuRequest  true  "bind code"
// @Success      200   {object}  response.Envelope
// @Failure      400   {object}  response.Envelope
// @Failure      401   {object}  response.Envelope
// @Router       /internal/v1/feishu/bindings/confirm [post]
func (h *InternalFeishuHandler) ConfirmBinding(c echo.Context) error {
	var req internalConfirmFeishuRequest
	if err := c.Bind(&req); err != nil {
		return mapServiceError(c, service.ErrInvalidParam)
	}

	result, err := h.bindings.ConfirmFeishuBinding(c.Request().Context(), service.FeishuConfirmInput{
		BindCode:      req.BindCode,
		OpenID:        req.OpenID,
		UnionID:       req.UnionID,
		ProviderAppID: req.ProviderAppID,
		DisplayName:   req.DisplayName,
	})
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

// UnbindBinding godoc
// @Summary      土地神按 open_id 解绑飞书
// @Tags         internal
// @Accept       json
// @Produce      json
// @Param        X-Service-Token  header  string  true  "service token"
// @Param        body  body  internalUnbindFeishuRequest  true  "open id"
// @Success      200   {object}  response.Envelope
// @Failure      400   {object}  response.Envelope
// @Failure      401   {object}  response.Envelope
// @Router       /internal/v1/feishu/bindings/unbind [post]
func (h *InternalFeishuHandler) UnbindBinding(c echo.Context) error {
	var req internalUnbindFeishuRequest
	if err := c.Bind(&req); err != nil {
		return mapServiceError(c, service.ErrInvalidParam)
	}

	result, err := h.bindings.UnbindFeishuByOpenID(c.Request().Context(), service.FeishuUnbindByOpenIDInput{
		OpenID:        req.OpenID,
		ProviderAppID: req.ProviderAppID,
	})
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}
