package model

import "time"

type AdminUser struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"uniqueIndex;size:50" json:"username"`
	Password  string    `gorm:"size:100" json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
