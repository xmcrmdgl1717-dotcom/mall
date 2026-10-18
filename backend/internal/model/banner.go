package model

import "time"

type Banner struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	ImageURL  string    `gorm:"size:255" json:"image_url"`
	LinkURL   string    `gorm:"size:255" json:"link_url"`
	Sort      int       `json:"sort"`
	Status    int8      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
