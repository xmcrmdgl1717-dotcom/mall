package repository

import (
	"mall/backend/internal/model"

	"gorm.io/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

// List 接收 5 个参数，包含 dealType 用来区分特惠
func (r *ProductRepo) List(categoryID uint64, sort string, page, size int, dealType string) ([]model.Product, int64, error) {
	q := r.db.Model(&model.Product{}).Where("status = ?", 1)
	
	if categoryID > 0 {
		q = q.Where("category_id = ?", categoryID)
	}

	// 新增的判断：如果是特惠商品
	if dealType == "daily" {
		q = q.Where("is_daily_deal = 1")
	} else if dealType == "hot" {
		q = q.Where("is_hot_deal = 1")
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	order := "created_at DESC"
	switch sort {
	case "sales":
		order = "total_sales DESC"
	case "price_asc":
		order = "min_price ASC"
	case "price_desc":
		order = "min_price DESC"
	}

	var list []model.Product
	err := q.Order(order).Offset((page - 1) * size).Limit(size).Find(&list).Error
	return list, total, err
}

// GetByID 获取单个商品详情（带 SKU）
func (r *ProductRepo) GetByID(id uint64) (*model.Product, error) {
	var product model.Product
	err := r.db.Preload("Skus").First(&product, id).Error
	return &product, err
}

func (r *ProductRepo) Create(product *model.Product) error {
	return r.db.Create(product).Error
}
