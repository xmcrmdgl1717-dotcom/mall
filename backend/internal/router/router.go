package router

import (
    "mall/backend/internal/handler"
    "mall/backend/internal/repository"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func New(db *gorm.DB) *gin.Engine {
    r := gin.Default()
    productRepo := repository.NewProductRepo(db)
    productHandler := handler.NewProductHandler(productRepo)

    api := r.Group("/api/v1")
    {
        api.GET("/products", productHandler.List)
    }
    return r
}
