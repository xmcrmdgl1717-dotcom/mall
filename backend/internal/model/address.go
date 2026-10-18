package model

import "time"

type Address struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	UserID    uint64    `json:"user_id"`
	Name      string    `gorm:"size:50" json:"name"`       // 收货人
	Phone     string    `gorm:"size:20" json:"phone"`      // 电话
	Province  string    `gorm:"size:50" json:"province"`   // 省份
	City      string    `gorm:"size:50" json:"city"`       // 城市
	Detail    string    `gorm:"size:255" json:"detail"`    // 详细地址
	IsDefault int8      `gorm:"default:0" json:"is_default"` // 1为默认地址
	CreatedAt time.Time `json:"created_at"`
}
