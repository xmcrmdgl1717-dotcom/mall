package router

import (
	"time"

	"mall/backend/internal/handler"
	"mall/backend/internal/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(db *gorm.DB, secret string) *gin.Engine {
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
	authHandler := handler.NewAuthHandler(db, secret) // 注意这里传了 secret

	// 注册路由组
	api := r.Group("/api/v1")
	{
		// 登录接口
		api.POST("/admin/login", authHandler.Login)

		// 商品接口
		api.GET("/products", productHandler.List)
		api.GET("/products/:id", productHandler.Detail)
		api.POST("/products", productHandler.Create)

		// 分类与店铺接口
		api.GET("/categories", categoryHandler.List)
		api.GET("/merchants/top", merchantHandler.ListTop)
	}
	
	return r
}
