package handler

import (
	"mall/backend/internal/pkg/response"
	"mall/backend/internal/repository"
	"github.com/gin-gonic/gin"
)

type MerchantHandler struct { 
	repo *repository.MerchantRepo 
}

func NewMerchantHandler(repo *repository.MerchantRepo) *MerchantHandler { 
	return &MerchantHandler{repo: repo} 
}

func (h *MerchantHandler) ListTop(c *gin.Context) {
	list, err := h.repo.GetTopMerchants(3)
	if err != nil {
		response.Fail(c, 500, "query failed")
		return
	}
	response.OK(c, list)
}
