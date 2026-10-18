package handler

import (
	"mall/backend/internal/pkg/jwt"
	"mall/backend/internal/pkg/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	db     *gorm.DB
	secret string
}

func NewAuthHandler(db *gorm.DB, secret string) *AuthHandler {
	return &AuthHandler{db: db, secret: secret}
}

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	// 1. 定义临时结构体接收查询结果
	type AdminUser struct {
		ID     uint64
		RoleID uint64
	}
	var admin AdminUser

	// 2. 验证账号密码（使用 GORM 原生查询）
	h.db.Table("admin_users").Where("username = ? AND password = ?", req.Username, req.Password).First(&admin)

	if admin.ID == 0 {
		response.Fail(c, 401, "账号或密码错误")
		return
	}

	// 3. 签发 Token (必须传 3 个参数，带上 admin.RoleID)
	token, _ := jwt.GenerateToken(h.secret, admin.ID, admin.RoleID)
	response.OK(c, gin.H{"token": token, "role_id": admin.RoleID})
}
