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

// List 接收 6 个参数，包含 keyword 模糊搜索
func (r *ProductRepo) List(categoryID uint64, sort string, page, size int, dealType string, keyword string) ([]model.Product, int64, error) {
	q := r.db.Model(&model.Product{}).Where("status = ?", 1)
	
	if categoryID > 0 {
		q = q.Where("category_id = ?", categoryID)
	}

	// 关键词搜索 (Postgres 下使用 ILIKE 忽略大小写)
	if keyword != "" {
		q = q.Where("title ILIKE ?", "%"+keyword+"%")
	}

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

// Create 支持带有 SKU 的商品创建（使用数据库事务）
func (r *ProductRepo) Create(product *model.Product) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if len(product.Skus) > 0 {
			minPrice := product.Skus[0].Price
			maxPrice := product.Skus[0].Price
			for _, sku := range product.Skus {
				if sku.Price < minPrice { minPrice = sku.Price }
				if sku.Price > maxPrice { maxPrice = sku.Price }
			}
			product.MinPrice = minPrice
			product.MaxPrice = maxPrice
		}

		if err := tx.Create(product).Error; err != nil {
			return err
		}
		return nil
	})
}
