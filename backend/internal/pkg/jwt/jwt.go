package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 生成 Token
func GenerateToken(secret string, id uint64, roleId uint64, isBuyer bool) (string, error) {
	userType := "admin"
	if isBuyer {
		userType = "buyer"
	}

	claims := jwt.MapClaims{
		"id":        id,
		"role_id":   roleId,
		"user_type": userType,
		"exp":       time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// 解析 Token
func ParseToken(secret, tokenString string) (uint64, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return 0, "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", errors.New("invalid claims")
	}

	id := uint64(claims["id"].(float64))
	userType, _ := claims["user_type"].(string)
	return id, userType, nil
}
