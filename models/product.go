package models

import (
	"ecom/utils"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Variation   []ProductVariation `gorm:"foreignkey:ProductID;"`
	Description string
	Sold        uint
	Liked       uint
	ProductImg  []string `gorm:"serializer:json"`
}

type ProductVariation struct {
	gorm.Model
	Stock          uint
	Price          int
	Specifications map[string]string `gorm:"serializer:json"`
	ProductID      uint
}

type CreateProductType struct {
	Name        string                       `json:"name"`
	Variation   []CreateProductVariationType `json:"variation"  example:CreateProductVariationType`
	Description string                       `json:"description"`
	ProductImg  []string                     `json:"productImg"`
}
type CreateProductVariationType struct {
	Stock          uint              `json:"stock"`
	Price          int               `json:"price"`
	Specifications map[string]string `json:"specifications"`
}
type UpdateProductType struct {
	ID          uint                         `json:"id"`
	Name        string                       `json:"name"`
	Variation   []UpdateProductVariationType `json:"variation"  example:UpdateProductVariationType`
	Description string                       `json:"description"`
	ProductImg  []string                     `json:"productImg"`
}
type UpdateProductVariationType struct {
	Stock          uint              `json:"stock"`
	Price          int               `json:"price"`
	Specifications map[string]string `json:"specifications"`
}

func CreateProduct(product Product) *gorm.DB {
	return utils.DB.Create(&product)
}

func DeleteProduct(product Product) *gorm.DB {
	return utils.DB.Delete(&product)
}

func UpdateProduct(product Product) *gorm.DB {

	// utils.DB.Model(&product).Preload("Variation").Omit().Updates(Product{
	// 	Name:        product.Name,
	// 	Variation:   product.Variation,
	// 	Description: product.Description,
	// 	ProductImg:  product.ProductImg,
	// })
	return utils.DB.Preload("Variation").Model(&product).Omit().Updates(Product{
		Name:        product.Name,
		Variation:   product.Variation,
		Description: product.Description,
		ProductImg:  product.ProductImg,
	})

}
func FindProduct(productId uint) Product {
	product := Product{}

	utils.DB.Preload("Variation").Where("Id = ?", productId).First(&product)

	return product
}
