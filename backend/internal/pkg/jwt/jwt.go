package jwt

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(secret string, adminId uint64, roleId uint64) (string, error) {
	claims := jwt.MapClaims{
		"admin_id": adminId,
		"role_id":  roleId, // 把角色 ID 也塞进 Token
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
