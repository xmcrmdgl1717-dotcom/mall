package repository

import (
	"mall/backend/internal/model"
	"gorm.io/gorm"
)

type CartRepo struct { db *gorm.DB }
func NewCartRepo(db *gorm.DB) *CartRepo { return &CartRepo{db: db} }

// AddToCart 加入购物车（如果已经存在则增加数量）
func (r *CartRepo) AddToCart(cart *model.Cart) error {
	var existing model.Cart
	err := r.db.Where("user_id = ? AND sku_id = ?", cart.UserID, cart.SkuID).First(&existing).Error
	if err == gorm.ErrRecordNotFound {
		// 没找到，直接新建
		return r.db.Create(cart).Error
	}
	// 找到了，数量累加
	existing.Quantity += cart.Quantity
	return r.db.Save(&existing).Error
}

// GetCartList 获取购物车列表（带上商品和 SKU 信息）
func (r *CartRepo) GetCartList(userID uint64) ([]model.Cart, error) {
	var list []model.Cart
	err := r.db.Preload("Product").Preload("Sku").Where("user_id = ?", userID).Find(&list).Error
	return list, err
}

// DeleteCartItem 删除购物车商品
func (r *CartRepo) DeleteCartItem(id uint64, userID uint64) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&model.Cart{}).Error
}
