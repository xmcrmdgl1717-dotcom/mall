package handler

import (
	"mall/backend/internal/model"
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

// 后台管理员登录
func (h *AuthHandler) AdminLogin(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil { response.Fail(c, 400, "参数错误"); return }

	type AdminUser struct { ID uint64; RoleID uint64 }
	var admin AdminUser
	h.db.Table("admin_users").Where("username = ? AND password = ?", req.Username, req.Password).First(&admin)

	if admin.ID == 0 { response.Fail(c, 401, "账号或密码错误"); return }

	token, _ := jwt.GenerateToken(h.secret, admin.ID, admin.RoleID, false)
	response.OK(c, gin.H{"token": token, "role_id": admin.RoleID})
}

// 买家注册
func (h *AuthHandler) BuyerRegister(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil { response.Fail(c, 400, "参数错误"); return }

	var count int64
	h.db.Model(&model.User{}).Where("username = ?", req.Username).Count(&count)
	if count > 0 { response.Fail(c, 400, "用户名已存在"); return }

	user := &model.User{Username: req.Username, Password: req.Password, Status: 1}
	if err := h.db.Create(user).Error; err != nil { response.Fail(c, 500, "注册失败"); return }
	response.OK(c, "注册成功")
}

// 买家登录
func (h *AuthHandler) BuyerLogin(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil { response.Fail(c, 400, "参数错误"); return }

	var user model.User
	h.db.Where("username = ? AND password = ?", req.Username, req.Password).First(&user)
	if user.ID == 0 { response.Fail(c, 401, "账号或密码错误"); return }
	if user.Status == 0 { response.Fail(c, 403, "账号已被封禁"); return }

	token, _ := jwt.GenerateToken(h.secret, user.ID, 0, true)
	response.OK(c, gin.H{"token": token, "username": user.Username})
}
