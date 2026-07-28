package handler

import (
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/seedao/seedao-api-server/internal/pkg/response"
	"github.com/seedao/seedao-api-server/internal/service"
)

type InternalMembersHandler struct {
	members *service.MemberListService
}

func NewInternalMembersHandler(members *service.MemberListService) *InternalMembersHandler {
	return &InternalMembersHandler{members: members}
}

// List godoc
// @Summary      土地神分页拉取注册会员列表
// @Tags         internal
// @Produce      json
// @Param        X-Service-Token  header  string  true  "service token"
// @Param        page       query  int     false  "页码，从 1 开始，默认 1"
// @Param        page_size  query  int     false  "每页条数，默认 20，最大 100"
// @Param        sort_by    query  string  false  "排序字段：user_id|registered_at|membership_expires_at，默认 user_id"
// @Param        order      query  string  false  "排序方向：asc|desc，默认 asc"
// @Success      200  {object}  response.Envelope
// @Failure      401  {object}  response.Envelope
// @Router       /internal/v1/members [get]
func (h *InternalMembersHandler) List(c echo.Context) error {
	page := queryPositiveInt(c, "page", 1)
	pageSize := queryPositiveInt(c, "page_size", 20)

	result, err := h.members.ListRegisteredMembers(c.Request().Context(), service.ListRegisteredMembersInput{
		Page:     page,
		PageSize: pageSize,
		SortBy:   strings.TrimSpace(c.QueryParam("sort_by")),
		Order:    strings.TrimSpace(c.QueryParam("order")),
	})
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

func queryPositiveInt(c echo.Context, key string, def int) int {
	v := c.QueryParam(key)
	if v == "" {
		return def
	}
	n, err := strconv.Atoi(v)
	if err != nil || n <= 0 {
		return def
	}
	return n
}
