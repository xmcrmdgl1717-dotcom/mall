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

	// 允许跨域
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 初始化所有 Handler
	productHandler := handler.NewProductHandler(repository.NewProductRepo(db))
	categoryHandler := handler.NewCategoryHandler(repository.NewCategoryRepo(db))
	merchantHandler := handler.NewMerchantHandler(repository.NewMerchantRepo(db))
	cartHandler := handler.NewCartHandler(repository.NewCartRepo(db))
	orderHandler := handler.NewOrderHandler(repository.NewOrderRepo(db))
	authHandler := handler.NewAuthHandler(db, secret)
	userHandler := handler.NewUserHandler(db)

	// 注册路由
	api := r.Group("/api/v1")
	{
		// 鉴权与用户
		api.POST("/admin/login", authHandler.Login)
		api.GET("/admin/users", userHandler.GetBuyers)

		// 商品
		api.GET("/products", productHandler.List)
		api.GET("/products/:id", productHandler.Detail)
		api.POST("/products", productHandler.Create)

		// 类目与店铺
		api.GET("/categories", categoryHandler.List)
		api.GET("/merchants/top", merchantHandler.ListTop)

		// 购物车
		api.GET("/cart", cartHandler.List)
		api.POST("/cart", cartHandler.Add)
		api.DELETE("/cart/:id", cartHandler.Delete)

		// 订单
		api.POST("/orders", orderHandler.Create)
		api.GET("/orders", orderHandler.List)
	}
	
	return r
}
