package handler

import (
	"mall/backend/internal/model"
	"mall/backend/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct { db *gorm.DB }
func NewUserHandler(db *gorm.DB) *UserHandler { return &UserHandler{db: db} }

// GetBuyers 获取所有买家列表 (对应后台用户管理)
func (h *UserHandler) GetBuyers(c *gin.Context) {
	var users []model.User
	h.db.Order("created_at DESC").Find(&users)
	response.OK(c, users)
}
