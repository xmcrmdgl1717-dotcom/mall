package repository

import (
	"mall/backend/internal/model"
	"gorm.io/gorm"
)

type DashboardRepo struct { db *gorm.DB }
func NewDashboardRepo(db *gorm.DB) *DashboardRepo { return &DashboardRepo{db: db} }

func (r *DashboardRepo) GetStats() (map[string]interface{}, error) {
	var userCount, orderCount, productCount int64
	var totalSales float64

	r.db.Model(&model.User{}).Count(&userCount)
	r.db.Model(&model.Order{}).Count(&orderCount)
	r.db.Model(&model.Product{}).Count(&productCount)
	
	// 统计已付款/已发货/已完成的订单总金额
	r.db.Model(&model.Order{}).Where("status >= ?", 1).Select("COALESCE(SUM(total_amount), 0)").Scan(&totalSales)

	return map[string]interface{}{
		"user_count":    userCount,
		"order_count":   orderCount,
		"product_count": productCount,
		"total_sales":   totalSales,
	}, nil
}
