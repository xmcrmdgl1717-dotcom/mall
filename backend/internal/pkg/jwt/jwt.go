package jwt

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

// 生成 Token
func GenerateToken(secret string, adminId uint64) (string, error) {
	claims := jwt.MapClaims{
		"admin_id": adminId,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 24小时过期
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
