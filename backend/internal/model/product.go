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
	IsDailyDeal int8     `json:"is_daily_deal"`
	IsHotDeal   int8     `json:"is_hot_deal"`
	CreatedAt  time.Time `json:"created_at"`
	
	// 新增：关联 SKU，用于详情页返回
	Skus []ProductSku `gorm:"foreignKey:ProductID" json:"skus"`
}

// SKU 模型
type ProductSku struct {
	ID        uint64  `gorm:"primaryKey" json:"id"`
	ProductID uint64  `json:"product_id"`
	Specs     string  `json:"specs"` // 例如："颜色:黑色;尺寸:M"
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
	SkuCode   string  `json:"sku_code"`
}
