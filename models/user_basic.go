package models

import (
	"time"

	"gorm.io/gorm"
)

// 基本使用者資訊
type UserBasic struct {
	gorm.Model
	UserName     string
	Password     string
	Phone        string `valid:"matches(^09\\d{8}$)"`
	Email        string `valid:"email"`
	Identity     string
	Salt         string
	LoginTime    *time.Time
	LogoutTime   *time.Time
	Birthday     *time.Time
	Sex          int
	AvatarURL    string
	IsLogout     bool
	ShippingInfo []ShippingInfo `gorm:"foreignkey:UserID"`

	// 購物車
	ShoppingCart ShoppingCart `gorm:"foreignkey:UserID"`
	// 我的最愛
	// 訂購紀錄
}

// 購物車
type ShoppingCart struct {
	gorm.Model
	// productList []
	UserID uint
}

// 我的最愛
type Favourite struct {
	// many2many
}

// 收貨資訊
type ShippingInfo struct {
	gorm.Model
	Phone           string `valid:"matches(^09\\d{8}$)"`
	Consignee       string
	ShippingAddress string
	UserID          uint
}
