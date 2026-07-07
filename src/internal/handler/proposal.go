package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/seedao/seedao-api-server/internal/middleware"
	"github.com/seedao/seedao-api-server/internal/pkg/errcode"
	"github.com/seedao/seedao-api-server/internal/pkg/response"
	"github.com/seedao/seedao-api-server/internal/service"
)

type ProposalHandler struct {
	proposals *service.ProposalService
	profile   *service.ProfileService
}

func NewProposalHandler(proposals *service.ProposalService, profile *service.ProfileService) *ProposalHandler {
	return &ProposalHandler{proposals: proposals, profile: profile}
}

// List godoc
// @Summary      分页获取提案列表
// @Tags         proposals
// @Produce      json
// @Param        limit     query  int     false  "每页条数，默认 20，最大 50"
// @Param        before_id query  int     false  "游标：传入上一页最后一条的 id，返回更早创建的记录"
// @Param        type      query  string  false  "提案类型"
// @Param        status    query  string  false  "提案状态"
// @Success      200  {object}  response.Envelope{data=service.ProposalListResult}
// @Router       /v1/proposals [get]
func (h *ProposalHandler) List(c echo.Context) error {
	limit := 20
	beforeID := int64(0)
	if v := c.QueryParam("limit"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			limit = n
		}
	}
	if v := c.QueryParam("before_id"); v != "" {
		if n, err := strconv.ParseInt(v, 10, 64); err == nil {
			beforeID = n
		}
	}

	result, err := h.proposals.List(limit, beforeID, c.QueryParam("type"), c.QueryParam("status"))
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

// Get godoc
// @Summary      获取提案详情
// @Tags         proposals
// @Produce      json
// @Param        id                path   int     true   "提案 ID"
// @Param        include_votes     query  bool    false  "是否包含投票列表"
// @Param        include_comments  query  bool    false  "是否包含评论列表"
// @Success      200  {object}  response.Envelope{data=service.ProposalDetail}
// @Failure      400  {object}  response.Envelope
// @Router       /v1/proposals/{id} [get]
func (h *ProposalHandler) Get(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}

	result, err := h.proposals.GetDetail(id,
		c.QueryParam("include_votes") == "true",
		c.QueryParam("include_comments") == "true",
		queryLimit(c, "vote_limit", 50),
		queryLimit(c, "comment_limit", 50),
	)
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

// VoteSummary godoc
// @Summary      获取提案投票计票摘要
// @Tags         proposals
// @Produce      json
// @Param        id  path  int  true  "提案 ID"
// @Success      200  {object}  response.Envelope{data=service.VoteTally}
// @Router       /v1/proposals/{id}/vote-summary [get]
func (h *ProposalHandler) VoteSummary(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}
	result, err := h.proposals.GetVoteSummary(id)
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

// ListVotes godoc
// @Summary      获取提案投票列表
// @Tags         proposals
// @Produce      json
// @Param        id     path   int  true   "提案 ID"
// @Param        limit  query  int  false  "条数上限"
// @Success      200  {object}  response.Envelope
// @Router       /v1/proposals/{id}/votes [get]
func (h *ProposalHandler) ListVotes(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}
	result, err := h.proposals.ListVotes(id, queryLimit(c, "limit", 50))
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, map[string]interface{}{"items": result})
}

// GetMyVote godoc
// @Summary      获取当前用户对提案的投票
// @Tags         proposals
// @Produce      json
// @Security     BearerAuth
// @Param        id  path  int  true  "提案 ID"
// @Success      200  {object}  response.Envelope{data=service.ProposalVoteItem}
// @Failure      400  {object}  response.Envelope
// @Router       /v1/proposals/{id}/votes/me [get]
func (h *ProposalHandler) GetMyVote(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}
	sns, _, err := h.actorFromContext(c)
	if err != nil {
		return err
	}
	result, err := h.proposals.GetMyVote(id, sns)
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

