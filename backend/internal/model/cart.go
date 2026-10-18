package model

import "time"

type Cart struct {
	ID         uint64    `gorm:"primaryKey" json:"id"`
	UserID     uint64    `json:"user_id"`     // 暂时写死为 1，等以后做了用户系统再改
	ProductID  uint64    `json:"product_id"`
	SkuID      uint64    `json:"sku_id"`      // 选的规格
	Quantity   int       `json:"quantity"`    // 购买数量
	Selected   int8      `json:"selected"`    // 1 表示勾选，0 表示未勾选
	CreatedAt  time.Time `json:"created_at"`
	
	// 关联商品信息，用于前端直接展示
	Product    Product   `gorm:"foreignKey:ProductID" json:"product"`
	Sku        ProductSku `gorm:"foreignKey:SkuID" json:"sku"`
}
