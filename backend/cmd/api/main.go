package main

import (
	"log"
	"mall/backend/internal/config"
	"mall/backend/internal/repository"
	"mall/backend/internal/router"
)
// ... 前面连接 db 的代码
db, err := repository.NewDB(cfg.DBDSN)
if err != nil {
    log.Fatal("db connect failed: ", err)
}

// ====== 新增这段自动迁移代码 ======
err = db.AutoMigrate(
    &model.Category{},
    &model.Product{},
    &model.ProductSku{},
    &model.Merchant{},
    &model.AdminUser{}, // 你需要先在 model 里建好这个结构体
)
if err != nil {
    log.Println("AutoMigrate failed:", err)
} else {
    log.Println("Database migrated successfully")
}
// =================================

r := router.New(db, cfg.Secret)
// ...

func main() {
	cfg := config.Load()
	db, err := repository.NewDB(cfg.DBDSN)
	if err != nil {
		log.Fatal("db connect failed: ", err)
	}

	// 把 cfg.Secret 传进去
    r := router.New(db, cfg.Secret)
	log.Printf("server start at :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
