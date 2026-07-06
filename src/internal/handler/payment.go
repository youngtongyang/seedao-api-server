package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/seedao/seedao-api-server/internal/middleware"
	"github.com/seedao/seedao-api-server/internal/pkg/errcode"
	"github.com/seedao/seedao-api-server/internal/pkg/response"
	"github.com/seedao/seedao-api-server/internal/service"
)

type PaymentHandler struct {
	payment *service.PaymentService
	devMode bool
}

func NewPaymentHandler(payment *service.PaymentService, devMode bool) *PaymentHandler {
	return &PaymentHandler{payment: payment, devMode: devMode}
}

type createPaymentOrderRequest struct {
	ProductType string `json:"product_type"`
	Days        int    `json:"days"`
	PayChannel  string `json:"pay_channel"`
	CouponCode  string `json:"coupon_code"`
}

type previewCouponRequest struct {
	ProductType string `json:"product_type"`
	Days        int    `json:"days"`
	CouponCode  string `json:"coupon_code"`
}

// PreviewCoupon godoc
// @Summary      会员续费优惠券试算
// @Tags         payments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body  previewCouponRequest  true  "券码试算"
// @Success      200  {object}  response.Envelope
// @Router       /v1/payments/coupons/preview [post]
func (h *PaymentHandler) PreviewCoupon(c echo.Context) error {
	var req previewCouponRequest
	if err := c.Bind(&req); err != nil || req.CouponCode == "" {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}

	userID, userKnown := middleware.UserIDFromContext(c)
	if !userKnown {
		userID = 0
	}

	result, err := h.payment.PreviewCoupon(service.CouponQuoteInput{
		UserID:      userID,
		UserKnown:   userKnown,
		ProductType: req.ProductType,
		Days:        req.Days,
		CouponCode:  req.CouponCode,
	})
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

// CreateOrder godoc
// @Summary      创建支付订单（会员续费）
// @Tags         payments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body  createPaymentOrderRequest  true  "下单"
// @Success      200  {object}  response.Envelope
// @Router       /v1/payments/orders [post]
func (h *PaymentHandler) CreateOrder(c echo.Context) error {
	userID, ok := middleware.UserIDFromContext(c)
	if !ok {
		return response.Fail(c, http.StatusUnauthorized, 40100)
	}

	var req createPaymentOrderRequest
	if err := c.Bind(&req); err != nil {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}

	result, err := h.payment.CreateOrder(c.Request().Context(), userID, service.CreatePaymentOrderInput{
		ProductType: req.ProductType,
		Days:        req.Days,
		PayChannel:  req.PayChannel,
		ClientIP:    c.RealIP(),
		CouponCode:  req.CouponCode,
	})
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

// GetOrder godoc
// @Summary      查询支付订单
// @Tags         payments
// @Produce      json
// @Security     BearerAuth
// @Param        order_no  path  string  true  "订单号"
// @Success      200  {object}  response.Envelope
// @Router       /v1/payments/orders/{order_no} [get]
func (h *PaymentHandler) GetOrder(c echo.Context) error {
	userID, ok := middleware.UserIDFromContext(c)
	if !ok {
		return response.Fail(c, http.StatusUnauthorized, 40100)
	}

	orderNo := c.Param("order_no")
	result, err := h.payment.GetOrder(userID, orderNo)
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

// ListOrders godoc
// @Summary      支付订单列表
// @Tags         payments
// @Produce      json
// @Security     BearerAuth
// @Param        limit  query  int  false  "条数，默认 20"
// @Success      200  {object}  response.Envelope
// @Router       /v1/payments/orders [get]
func (h *PaymentHandler) ListOrders(c echo.Context) error {
	userID, ok := middleware.UserIDFromContext(c)
	if !ok {
		return response.Fail(c, http.StatusUnauthorized, 40100)
	}

	limit := 20
	if err := echo.QueryParamsBinder(c).Int("limit", &limit).BindError(); err != nil {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}

	result, err := h.payment.ListOrders(userID, limit)
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

// CloseOrder godoc
// @Summary      关闭未支付订单
// @Tags         payments
// @Produce      json
// @Security     BearerAuth
// @Param        order_no  path  string  true  "订单号"
// @Success      200  {object}  response.Envelope
// @Router       /v1/payments/orders/{order_no}/close [post]
func (h *PaymentHandler) CloseOrder(c echo.Context) error {
	userID, ok := middleware.UserIDFromContext(c)
	if !ok {
		return response.Fail(c, http.StatusUnauthorized, 40100)
	}

	orderNo := c.Param("order_no")
	result, err := h.payment.CloseOrder(c.Request().Context(), userID, orderNo)
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

// WechatNotify godoc
// @Summary      微信支付回调（微信服务器调用）
// @Tags         payments
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /v1/payments/wechat/notify [post]
func (h *PaymentHandler) WechatNotify(c echo.Context) error {
	if err := h.payment.HandleWechatNotify(c.Request().Context(), c.Request()); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, map[string]string{
			"code":    "FAIL",
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"code":    "SUCCESS",
		"message": "成功",
	})
}

type simulatePayRequest struct {
	OrderNo string `json:"order_no"`
}

// SimulatePayDev godoc
// @Summary      模拟支付（仅 dev_mode）
// @Tags         payments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body  simulatePayRequest  true  "订单号"
// @Success      200  {object}  response.Envelope
// @Router       /v1/payments/dev/simulate-pay [post]
func (h *PaymentHandler) SimulatePayDev(c echo.Context) error {
	if !h.devMode {
		return response.Fail(c, http.StatusNotFound, errcode.InvalidParam)
	}

	userID, ok := middleware.UserIDFromContext(c)
	if !ok {
		return response.Fail(c, http.StatusUnauthorized, 40100)
	}

	var req simulatePayRequest
	if err := c.Bind(&req); err != nil || req.OrderNo == "" {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}

	result, err := h.payment.SimulatePayDev(c.Request().Context(), userID, req.OrderNo)
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}
