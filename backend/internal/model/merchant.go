package model

import "time"

type Merchant struct {
	ID            uint64    `gorm:"primaryKey" json:"id"`
	Name          string    `json:"name"`
	Logo          string    `json:"logo"`
	Description   string    `json:"description"`
	TotalSales    uint64    `json:"total_sales"`
	GoodRate      float64   `json:"good_rate"`
	Status        int8      `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}
