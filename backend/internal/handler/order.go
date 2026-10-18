package handler

import (
	"mall/backend/internal/pkg/response"
	"mall/backend/internal/repository"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct { repo *repository.OrderRepo }
func NewOrderHandler(repo *repository.OrderRepo) *OrderHandler { return &OrderHandler{repo: repo} }

type CheckoutReq struct {
	AddressName   string   `json:"address_name" binding:"required"`
	AddressPhone  string   `json:"address_phone" binding:"required"`
	AddressDetail string   `json:"address_detail" binding:"required"`
	CartIDs       []uint64 `json:"cart_ids" binding:"required"`
}

func (h *OrderHandler) Create(c *gin.Context) {
	var req CheckoutReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	// 暂时写死 UserID 为 1
	order, err := h.repo.CreateOrder(repository.CreateOrderReq{
		UserID:        1,
		AddressName:   req.AddressName,
		AddressPhone:  req.AddressPhone,
		AddressDetail: req.AddressDetail,
		CartIDs:       req.CartIDs,
	})

	if err != nil {
		response.Fail(c, 500, "下单失败: "+err.Error())
		return
	}
	response.OK(c, order)
}

func (h *OrderHandler) List(c *gin.Context) {
	list, err := h.repo.GetUserOrders(1) // 暂时写死 1
	if err != nil {
		response.Fail(c, 500, "查询失败")
		return
	}
	response.OK(c, list)
}
