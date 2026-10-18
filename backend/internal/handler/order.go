package handler

import (
	"strconv"

	"mall/backend/internal/pkg/response"
	"mall/backend/internal/repository"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	repo *repository.OrderRepo
}

func NewOrderHandler(repo *repository.OrderRepo) *OrderHandler {
	return &OrderHandler{repo: repo}
}

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
	list, err := h.repo.GetUserOrders(1)
	if err != nil {
		response.Fail(c, 500, "查询失败")
		return
	}
	response.OK(c, list)
}

// ================= 新增：发货操作 =================
func (h *OrderHandler) ShipOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.Fail(c, 400, "invalid order id")
		return
	}

	// 假设状态 2 代表已发货
	if err := h.repo.UpdateStatus(id, 2); err != nil {
		response.Fail(c, 500, "发货失败")
		return
	}
	response.OK(c, "发货成功")
}
