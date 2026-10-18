package router

import (
	"time"

	"mall/backend/internal/handler"
	"mall/backend/internal/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// 配置跨域
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 初始化 Repo
	productRepo := repository.NewProductRepo(db)
	categoryRepo := repository.NewCategoryRepo(db)
	merchantRepo := repository.NewMerchantRepo(db)

	// 初始化 Handler
	productHandler := handler.NewProductHandler(productRepo)
	categoryHandler := handler.NewCategoryHandler(categoryRepo)
	merchantHandler := handler.NewMerchantHandler(merchantRepo)

	// 注册路由
	api := r.Group("/api/v1")
	{
		api.GET("/products", productHandler.List)
		api.GET("/products/:id", productHandler.Detail)
		api.GET("/categories", categoryHandler.List)
		api.GET("/merchants/top", merchantHandler.ListTop)
	}
	
	return r
}
