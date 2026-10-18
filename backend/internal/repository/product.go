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

func (r *ProductRepo) List(categoryID uint64, sort string, page, size int) ([]model.Product, int64, error) {
    q := r.db.Model(
    if categoryID > 0 {
        q = q.Where("category_id = ?", categoryID)
    }
    var total int64
    if err := q.Count(
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
    err := q.Order(order).Offset((page - 1) * size).Limit(size).Find(
    return list, total, err
}
