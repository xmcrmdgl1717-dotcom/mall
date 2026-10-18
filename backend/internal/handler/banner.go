package handler

import (
	"mall/backend/internal/model"
	"mall/backend/internal/pkg/response"
	"mall/backend/internal/repository"
	"github.com/gin-gonic/gin"
)

type BannerHandler struct { repo *repository.BannerRepo }
func NewBannerHandler(repo *repository.BannerRepo) *BannerHandler { return &BannerHandler{repo: repo} }

// ListActive 前台获取启用的 Banner
func (h *BannerHandler) ListActive(c *gin.Context) {
	list, err := h.repo.ListActive()
	if err != nil {
		response.Fail(c, 500, "获取失败")
		return
	}
	response.OK(c, list)
}

// ListAll 后台获取所有 Banner
func (h *BannerHandler) ListAll(c *gin.Context) {
	list, err := h.repo.ListAll()
	if err != nil {
		response.Fail(c, 500, "获取失败")
		return
	}
	response.OK(c, list)
}

type BannerReq struct {
	ImageURL string `json:"image_url" binding:"required"`
	LinkURL  string `json:"link_url"`
	Sort     int    `json:"sort"`
}

func (h *BannerHandler) Create(c *gin.Context) {
	var req BannerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}
	banner := &model.Banner{
		ImageURL: req.ImageURL,
		LinkURL:  req.LinkURL,
		Sort:     req.Sort,
		Status:   1,
	}
	if err := h.repo.Create(banner); err != nil {
		response.Fail(c, 500, "创建失败")
		return
	}
	response.OK(c, banner)
}

func (h *BannerHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	var id uint64
	for _, char := range idStr { if char >= '0' && char <= '9' { id = id*10 + uint64(char-'0') } }
	
	if err := h.repo.Delete(id); err != nil {
		response.Fail(c, 500, "删除失败")
		return
	}
	response.OK(c, "删除成功")
}
