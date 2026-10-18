package repository

import (
	"mall/backend/internal/model"
	"gorm.io/gorm"
)

type MerchantRepo struct { db *gorm.DB }
func NewMerchantRepo(db *gorm.DB) *MerchantRepo { return &MerchantRepo{db: db} }

// GetTopMerchants 获取热门店铺
func (r *MerchantRepo) GetTopMerchants(limit int) ([]model.Merchant, error) {
	var list []model.Merchant
	err := r.db.Where("status = ?", 1).Order("total_sales DESC").Limit(limit).Find(&list).Error
	return list, err
}
