package models

import (
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

	ProductID      uint
	ShoppingCartID uint
	OrderID        uint
}
