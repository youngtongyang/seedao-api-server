package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/seedao/seedao-api-server/internal/middleware"
	"github.com/seedao/seedao-api-server/internal/pkg/response"
	"github.com/seedao/seedao-api-server/internal/service"
)

type BindingsHandler struct {
	bindings *service.OAuthBindingService
}

func NewBindingsHandler(bindings *service.OAuthBindingService) *BindingsHandler {
	return &BindingsHandler{bindings: bindings}
}

type bindFeishuTokenRequest struct {
	BindToken string `json:"bind_token"`
}

// List godoc
// @Summary      查询第三方账号绑定状态
// @Tags         user
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  response.Envelope
// @Failure      401  {object}  response.Envelope
// @Router       /v1/user/bindings [get]
func (h *BindingsHandler) List(c echo.Context) error {
	userID, ok := middleware.UserIDFromContext(c)
	if !ok {
		return response.Fail(c, http.StatusUnauthorized, 40100)
	}

	result, err := h.bindings.ListByUserID(c.Request().Context(), userID)
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

// BindFeishuToken godoc
// @Summary      使用土地神 bind_token 绑定飞书
// @Tags         user
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body  bindFeishuTokenRequest  true  "bind token"
// @Success      200   {object}  response.Envelope
// @Failure      400   {object}  response.Envelope
// @Failure      401   {object}  response.Envelope
// @Router       /v1/user/bindings/feishu/token [post]
func (h *BindingsHandler) BindFeishuToken(c echo.Context) error {
	userID, ok := middleware.UserIDFromContext(c)
	if !ok {
		return response.Fail(c, http.StatusUnauthorized, 40100)
	}

	var req bindFeishuTokenRequest
	if err := c.Bind(&req); err != nil {
		return mapServiceError(c, service.ErrInvalidParam)
	}

	result, err := h.bindings.BindFeishuWithToken(c.Request().Context(), userID, req.BindToken)
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

// CreateFeishuChallenge godoc
// @Summary      生成飞书绑定验证码（官网 → 飞书路径）
// @Tags         user
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  response.Envelope
// @Failure      400  {object}  response.Envelope
// @Failure      401  {object}  response.Envelope
// @Router       /v1/user/bindings/feishu/challenge [post]
func (h *BindingsHandler) CreateFeishuChallenge(c echo.Context) error {
	userID, ok := middleware.UserIDFromContext(c)
	if !ok {
		return response.Fail(c, http.StatusUnauthorized, 40100)
	}

	result, err := h.bindings.CreateFeishuChallenge(c.Request().Context(), userID)
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

// UnbindFeishu godoc
// @Summary      解绑飞书账号
// @Tags         user
// @Produce      json
// @Security     BearerAuth
// @Success      200  {object}  response.Envelope
// @Failure      400  {object}  response.Envelope
// @Failure      401  {object}  response.Envelope
// @Router       /v1/user/bindings/feishu [delete]
func (h *BindingsHandler) UnbindFeishu(c echo.Context) error {
	userID, ok := middleware.UserIDFromContext(c)
	if !ok {
		return response.Fail(c, http.StatusUnauthorized, 40100)
	}

	result, err := h.bindings.UnbindFeishu(c.Request().Context(), userID)
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}
