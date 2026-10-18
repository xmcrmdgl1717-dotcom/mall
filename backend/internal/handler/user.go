package handler

import (
	"mall/backend/internal/model"
	"mall/backend/internal/pkg/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

// GetBuyers 获取所有买家列表
func (h *UserHandler) GetBuyers(c *gin.Context) {
	var users []model.User
	if err := h.db.Order("created_at DESC").Find(&users).Error; err != nil {
		response.Fail(c, 500, "获取用户失败")
		return
	}
	response.OK(c, users)
}
