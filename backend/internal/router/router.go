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
	// ... CORS 代码保持不变 ...

	productRepo := repository.NewProductRepo(db)
	productHandler := handler.NewProductHandler(productRepo)
	
	categoryRepo := repository.NewCategoryRepo(db)
	categoryHandler := handler.NewCategoryHandler(categoryRepo)
   // 在已有的路由组里加上：
    merchantRepo := repository.NewMerchantRepo(db)
    merchantHandler := handler.NewMerchantHandler(merchantRepo)

api.GET("/merchants/top", merchantHandler.ListTop)
	api := r.Group("/api/v1")
	{
		api.GET("/products", productHandler.List)
        api.GET("/products/:id", productHandler.Detail) // 新增详情接口
        api.GET("/categories", categoryHandler.List)
	}
	return r
}
