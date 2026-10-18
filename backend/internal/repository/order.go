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

// CreateOrder 创建订单
func (r *OrderRepo) CreateOrder(req CreateOrderReq) (*model.Order, error) {
	var newOrder *model.Order

	err := r.db.Transaction(func(tx *gorm.DB) error {
		var carts []model.Cart
		if err := tx.Preload("Product").Preload("Sku").Where("id IN ? AND user_id = ?", req.CartIDs, req.UserID).Find(&carts).Error; err != nil {
			return err
		}
		if len(carts) == 0 {
			return errors.New("no items to checkout")
		}

		var totalAmount float64
		var orderItems []model.OrderItem

		for _, cart := range carts {
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

		order := &model.Order{
			OrderNo:       "ORD" + time.Now().Format("20060102150405"),
			UserID:        req.UserID,
			TotalAmount:   totalAmount,
			Status:        1, // 1 代表待发货
			AddressName:   req.AddressName,
			AddressPhone:  req.AddressPhone,
			AddressDetail: req.AddressDetail,
			Items:         orderItems,
		}

		if err := tx.Create(order).Error; err != nil {
			return err
		}

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

// ================= 新增：更新订单状态 =================
func (r *OrderRepo) UpdateStatus(orderID uint64, status int8) error {
	return r.db.Model(&model.Order{}).Where("id = ?", orderID).Update("status", status).Error
}
