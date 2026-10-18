package handler

import (
	"mall/backend/internal/pkg/response"
	"mall/backend/internal/repository"
	"github.com/gin-gonic/gin"
)

type DashboardHandler struct { repo *repository.DashboardRepo }
func NewDashboardHandler(repo *repository.DashboardRepo) *DashboardHandler { return &DashboardHandler{repo: repo} }

func (h *DashboardHandler) GetStats(c *gin.Context) {
	stats, err := h.repo.GetStats()
	if err != nil { response.Fail(c, 500, "获取失败"); return }
	response.OK(c, stats)
}
