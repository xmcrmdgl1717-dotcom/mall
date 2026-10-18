package router

import (
	"time"

	"mall/backend/internal/handler"
	"mall/backend/internal/middleware"
	"mall/backend/internal/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(db *gorm.DB, secret string) *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Static("/uploads", "./uploads")

	categoryHandler := handler.NewCategoryHandler(repository.NewCategoryRepo(db))
	dashboardHandler := handler.NewDashboardHandler(repository.NewDashboardRepo(db))
	productHandler := handler.NewProductHandler(repository.NewProductRepo(db))
	merchantHandler := handler.NewMerchantHandler(repository.NewMerchantRepo(db))
	cartHandler := handler.NewCartHandler(repository.NewCartRepo(db))
	orderHandler := handler.NewOrderHandler(repository.NewOrderRepo(db))
	authHandler := handler.NewAuthHandler(db, secret)
	userHandler := handler.NewUserHandler(db)
	bannerHandler := handler.NewBannerHandler(repository.NewBannerRepo(db))
	uploadHandler := handler.NewUploadHandler()

	api := r.Group("/api/v1")
	{
		// --- 公开接口 ---
		api.POST("/admin/login", authHandler.AdminLogin) // 注意：后台登录改为了 AdminLogin
		api.POST("/buyer/register", authHandler.BuyerRegister)
		api.POST("/buyer/login", authHandler.BuyerLogin)
		
		api.GET("/products", productHandler.List)
		api.GET("/products/:id", productHandler.Detail)
		api.GET("/categories", categoryHandler.List)
		api.GET("/merchants/top", merchantHandler.ListTop)
		api.GET("/banners", bannerHandler.ListActive)

		// --- 需要买家登录的接口 (挂载中间件) ---
		buyerGroup := api.Group("")
		buyerGroup.Use(middleware.BuyerAuth(secret))
		{
			buyerGroup.GET("/cart", cartHandler.List)
			buyerGroup.POST("/cart", cartHandler.Add)
			buyerGroup.DELETE("/cart/:id", cartHandler.Delete)
			buyerGroup.POST("/orders", orderHandler.Create)
			buyerGroup.GET("/orders", orderHandler.List)
		}

		// --- 后台管理接口 ---
		adminGroup := api.Group("/admin")
		{
			adminGroup.GET("/users", userHandler.GetBuyers)
			adminGroup.GET("/dashboard", dashboardHandler.GetStats)
			adminGroup.POST("/upload", uploadHandler.UploadImage)
			adminGroup.POST("/categories", categoryHandler.Create)
			adminGroup.DELETE("/categories/:id", categoryHandler.Delete)
			adminGroup.POST("/products", productHandler.Create)
			adminGroup.GET("/banners", bannerHandler.ListAll)
			adminGroup.POST("/banners", bannerHandler.Create)
			adminGroup.DELETE("/banners/:id", bannerHandler.Delete)
			adminGroup.GET("/merchants", merchantHandler.ListAll)
			adminGroup.PUT("/merchants/:id/status", merchantHandler.UpdateStatus)
			adminGroup.PUT("/orders/:id/ship", orderHandler.ShipOrder)
		}
	}
	
	return r
}
