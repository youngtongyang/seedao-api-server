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

func (h *PaymentHandler) PreviewCoupon(c echo.Context) error {
	userID, ok := middleware.UserIDFromContext(c)
	if !ok {
		return response.Fail(c, http.StatusUnauthorized, 40100)
	}

	var req previewCouponRequest
	if err := c.Bind(&req); err != nil || req.CouponCode == "" {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}

	result, err := h.payment.PreviewCoupon(userID, service.CouponQuoteInput{
		ProductType: req.ProductType,
		Days:        req.Days,
		CouponCode:  req.CouponCode,
	})
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

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
