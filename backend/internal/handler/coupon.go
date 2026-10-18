package handler

import (
	"mall/backend/internal/model"
	"mall/backend/internal/pkg/response"
	"mall/backend/internal/repository"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CouponHandler struct { repo *repository.CouponRepo }
func NewCouponHandler(repo *repository.CouponRepo) *CouponHandler { return &CouponHandler{repo: repo} }

// AdminList 后台列表
func (h *CouponHandler) AdminList(c *gin.Context) {
	list, _ := h.repo.ListAll()
	response.OK(c, list)
}

type CreateCouponReq struct {
	Name       string  `json:"name" binding:"required"`
	Amount     float64 `json:"amount" binding:"required"`
	MinAmount  float64 `json:"min_amount"`
	TotalCount int     `json:"total_count" binding:"required"`
	Days       int     `json:"days"` // 有效期天数
}

// Create 后台发布优惠券
func (h *CouponHandler) Create(c *gin.Context) {
	var req CreateCouponReq
	if err := c.ShouldBindJSON(&req); err != nil { response.Fail(c, 400, "参数错误"); return }

	coupon := &model.Coupon{
		Name:       req.Name,
		Amount:     req.Amount,
		MinAmount:  req.MinAmount,
		TotalCount: req.TotalCount,
		StartTime:  time.Now(),
		EndTime:    time.Now().AddDate(0, 0, req.Days),
	}
	if err := h.repo.Create(coupon); err != nil { response.Fail(c, 500, "创建失败"); return }
	response.OK(c, coupon)
}

// ListActive 前台领券中心
func (h *CouponHandler) ListActive(c *gin.Context) {
	list, _ := h.repo.ListActive()
	response.OK(c, list)
}

// Claim 用户领券
func (h *CouponHandler) Claim(c *gin.Context) {
	couponID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.repo.Claim(1, couponID); err != nil { // UserID 暂时写死为 1
		response.Fail(c, 400, err.Error())
		return
	}
	response.OK(c, "领取成功")
}

// MyCoupons 我的优惠券
func (h *CouponHandler) MyCoupons(c *gin.Context) {
	list, _ := h.repo.GetUserCoupons(1)
	response.OK(c, list)
}
