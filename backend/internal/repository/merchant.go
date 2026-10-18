package repository

import (
	"mall/backend/internal/model"
	"gorm.io/gorm"
)

type MerchantRepo struct { db *gorm.DB }
func NewMerchantRepo(db *gorm.DB) *MerchantRepo { return &MerchantRepo{db: db} }

// GetTopMerchants 首页热门店铺
func (r *MerchantRepo) GetTopMerchants(limit int) ([]model.Merchant, error) {
	var list []model.Merchant
	err := r.db.Where("status = ?", 1).Order("total_sales DESC").Limit(limit).Find(&list).Error
	return list, err
}

// ListAll 后台获取所有商家
func (r *MerchantRepo) ListAll() ([]model.Merchant, error) {
	var list []model.Merchant
	err := r.db.Order("created_at DESC").Find(&list).Error
	return list, err
}

// UpdateStatus 更改商家状态 (0待审核 1已通过 2已驳回 3已冻结)
func (r *MerchantRepo) UpdateStatus(id uint64, status int8) error {
	return r.db.Model(&model.Merchant{}).Where("id = ?", id).Update("status", status).Error
}
