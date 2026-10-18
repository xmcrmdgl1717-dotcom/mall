package router

import (
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mall/backend/internal/handler"
	"mall/backend/internal/repository"
)

func New(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// 配置跨域，允许所有来源（开发阶段）
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	productRepo := repository.NewProductRepo(db)
	productHandler := handler.NewProductHandler(productRepo)

	api := r.Group("/api/v1")
	{
		api.GET("/products", productHandler.List)
	}
	return r
}
