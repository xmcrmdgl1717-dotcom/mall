package repository

import (
	"errors"
	"time"

	"mall/backend/internal/model"

	"gorm.io/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

type CreateOrderReq struct {
	UserID        uint64
	AddressName   string
	AddressPhone  string
	AddressDetail string
	CartIDs       []uint64
}

// CreateOrder 创建订单（核心逻辑：事务处理、扣库存、生成订单、删购物车）
func (r *OrderRepo) CreateOrder(req CreateOrderReq) (*model.Order, error) {
	var newOrder *model.Order

	err := r.db.Transaction(func(tx *gorm.DB) error {
		// 1. 查出要结算的购物车商品
		var carts []model.Cart
		if err := tx.Preload("Product").Preload("Sku").Where("id IN ? AND user_id = ?", req.CartIDs, req.UserID).Find(&carts).Error; err != nil {
			return err
		}
		if len(carts) == 0 {
			return errors.New("no items to checkout")
		}

		var totalAmount float64
		var orderItems []model.OrderItem

		// 2. 遍历商品，扣减库存，计算总价，组装订单明细
		for _, cart := range carts {
			// 使用 gorm.Expr 实现原子扣减库存，并判断库存是否充足 (stock >= quantity)
			res := tx.Model(&model.ProductSku{}).
				Where("id = ? AND stock >= ?", cart.SkuID, cart.Quantity).
				Update("stock", gorm.Expr("stock - ?", cart.Quantity))
				
			if res.RowsAffected == 0 {
				return errors.New("stock insufficient for " + cart.Product.Title)
			}

			price := cart.Sku.Price
			totalAmount += price * float64(cart.Quantity)

			orderItems = append(orderItems, model.OrderItem{
				ProductID: cart.ProductID,
				SkuID:     cart.SkuID,
				Title:     cart.Product.Title,
				Image:     cart.Product.MainImage,
				Specs:     cart.Sku.Specs,
				Price:     price,
				Quantity:  cart.Quantity,
			})
		}

		// 3. 生成订单记录
		order := &model.Order{
			OrderNo:       "ORD" + time.Now().Format("20060102150405"), // 简单生成订单号
			UserID:        req.UserID,
			TotalAmount:   totalAmount,
			Status:        1, // 假设直接支付成功，状态置为已付款
			AddressName:   req.AddressName,
			AddressPhone:  req.AddressPhone,
			AddressDetail: req.AddressDetail,
			Items:         orderItems,
		}

		if err := tx.Create(order).Error; err != nil {
			return err
		}

		// 4. 删除已结算的购物车记录
		if err := tx.Where("id IN ?", req.CartIDs).Delete(&model.Cart{}).Error; err != nil {
			return err
		}

		newOrder = order
		return nil
	})

	return newOrder, err
}

// GetUserOrders 获取用户的订单列表
func (r *OrderRepo) GetUserOrders(userID uint64) ([]model.Order, error) {
	var list []model.Order
	err := r.db.Preload("Items").Where("user_id = ?", userID).Order("created_at DESC").Find(&list).Error
	return list, err
}
