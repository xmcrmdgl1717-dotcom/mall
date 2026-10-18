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
	if err := c.ShouldBindJSON(&req); err != nil { response.Fail(c, 400, "参数错误"); return }

	userID := c.GetUint64("buyer_id") // 获取拦截器传过来的用户 ID
	cart := &model.Cart{
		UserID:    userID, 
		ProductID: req.ProductID,
		SkuID:     req.SkuID,
		Quantity:  req.Quantity,
		Selected:  1,
	}

	if err := h.repo.AddToCart(cart); err != nil { response.Fail(c, 500, "添加失败"); return }
	response.OK(c, "加入购物车成功")
}

func (h *CartHandler) List(c *gin.Context) {
	userID := c.GetUint64("buyer_id")
	list, err := h.repo.GetCartList(userID)
	if err != nil { response.Fail(c, 500, "查询失败"); return }
	response.OK(c, list)
}

func (h *CartHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	var id uint64
	for _, char := range idStr { if char >= '0' && char <= '9' { id = id*10 + uint64(char-'0') } }
	
	userID := c.GetUint64("buyer_id")
	if err := h.repo.DeleteCartItem(id, userID); err != nil { response.Fail(c, 500, "删除失败"); return }
	response.OK(c, "删除成功")
}
