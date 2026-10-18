package handler

import (
	"mall/backend/internal/model"
	"mall/backend/internal/pkg/response"
	"mall/backend/internal/repository"
	"github.com/gin-gonic/gin"
)

type CartHandler struct { repo *repository.CartRepo }
func NewCartHandler(repo *repository.CartRepo) *CartHandler { return &CartHandler{repo: repo} }

type AddCartReq struct {
	ProductID uint64 `json:"product_id" binding:"required"`
	SkuID     uint64 `json:"sku_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required,gt=0"`
}

func (h *CartHandler) Add(c *gin.Context) {
	var req AddCartReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	// 暂时写死 UserID 为 1
	cart := &model.Cart{
		UserID:    1, 
		ProductID: req.ProductID,
		SkuID:     req.SkuID,
		Quantity:  req.Quantity,
		Selected:  1,
	}

	if err := h.repo.AddToCart(cart); err != nil {
		response.Fail(c, 500, "添加失败")
		return
	}
	response.OK(c, "加入购物车成功")
}

func (h *CartHandler) List(c *gin.Context) {
	list, err := h.repo.GetCartList(1) // 暂时写死 1
	if err != nil {
		response.Fail(c, 500, "查询失败")
		return
	}
	response.OK(c, list)
}

func (h *CartHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	// 简单解析
	var id uint64
	for _, char := range idStr {
		if char >= '0' && char <= '9' {
			id = id*10 + uint64(char-'0')
		}
	}
	if err := h.repo.DeleteCartItem(id, 1); err != nil {
		response.Fail(c, 500, "删除失败")
		return
	}
	response.OK(c, "删除成功")
}
