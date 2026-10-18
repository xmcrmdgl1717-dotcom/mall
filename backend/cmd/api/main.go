package main

import (
	"log"

	"mall/backend/internal/config"
	"mall/backend/internal/model" // 引入 model
	"mall/backend/internal/repository"
	"mall/backend/internal/router"
)

func main() {
	// 1. 加载配置
	cfg := config.Load()

	// 2. 连接数据库
	db, err := repository.NewDB(cfg.DBDSN)
	if err != nil {
		log.Fatal("db connect failed: ", err)
	}

	// 3. 自动迁移建表 (AutoMigrate)
	err = db.AutoMigrate(
		&model.Category{},
		&model.Product{},
		&model.ProductSku{},
		&model.Merchant{},
		&model.AdminUser{},
	)
	if err != nil {
		log.Println("AutoMigrate failed:", err)
	} else {
		log.Println("Database migrated successfully")
	}

	// 4. 初始化默认管理员账号
	var count int64
	db.Model(&model.AdminUser{}).Where("username = ?", "admin").Count(&count)
	if count == 0 {
		db.Create(&model.AdminUser{
			Username: "admin",
			Password: "admin123",
		})
		log.Println("Default admin account created: admin / admin123")
	}

	// 5. 启动 HTTP 服务
	r := router.New(db, cfg.Secret)
	log.Printf("server start at :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
