package repository

import (
	"errors"
	"mall/backend/internal/model"
	"time"

	"gorm.io/gorm"
)

type CouponRepo struct { db *gorm.DB }
func NewCouponRepo(db *gorm.DB) *CouponRepo { return &CouponRepo{db: db} }

// ListActive 获取当前可领取的优惠券
func (r *CouponRepo) ListActive() ([]model.Coupon, error) {
	var list []model.Coupon
	now := time.Now()
	err := r.db.Where("left_count > 0 AND start_time <= ? AND end_time >= ?", now, now).Find(&list).Error
	return list, err
}

// ListAll 后台获取所有优惠券
func (r *CouponRepo) ListAll() ([]model.Coupon, error) {
	var list []model.Coupon
	err := r.db.Order("created_at DESC").Find(&list).Error
	return list, err
}

// Create 后台新增优惠券
func (r *CouponRepo) Create(c *model.Coupon) error {
	c.LeftCount = c.TotalCount
	return r.db.Create(c).Error
}

// Claim 用户领券
func (r *CouponRepo) Claim(userID, couponID uint64) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var coupon model.Coupon
		if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&coupon, couponID).Error; err != nil {
			return err
		}
		if coupon.LeftCount <= 0 {
			return errors.New("优惠券已被领完")
		}

		// 检查是否已经领过
		var count int64
		tx.Model(&model.UserCoupon{}).Where("user_id = ? AND coupon_id = ?", userID, couponID).Count(&count)
		if count > 0 {
			return errors.New("您已领取过该优惠券")
		}

		// 扣减库存
		if err := tx.Model(&coupon).Update("left_count", gorm.Expr("left_count - 1")).Error; err != nil {
			return err
		}

		// 记录到用户券包
		return tx.Create(&model.UserCoupon{
			UserID:   userID,
			CouponID: couponID,
			Status:   0,
		}).Error
	})
}

// GetUserCoupons 获取用户的券包
func (r *CouponRepo) GetUserCoupons(userID uint64) ([]model.UserCoupon, error) {
	var list []model.UserCoupon
	err := r.db.Preload("Coupon").Where("user_id = ?", userID).Order("created_at DESC").Find(&list).Error
	return list, err
}
