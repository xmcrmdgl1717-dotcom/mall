package handler

import (
	"strconv"

	"mall/backend/internal/pkg/response"
	"mall/backend/internal/repository"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	repo *repository.ProductRepo
}

func NewProductHandler(repo *repository.ProductRepo) *ProductHandler {
	return &ProductHandler{repo: repo}
}

func (h *ProductHandler) List(c *gin.Context) {
	categoryID, _ := strconv.ParseUint(c.DefaultQuery("category_id", "0"), 10, 64)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "20"))
	dealType := c.Query("deal_type") // 新增：daily 或 hot
	
	if page < 1 { page = 1 }
	if size < 1 || size > 100 { size = 20 }

	list, total, err := h.repo.List(categoryID, c.Query("sort"), page, size, dealType)
	if err != nil {
		response.Fail(c, 500, "query failed")
		return
	}
	response.OK(c, gin.H{"list": list, "total": total, "page": page, "size": size})
}
