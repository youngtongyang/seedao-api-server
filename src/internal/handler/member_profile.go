package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/seedao/seedao-api-server/internal/pkg/errcode"
	"github.com/seedao/seedao-api-server/internal/pkg/response"
	"github.com/seedao/seedao-api-server/internal/service"
)

type MemberProfileHandler struct {
	svc *service.MemberProfileService
}

func NewMemberProfileHandler(svc *service.MemberProfileService) *MemberProfileHandler {
	return &MemberProfileHandler{svc: svc}
}

// Get godoc
// @Summary      获取成员社区资料（SEE/WANG/SCR/角色/SEED/SBT）
// @Description  按 sns_name 聚合 1.0 镜像表数据；图片优先返回 OSS 地址
// @Tags         sns
// @Produce      json
// @Param        sns_name  query  string  true  "SNS 名称，如 baiyu.seedao"
// @Success      200  {object}  response.Envelope{data=service.MemberProfileResult}
// @Failure      400  {object}  response.Envelope
// @Router       /v1/sns/profile [get]
func (h *MemberProfileHandler) Get(c echo.Context) error {
	snsName := c.QueryParam("sns_name")
	if snsName == "" {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}

	result, err := h.svc.GetBySNS(c.Request().Context(), snsName)
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}
