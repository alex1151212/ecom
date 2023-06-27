package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Variation   []ProductVariation `gorm:"foreignkey:ProductID"`
	Description string
	Sold        uint
	Liked       uint
	ProductImg  []string
}

type ProductVariation struct {
	gorm.Model
	// ProductSubcode uint
	Stock          uint
	Price          int
	Specifications map[string]string
	ProductID      uint
}
