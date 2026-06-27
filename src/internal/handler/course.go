package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/seedao/seedao-api-server/internal/middleware"
	"github.com/seedao/seedao-api-server/internal/pkg/errcode"
	"github.com/seedao/seedao-api-server/internal/pkg/response"
	"github.com/seedao/seedao-api-server/internal/service"
)

type CourseHandler struct {
	course  *service.CourseService
	payment *service.PaymentService
	devMode bool
}

func NewCourseHandler(course *service.CourseService, payment *service.PaymentService, devMode bool) *CourseHandler {
	return &CourseHandler{course: course, payment: payment, devMode: devMode}
}

type courseCheckoutPreviewRequest struct {
	CourseID            string `json:"course_id"`
	SKUID               string `json:"sku_id"`
	ChannelCode         string `json:"channel_code"`
	ApplyMemberDiscount bool   `json:"apply_member_discount"`
}

type createCourseOrderRequest struct {
	CourseID            string `json:"course_id"`
	SKUID               string `json:"sku_id"`
	ChannelCode         string `json:"channel_code"`
	ApplyMemberDiscount bool   `json:"apply_member_discount"`
	PayChannel          string `json:"pay_channel"`
}

type courseOrderTokenRequest struct {
	OrderToken string `json:"order_token"`
}

type claimCourseOrderRequest struct {
	OrderNo    string `json:"order_no"`
	OrderToken string `json:"order_token"`
}

func (h *CourseHandler) PreviewCheckout(c echo.Context) error {
	var req courseCheckoutPreviewRequest
	if err := c.Bind(&req); err != nil || req.CourseID == "" || req.SKUID == "" {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}

	userID, userKnown := middleware.UserIDFromContext(c)
	if !userKnown {
		userID = 0
	}

	result, err := h.course.PreviewCheckout(service.CourseCheckoutInput{
		CourseID:            req.CourseID,
		SKUID:               req.SKUID,
		ChannelCode:         req.ChannelCode,
		ApplyMemberDiscount: req.ApplyMemberDiscount,
		UserID:              userID,
		UserKnown:           userKnown,
	})
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

func (h *CourseHandler) CreateOrder(c echo.Context) error {
	var req createCourseOrderRequest
	if err := c.Bind(&req); err != nil || req.CourseID == "" || req.SKUID == "" {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}

	userID, userKnown := middleware.UserIDFromContext(c)
	if !userKnown {
		userID = 0
	}
	if req.ApplyMemberDiscount && !userKnown {
		return response.Fail(c, http.StatusUnauthorized, 40100)
	}

	result, err := h.course.CreateOrder(c.Request().Context(), service.CreateCourseOrderInput{
		CourseCheckoutInput: service.CourseCheckoutInput{
			CourseID:            req.CourseID,
			SKUID:               req.SKUID,
			ChannelCode:         req.ChannelCode,
			ApplyMemberDiscount: req.ApplyMemberDiscount,
			UserID:              userID,
			UserKnown:           userKnown,
		},
		PayChannel: req.PayChannel,
		ClientIP:   c.RealIP(),
	})
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

func (h *CourseHandler) GetOrder(c echo.Context) error {
	orderNo := c.Param("order_no")
	token := c.QueryParam("token")

	userID, userKnown := middleware.UserIDFromContext(c)
	if !userKnown {
		userID = 0
	}
	if !userKnown && token == "" {
		return response.Fail(c, http.StatusBadRequest, errcode.CourseOrderTokenRequired)
	}

	result, err := h.course.GetOrder(userID, userKnown, orderNo, token)
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

func (h *CourseHandler) CloseOrder(c echo.Context) error {
	orderNo := c.Param("order_no")

	userID, userKnown := middleware.UserIDFromContext(c)
	if !userKnown {
		userID = 0
	}

	token := c.QueryParam("token")
	if !userKnown {
		var req courseOrderTokenRequest
		if err := c.Bind(&req); err == nil && req.OrderToken != "" {
			token = req.OrderToken
		}
		if token == "" {
			return response.Fail(c, http.StatusBadRequest, errcode.CourseOrderTokenRequired)
		}
	}

	result, err := h.course.CloseOrder(c.Request().Context(), userID, userKnown, orderNo, token)
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

func (h *CourseHandler) ClaimOrder(c echo.Context) error {
	userID, ok := middleware.UserIDFromContext(c)
	if !ok {
		return response.Fail(c, http.StatusUnauthorized, 40100)
	}

	var req claimCourseOrderRequest
	if err := c.Bind(&req); err != nil || req.OrderNo == "" || req.OrderToken == "" {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}

	result, err := h.course.ClaimOrder(userID, req.OrderNo, req.OrderToken)
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

func (h *CourseHandler) SimulatePayDev(c echo.Context) error {
	if !h.devMode {
		return response.Fail(c, http.StatusNotFound, errcode.InvalidParam)
	}

	orderNo := c.Param("order_no")
	userID, userKnown := middleware.UserIDFromContext(c)
	if !userKnown {
		userID = 0
	}

	token := c.QueryParam("token")
	if !userKnown {
		var req courseOrderTokenRequest
		if err := c.Bind(&req); err == nil && req.OrderToken != "" {
			token = req.OrderToken
		}
		if token == "" {
			return response.Fail(c, http.StatusBadRequest, errcode.CourseOrderTokenRequired)
		}
	}

	result, err := h.course.SimulatePayDev(
		c.Request().Context(),
		userID,
		userKnown,
		orderNo,
		token,
		h.payment.CompleteOrderPayment,
	)
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}
