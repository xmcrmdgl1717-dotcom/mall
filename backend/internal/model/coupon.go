package model

import "time"

// Coupon 优惠券模板 (后台发布的)
type Coupon struct {
	ID         uint64    `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"size:100" json:"name"`        // 优惠券名称
	Amount     float64   `json:"amount"`                      // 抵扣金额
	MinAmount  float64   `json:"min_amount"`                  // 满多少可用 (0为无门槛)
	TotalCount int       `json:"total_count"`                 // 发行总量
	LeftCount  int       `json:"left_count"`                  // 剩余数量
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	CreatedAt  time.Time `json:"created_at"`
}

// UserCoupon 用户领取的优惠券记录
type UserCoupon struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	UserID    uint64    `json:"user_id"`
	CouponID  uint64    `json:"coupon_id"`
	Status    int8      `json:"status"` // 0未使用 1已使用 2已过期
	CreatedAt time.Time `json:"created_at"`
	
	Coupon    Coupon    `gorm:"foreignKey:CouponID" json:"coupon"` // 关联模板信息
}
