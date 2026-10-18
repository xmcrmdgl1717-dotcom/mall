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

// List 获取商品列表
func (h *ProductHandler) List(c *gin.Context) {
	categoryID, _ := strconv.ParseUint(c.DefaultQuery("category_id", "0"), 10, 64)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "20"))
	dealType := c.Query("deal_type") // 接收前端传来的 deal_type
	
	if page < 1 { page = 1 }
	if size < 1 || size > 100 { size = 20 }

	// 重点修复：这里要传 5 个参数，顺序依次是 categoryID, sort, page, size, dealType
	list, total, err := h.repo.List(categoryID, c.Query("sort"), page, size, dealType)
	if err != nil {
		response.Fail(c, 500, "query failed")
		return
	}
	response.OK(c, gin.H{"list": list, "total": total, "page": page, "size": size})
}

// Detail 获取商品详情 (含 SKU)
func (h *ProductHandler) Detail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Fail(c, 400, "invalid product id")
		return
	}

	product, err := h.repo.GetByID(id)
	if err != nil {
		response.Fail(c, 404, "product not found")
		return
	}
	response.OK(c, product)
}
