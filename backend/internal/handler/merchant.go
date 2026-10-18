package handler

import (
	"mall/backend/internal/pkg/response"
	"mall/backend/internal/repository"
	"strconv"
	"github.com/gin-gonic/gin"
)

type MerchantHandler struct { repo *repository.MerchantRepo }
func NewMerchantHandler(repo *repository.MerchantRepo) *MerchantHandler { return &MerchantHandler{repo: repo} }

func (h *MerchantHandler) ListTop(c *gin.Context) {
	list, err := h.repo.GetTopMerchants(3)
	if err != nil {
		response.Fail(c, 500, "query failed")
		return
	}
	response.OK(c, list)
}

func (h *MerchantHandler) ListAll(c *gin.Context) {
	list, err := h.repo.ListAll()
	if err != nil {
		response.Fail(c, 500, "query failed")
		return
	}
	response.OK(c, list)
}

type UpdateStatusReq struct {
	Status int8 `json:"status"`
}

func (h *MerchantHandler) UpdateStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	
	var req UpdateStatusReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	if err := h.repo.UpdateStatus(id, req.Status); err != nil {
		response.Fail(c, 500, "更新失败")
		return
	}
	response.OK(c, "更新成功")
}
