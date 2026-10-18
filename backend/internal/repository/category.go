package repository

import (
	"mall/backend/internal/model"
	"gorm.io/gorm"
)

type CategoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{db: db}
}

func (r *CategoryRepo) ListAll() ([]model.Category, error) {
	var list []model.Category
	err := r.db.Order("sort ASC").Find(&list).Error
	return list, err
}
