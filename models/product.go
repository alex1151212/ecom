package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Variation   []ProductVariation `gorm:"foreignkey:ProductID;"`
	Description string
	Sold        uint
	Liked       uint
	ProductImg  []string `gorm:"type:json"`
}

type ProductVariation struct {
	gorm.Model
	ProductSubcode uint
	Stock          uint
	Price          int
	Specifications map[string]string `gorm:"type:json"`
	ProductID      uint
}

type CreateProduct struct {
	Name        string                   `json:"name"`
	Variation   []CreateProductVariation `json:"variation"  example:CreateProductVariation`
	Description string                   `json:"description"`
	ProductImg  []string                 `json:"productImg"`
}
type CreateProductVariation struct {
	ProductSubcode uint `json:"productSubcode"`
	Stock          uint `json:"stock"`
	Price          int  `json:"price"`
	ProductID      uint `json:"productID"`
}
