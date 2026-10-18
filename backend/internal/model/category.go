package model

import "time"

type Category struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	Sort      int       `json:"sort"`
	CreatedAt time.Time `json:"created_at"`
}
