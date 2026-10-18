package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(dsn string) (*gorm.DB, error) {
	// Render 提供的连接串有时带前缀，GORM 直接支持 standard 格式
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
