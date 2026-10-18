package model

import "time"

// 管理员表
type AdminUser struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"uniqueIndex;size:50" json:"username"`
	Password  string    `gorm:"size:100" json:"password"`
	RoleID    uint64    `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
}

// 角色表
type Role struct {
	ID          uint64    `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:50" json:"name"`
	Description string    `gorm:"size:100" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// 买家/用户表 (加了 Password)
type User struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"uniqueIndex;size:50" json:"username"`
	Password  string    `gorm:"size:100" json:"-"` // json:"-" 表示不返回给前端
	Phone     string    `gorm:"size:20" json:"phone"`
	Status    int8      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
