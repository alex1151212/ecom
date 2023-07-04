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
type ProductInfo struct {
	CreateProductType
	ID uint `json:"id"`
}

func CreateProduct(product Product) *gorm.DB {
	return utils.DB.Create(&product)
}

func DeleteProduct(product Product) *gorm.DB {
	return utils.DB.Delete(&product)
}

func UpdateProduct(product Product) *gorm.DB {

	utils.DB.Model(&product).Updates(Product{
		Name:        product.Name,
		Description: product.Description,
		ProductImg:  product.ProductImg,
	})

	utils.DB.Model(&product).Association("Variation").Replace(product.Variation)

	return utils.DB.Save(&product)
}

func FindProduct(productId uint) *Product {
	// TODO 只輸出指定欄位
	product := Product{}

	state := utils.DB.Preload("Variation").Model(&product).Where("Id = ?", productId).First(&product)

	if state.Error != nil {
		return nil
	}

	return &product
}
