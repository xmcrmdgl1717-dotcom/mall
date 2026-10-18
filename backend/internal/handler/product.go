package handler

import (
	"strconv"

	"mall/backend/internal/model"
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
	dealType := c.Query("deal_type")
	
	if page < 1 { page = 1 }
	if size < 1 || size > 100 { size = 20 }

	list, total, err := h.repo.List(categoryID, c.Query("sort"), page, size, dealType)
	if err != nil {
		response.Fail(c, 500, "query failed")
		return
	}
	response.OK(c, gin.H{"list": list, "total": total, "page": page, "size": size})
}

// Detail 获取商品详情
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

// ================= 新增：支持 SKU 的商品发布 =================

type SkuReq struct {
	Specs string  `json:"specs"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

type CreateProductReq struct {
	Title      string   `json:"title" binding:"required"`
	MainImage  string   `json:"main_image"`
	CategoryID uint64   `json:"category_id"`
	Skus       []SkuReq `json:"skus" binding:"required"` 
}

func (h *ProductHandler) Create(c *gin.Context) {
	var req CreateProductReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误: "+err.Error())
		return
	}

	// 组装 SKU 数组
	var skus []model.ProductSku
	for _, s := range req.Skus {
		skus = append(skus, model.ProductSku{
			Specs: s.Specs,
			Price: s.Price,
			Stock: s.Stock,
		})
	}

	product := &model.Product{
		Title:      req.Title,
		MainImage:  req.MainImage,
		CategoryID: req.CategoryID,
		MerchantID: 1, // 暂时写死
		Status:     1, // 默认上架
		Skus:       skus, 
	}

	if err := h.repo.Create(product); err != nil {
		response.Fail(c, 500, "保存失败: "+err.Error())
		return
	}
	response.OK(c, product)
}
