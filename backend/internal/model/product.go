package model

import "time"

type Product struct {
	ID         uint64    `gorm:"primaryKey" json:"id"`
	MerchantID uint64    `json:"merchant_id"`
	CategoryID uint64    `json:"category_id"`
	Title      string    `json:"title"`
	MainImage  string    `json:"main_image"`
	MinPrice   float64   `json:"min_price"`
	MaxPrice   float64   `json:"max_price"`
	TotalSales uint64    `json:"total_sales"`
	GoodRate   float64   `json:"good_rate"`
	Status     int8      `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}
