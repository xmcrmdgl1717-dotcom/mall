package main

import (
	"log"

	"mall/backend/internal/config"
	"mall/backend/internal/model"
	"mall/backend/internal/repository"
	"mall/backend/internal/router"
)

func main() {
	cfg := config.Load()
	db, err := repository.NewDB(cfg.DBDSN)
	if err != nil {
		log.Fatal("db connect failed: ", err)
	}

	// 自动迁移建表 (这里加上了 model.Cart)
	err = db.AutoMigrate(
		&model.Category{},
		&model.Product{},
		&model.ProductSku{},
		&model.Merchant{},
		&model.AdminUser{},
		&model.Cart{}, // 新增
	)
	if err != nil {
		log.Println("AutoMigrate failed:", err)
	} else {
		log.Println("Database migrated successfully")
	}

	// 初始化默认管理员账号
	var count int64
	db.Model(&model.AdminUser{}).Where("username = ?", "admin").Count(&count)
	if count == 0 {
		db.Create(&model.AdminUser{
			Username: "admin",
			Password: "admin123",
		})
		log.Println("Default admin account created: admin / admin123")
	}

	r := router.New(db, cfg.Secret)
	log.Printf("server start at :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
