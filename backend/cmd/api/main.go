package main

import (
	"log"

	"mall/backend/internal/config"
	"mall/backend/internal/model"
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
	// 注意：这里已经加上了 Order 和 OrderItem，启动后 Render 会自动建表
	err = db.AutoMigrate(
		&model.Category{},
		&model.Product{},
		&model.ProductSku{},
		&model.Merchant{},
		&model.AdminUser{},
		&model.Cart{},
		&model.Order{},     // 新增：订单表
		&model.OrderItem{}, // 新增：订单明细表
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
