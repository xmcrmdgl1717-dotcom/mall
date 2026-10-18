package main

import (
	"log"
	"mall/backend/internal/config"
	"mall/backend/internal/repository"
	"mall/backend/internal/router"
)

func main() {
	cfg := config.Load()
	db, err := repository.NewDB(cfg.DBDSN)
	if err != nil {
		log.Fatal("db connect failed: ", err)
	}

	r := router.New(db)
	log.Printf("server start at :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
