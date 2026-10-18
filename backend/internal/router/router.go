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

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 开启静态文件服务，让外界可以访问 /uploads/xxx.jpg
	r.Static("/uploads", "./uploads")

	// 初始化所有 Handler
	categoryHandler := handler.NewCategoryHandler(repository.NewCategoryRepo(db))
	dashboardHandler := handler.NewDashboardHandler(repository.NewDashboardRepo(db))
	productHandler := handler.NewProductHandler(repository.NewProductRepo(db))
	merchantHandler := handler.NewMerchantHandler(repository.NewMerchantRepo(db))
	cartHandler := handler.NewCartHandler(repository.NewCartRepo(db))
	orderHandler := handler.NewOrderHandler(repository.NewOrderRepo(db))
	authHandler := handler.NewAuthHandler(db, secret)
	userHandler := handler.NewUserHandler(db)
	bannerHandler := handler.NewBannerHandler(repository.NewBannerRepo(db))
	uploadHandler := handler.NewUploadHandler() // 新增

	api := r.Group("/api/v1")
	{
				// 商家管理 (后台)
		api.GET("/admin/merchants", merchantHandler.ListAll)
		api.PUT("/admin/merchants/:id/status", merchantHandler.UpdateStatus)
		
		api.POST("/admin/login", authHandler.Login)
		api.GET("/admin/users", userHandler.GetBuyers)
		
		// 图片上传接口
		api.POST("/admin/upload", uploadHandler.UploadImage)

		// 看板与类目
		api.GET("/admin/dashboard", dashboardHandler.GetStats)
		api.POST("/admin/categories", categoryHandler.Create)
		api.DELETE("/admin/categories/:id", categoryHandler.Delete)
		api.GET("/categories", categoryHandler.List)

		// 商品
		api.GET("/products", productHandler.List)
		api.GET("/products/:id", productHandler.Detail)
		api.POST("/products", productHandler.Create)

		api.GET("/merchants/top", merchantHandler.ListTop)

		// 购物车与订单
		api.GET("/cart", cartHandler.List)
		api.POST("/cart", cartHandler.Add)
		api.DELETE("/cart/:id", cartHandler.Delete)
		api.POST("/orders", orderHandler.Create)
		api.GET("/orders", orderHandler.List)
		api.PUT("/orders/:id/ship", orderHandler.ShipOrder)
		
		// 新增 Banner 路由
		api.GET("/banners", bannerHandler.ListActive)
		api.GET("/admin/banners", bannerHandler.ListAll)
		api.POST("/admin/banners", bannerHandler.Create)
		api.DELETE("/admin/banners/:id", bannerHandler.Delete)
	}
	
	return r
}