// ListComments godoc
// @Summary      获取提案评论列表
// @Tags         proposals
// @Produce      json
// @Param        id     path   int  true   "提案 ID"
// @Param        limit  query  int  false  "条数上限"
// @Success      200  {object}  response.Envelope
// @Router       /v1/proposals/{id}/comments [get]
func (h *ProposalHandler) ListComments(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}
	result, err := h.proposals.ListComments(id, queryLimit(c, "limit", 50))
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, map[string]interface{}{"items": result})
}

type createProposalRequest struct {
	Type           string          `json:"type"`
	Title          string          `json:"title"`
	Content        json.RawMessage `json:"content"`
	ResolutionRule string          `json:"resolution_rule"`
}

// Create godoc
// @Summary      创建提案草稿
// @Tags         proposals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        body  body  createProposalRequest  true  "提案内容"
// @Success      200  {object}  response.Envelope{data=service.ProposalResponse}
// @Failure      403  {object}  response.Envelope
// @Router       /v1/proposals [post]
func (h *ProposalHandler) Create(c echo.Context) error {
	var req createProposalRequest
	if err := c.Bind(&req); err != nil || req.Title == "" || len(req.Content) == 0 {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}
	sns, wallet, err := h.actorFromContext(c)
	if err != nil {
		return err
	}

	result, err := h.proposals.CreateDraft(service.CreateProposalInput{
		Type:           req.Type,
		Title:          req.Title,
		Content:        req.Content,
		ResolutionRule: req.ResolutionRule,
		ActorSNS:       sns,
		ActorWallet:    wallet,
	})
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

type updateProposalRequest struct {
	Title   *string         `json:"title"`
	Content json.RawMessage `json:"content"`
}

// Update godoc
// @Summary      修改提案草稿
// @Tags         proposals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path  int                    true  "提案 ID"
// @Param        body  body  updateProposalRequest  true  "更新字段"
// @Success      200  {object}  response.Envelope{data=service.ProposalResponse}
// @Router       /v1/proposals/{id} [patch]
func (h *ProposalHandler) Update(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}
	var req updateProposalRequest
	if err := c.Bind(&req); err != nil {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}
	sns, _, err := h.actorFromContext(c)
	if err != nil {
		return err
	}

	result, err := h.proposals.UpdateDraft(id, service.UpdateProposalInput{
		Title:    req.Title,
		Content:  req.Content,
		ActorSNS: sns,
	})
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

// Close godoc
// @Summary      关闭提案草稿
// @Tags         proposals
// @Produce      json
// @Security     BearerAuth
// @Param        id  path  int  true  "提案 ID"
// @Success      200  {object}  response.Envelope{data=service.ProposalResponse}
// @Router       /v1/proposals/{id}/close [post]
func (h *ProposalHandler) Close(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}
	sns, _, err := h.actorFromContext(c)
	if err != nil {
		return err
	}
	result, err := h.proposals.CloseDraft(id, sns)
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

type publishProposalRequest struct {
	PublicitySeconds    *int `json:"publicity_seconds"`
	VoteDurationSeconds *int `json:"vote_duration_seconds"`
}

// Publish godoc
// @Summary      发布提案（进入公示期）
// @Tags         proposals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path  int                     true  "提案 ID"
// @Param        body  body  publishProposalRequest  false  "公示/投票时长"
// @Success      200  {object}  response.Envelope{data=service.ProposalResponse}
// @Router       /v1/proposals/{id}/publish [post]
func (h *ProposalHandler) Publish(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}
	var req publishProposalRequest
	_ = c.Bind(&req)
	sns, _, err := h.actorFromContext(c)
	if err != nil {
		return err
	}

	result, err := h.proposals.Publish(id, service.PublishProposalInput{
		PublicitySeconds:    req.PublicitySeconds,
		VoteDurationSeconds: req.VoteDurationSeconds,
		ActorSNS:            sns,
	})
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

type castVoteRequest struct {
	VoteResult string `json:"vote_result"`
}

// CastVote godoc
// @Summary      对提案投票
// @Description  节点白名单内 SNS 可投票；同一 SNS 对同一提案仅可投一票
// @Tags         proposals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path  int              true  "提案 ID"
// @Param        body  body  castVoteRequest  true  "投票选项，如 1/-1/0"
// @Success      200  {object}  response.Envelope{data=service.ProposalVoteItem}
// @Failure      403  {object}  response.Envelope
// @Router       /v1/proposals/{id}/votes [post]
func (h *ProposalHandler) CastVote(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}
	var req castVoteRequest
	if err := c.Bind(&req); err != nil || req.VoteResult == "" {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}
	sns, wallet, err := h.actorFromContext(c)
	if err != nil {
		return err
	}

	result, err := h.proposals.CastVote(service.CastVoteInput{
		ProposalID:  id,
		VoteResult:  req.VoteResult,
		ActorSNS:    sns,
		ActorWallet: wallet,
	})
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

type createCommentRequest struct {
	Content  string `json:"content"`
	ParentID *int64 `json:"parent_id"`
	IsReject bool   `json:"is_reject"`
}

// CreateComment godoc
// @Summary      发表评论
// @Tags         proposals
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id    path  int                   true  "提案 ID"
// @Param        body  body  createCommentRequest  true  "评论内容"
// @Success      200  {object}  response.Envelope{data=service.ProposalCommentItem}
// @Router       /v1/proposals/{id}/comments [post]
func (h *ProposalHandler) CreateComment(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}
	var req createCommentRequest
	if err := c.Bind(&req); err != nil || req.Content == "" {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}
	sns, wallet, err := h.actorFromContext(c)
	if err != nil {
		return err
	}

	result, err := h.proposals.CreateComment(service.CreateCommentInput{
		ProposalID:  id,
		Content:     req.Content,
		ParentID:    req.ParentID,
		IsReject:    req.IsReject,
		ActorSNS:    sns,
		ActorWallet: wallet,
	})
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

// Finalize godoc
// @Summary      投票结束后结票
// @Description  根据元规则计票并将状态更新为 vote_passed 或 vote_failed
// @Tags         proposals
// @Produce      json
// @Security     BearerAuth
// @Param        id  path  int  true  "提案 ID"
// @Success      200  {object}  response.Envelope{data=service.ProposalDetail}
// @Router       /v1/proposals/{id}/finalize [post]
func (h *ProposalHandler) Finalize(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil || id <= 0 {
		return response.Fail(c, http.StatusBadRequest, errcode.InvalidParam)
	}
	sns, _, err := h.actorFromContext(c)
	if err != nil {
		return err
	}
	result, err := h.proposals.FinalizeVote(id, sns)
	if err != nil {
		return mapServiceError(c, err)
	}
	return response.OK(c, result)
}

func (h *ProposalHandler) GovernanceNodes(c echo.Context) error {
	return response.OK(c, h.proposals.GovernanceNodes())
}

func (h *ProposalHandler) actorFromContext(c echo.Context) (string, string, error) {
	userID, ok := middleware.UserIDFromContext(c)
	if !ok {
		return "", "", response.Fail(c, http.StatusUnauthorized, errcode.InvalidParam)
	}
	profile, err := h.profile.GetByUserID(userID)
	if err != nil {
		return "", "", mapServiceError(c, err)
	}
	if profile.SNSName == "" {
		return "", "", response.Fail(c, http.StatusBadRequest, errcode.ProposalSNSRequired)
	}
	return profile.SNSName, "", nil
}

func queryLimit(c echo.Context, key string, def int) int {
	if v := c.QueryParam(key); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			return n
		}
	}
	if v := c.QueryParam("limit"); v != "" && key != "limit" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			return n
		}
	}
	return def
}
