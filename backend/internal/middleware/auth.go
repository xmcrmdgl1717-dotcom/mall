package middleware

import (
	"strings"

	"mall/backend/internal/pkg/jwt"
	"mall/backend/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

func BuyerAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			response.Fail(c, 401, "请先登录")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		id, userType, err := jwt.ParseToken(secret, tokenString)
		if err != nil || userType != "buyer" {
			response.Fail(c, 401, "登录已过期，请重新登录")
			c.Abort()
			return
		}

		c.Set("buyer_id", id)
		c.Next()
	}
}
