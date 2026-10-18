package model

import "time"

type Order struct {
	ID            uint64      `gorm:"primaryKey" json:"id"`
	OrderNo       string      `gorm:"uniqueIndex;size:32" json:"order_no"`
	UserID        uint64      `json:"user_id"`
	TotalAmount   float64     `json:"total_amount"`
	Status        int8        `json:"status"` // 0待付款, 1已付款, 2已发货, 3已完成
	AddressName   string      `json:"address_name"`
	AddressPhone  string      `json:"address_phone"`
	AddressDetail string      `json:"address_detail"`
	CreatedAt     time.Time   `json:"created_at"`
	
	Items         []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
}

type OrderItem struct {
	ID        uint64  `gorm:"primaryKey" json:"id"`
	OrderID   uint64  `json:"order_id"`
	ProductID uint64  `json:"product_id"`
	SkuID     uint64  `json:"sku_id"`
	Title     string  `json:"title"`
	Image     string  `json:"image"`
	Specs     string  `json:"specs"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
}
