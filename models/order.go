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
	OrderTime      *time.Time
	OrderNumber    string
	OrderAmount    int
	OriginAmount   int
	DiscountAmount int

	// payStatus
	// "outstanding"
	// invoiceStatus
	// "pending"
	// invoiceType
	// "member"

}
