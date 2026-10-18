package repository

import (
	"mall/backend/internal/model"
	"gorm.io/gorm"
)

type CategoryRepo struct { db *gorm.DB }
func NewCategoryRepo(db *gorm.DB) *CategoryRepo { return &CategoryRepo{db: db} }

func (r *CategoryRepo) ListAll() ([]model.Category, error) {
	var list []model.Category
	err := r.db.Order("sort ASC").Find(&list).Error
	return list, err
}

// 新增：创建类目
func (r *CategoryRepo) Create(c *model.Category) error {
	return r.db.Create(c).Error
}

// 新增：更新类目
func (r *CategoryRepo) Update(c *model.Category) error {
	return r.db.Save(c).Error
}

// 新增：删除类目
func (r *CategoryRepo) Delete(id uint64) error {
	return r.db.Delete(&model.Category{}, id).Error
}

// Banner 相关
type BannerRepo struct { db *gorm.DB }
func NewBannerRepo(db *gorm.DB) *BannerRepo { return &BannerRepo{db: db} }

func (r *BannerRepo) ListActive() ([]model.Banner, error) {
	var list []model.Banner
	err := r.db.Where("status = ?", 1).Order("sort ASC").Find(&list).Error
	return list, err
}
// 完善 Banner 数据库操作
func (r *BannerRepo) ListAll() ([]model.Banner, error) {
	var list []model.Banner
	err := r.db.Order("sort ASC").Find(&list).Error
	return list, err
}

func (r *BannerRepo) Create(b *model.Banner) error {
	return r.db.Create(b).Error
}

func (r *BannerRepo) Delete(id uint64) error {
	return r.db.Delete(&model.Banner{}, id).Error
}
