package repository

import (
	"mall/backend/internal/model"
	"gorm.io/gorm"
)

type AddressRepo struct { db *gorm.DB }
func NewAddressRepo(db *gorm.DB) *AddressRepo { return &AddressRepo{db: db} }

// List 获取用户的地址列表
func (r *AddressRepo) List(userID uint64) ([]model.Address, error) {
	var list []model.Address
	err := r.db.Where("user_id = ?", userID).Order("is_default DESC, created_at DESC").Find(&list).Error
	return list, err
}

// Create 新增地址
func (r *AddressRepo) Create(addr *model.Address) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 如果设为了默认地址，先把该用户的其他地址取消默认
		if addr.IsDefault == 1 {
			tx.Model(&model.Address{}).Where("user_id = ?", addr.UserID).Update("is_default", 0)
		}
		return tx.Create(addr).Error
	})
}

// Delete 删除地址
func (r *AddressRepo) Delete(id uint64, userID uint64) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&model.Address{}).Error
}
