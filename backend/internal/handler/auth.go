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

	// 验证账号密码
	var adminId uint64
	err := h.db.Raw("SELECT id FROM admin_user WHERE username = ? AND password = ?", req.Username, req.Password).Scan(&adminId).Error
	if err != nil || adminId == 0 {
		response.Fail(c, 401, "账号或密码错误")
		return
	}

	// 签发 Token
	token, _ := jwt.GenerateToken(h.secret, adminId)
	response.OK(c, gin.H{"token": token})
}
