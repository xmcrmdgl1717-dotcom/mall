package model

import "time"

// 管理员表 (更新：增加角色关联)
type AdminUser struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"uniqueIndex;size:50" json:"username"`
	Password  string    `gorm:"size:100" json:"password"`
	RoleID    uint64    `json:"role_id"` // 新增：关联角色
	CreatedAt time.Time `json:"created_at"`
}

// 角色表 (新增)
type Role struct {
	ID          uint64    `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:50" json:"name"` // 如：超级管理员、运营、客服
	Description string    `gorm:"size:100" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// 买家/用户表 (新增：对应前端的注册用户)
type User struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"uniqueIndex;size:50" json:"username"`
	Phone     string    `gorm:"size:20" json:"phone"`
	Status    int8      `json:"status"` // 1正常 0封禁
	CreatedAt time.Time `json:"created_at"`
}
