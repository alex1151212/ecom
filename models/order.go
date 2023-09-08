package models

import (
	"ecom/utils"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OwnerID        uint
	IsUsedCoupon   bool
	BillStatus     string // pending
	OrderStatus    string //confirm
	PayStatus      string
	OrderTime      *time.Time
	OrderNumber    string
	OrderAmount    int
	OriginAmount   int
	DiscountAmount int
	OrderProducts  []OrderProduct `gorm:"foreignkey:OrderID;"`

	/* 發票 */
	// invoicesStatus
	// "pending"
	// invoiceType
	// "member"

}

type OrderProduct struct {
	gorm.Model
	Product        Product `gorm:"foreignkey:ProductID;"`
	ProductSubcode uint
	Quantity       int
	Price          int
	Status         string

	ProductID      uint
	ShoppingCartID uint
	OrderID        uint
}

/* OrderProduct */
func CreateOrderProduct(product OrderProduct) *gorm.DB {
	return utils.DB.Create(&product)
}

func DeleteOrderProduct(product OrderProduct) *gorm.DB {
	return utils.DB.Delete(&product)
}

func UpdateOrderProduct(product OrderProduct) *gorm.DB {
	utils.DB.Model(&product).Updates(OrderProduct{
		ProductSubcode: product.ProductSubcode,
		Quantity:       product.Quantity,
		Price:          product.Price,
		Status:         product.Status,
		ProductID:      product.ProductID,
		ShoppingCartID: product.ShoppingCartID,
		OrderID:        product.OrderID,
	})

	// utils.DB.Model(&product).Association("Variation").Replace(product.Variation)

	return utils.DB.Save(&product)
}

func FindOrderProduct(orderProductId uint) {

}

/* Order */

func CreateOrder(order Order) *gorm.DB {
	return utils.DB.Create(&order)
}

func DeleteOrder(order Order) *gorm.DB {
	return utils.DB.Delete(&order)
}

func UpdateOrder(order Order) *gorm.DB {
	utils.DB.Model(&order).Updates(Order{
		BillStatus:  order.BillStatus,
		OrderStatus: order.OrderStatus,
		PayStatus:   order.PayStatus,
	})

	// utils.DB.Model(&product).Association("Variation").Replace(product.Variation)

	return utils.DB.Save(&order)
}
