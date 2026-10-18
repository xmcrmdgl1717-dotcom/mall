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

	// 自动迁移建表
	err = db.AutoMigrate(
		&model.Category{},
		&model.Product{},
		&model.ProductSku{},
		&model.Merchant{},
		&model.AdminUser{},
		&model.Cart{},
		&model.Order{},
		&model.OrderItem{},
		&model.Role{},
		&model.User{},
	)
	if err != nil {
		log.Println("AutoMigrate failed:", err)
	} else {
		log.Println("Database migrated successfully")
	}

	// ==========================================
	// 核心修复：自动创建超级管理员账号
	// ==========================================
	var roleCount int64
	db.Model(&model.Role{}).Count(&roleCount)
	if roleCount == 0 {
		db.Create(&model.Role{Name: "超级管理员", Description: "拥有所有权限"})
	}

	var adminCount int64
	db.Model(&model.AdminUser{}).Where("username = ?", "admin").Count(&adminCount)
	if adminCount == 0 {
		db.Create(&model.AdminUser{
			Username: "admin",
			Password: "admin123", // 如果登录失败，密码就绝对是这个
			RoleID:   1,
		})
		log.Println("✅ 默认管理员账号已自动生成: admin / admin123")
	}

	// 启动服务
	r := router.New(db, cfg.Secret)
	log.Printf("server start at :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
