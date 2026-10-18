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
	cartRepo := repository.NewCartRepo(db)
	orderRepo := repository.NewOrderRepo(db)

	// 初始化 Handler
	productHandler := handler.NewProductHandler(productRepo)
	categoryHandler := handler.NewCategoryHandler(categoryRepo)
	merchantHandler := handler.NewMerchantHandler(merchantRepo)
	authHandler := handler.NewAuthHandler(db, secret)
	cartHandler := handler.NewCartHandler(cartRepo)
	orderHandler := handler.NewOrderHandler(orderRepo)
	
	// 初始化 UserHandler (新增)
	userHandler := handler.NewUserHandler(db)

	// 注册路由组
	api := r.Group("/api/v1")
	{
		// 登录与用户管理
		api.POST("/admin/login", authHandler.Login)
		api.GET("/admin/users", userHandler.GetBuyers) // 新增：获取用户列表

		// 商品接口
		api.GET("/products", productHandler.List)
		api.GET("/products/:id", productHandler.Detail)
		api.POST("/products", productHandler.Create)

		// 分类与店铺接口
		api.GET("/categories", categoryHandler.List)
		api.GET("/merchants/top", merchantHandler.ListTop)

		// 购物车接口
		api.GET("/cart", cartHandler.List)
		api.POST("/cart", cartHandler.Add)
		api.DELETE("/cart/:id", cartHandler.Delete)

		// 订单接口
		api.POST("/orders", orderHandler.Create)
		api.GET("/orders", orderHandler.List)
	}
	
	return r
}
