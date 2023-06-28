package models

import (
	"ecom/utils"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// 基本使用者資訊
type User struct {
	gorm.Model
	Username     string
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
	ShippingInfo []ShippingInfo `gorm:"foreignkey:UserID;"`

	// 購物車
	ShoppingCart ShoppingCart `gorm:"foreignkey:OwnerID;"`
	// 我的最愛
	Favourite []uint `gorm:"type:json;"` // ProductID
}

// 購物車
type ShoppingCart struct {
	gorm.Model
	ProductList []OrderProduct `gorm:"foreignkey:ShoppingCartID;"`
	OwnerID     uint
}

// 收貨資訊
type ShippingInfo struct {
	gorm.Model
	Phone           string `valid:"matches(^09\\d{8}$)"`
	Consignee       string
	ShippingAddress string
	UserID          uint
}

func CreateUser(user User) *gorm.DB {
	return utils.DB.Create(&user)
}

func DeleteUser(user User) *gorm.DB {
	return utils.DB.Delete(&user)
}

func FindAllUser() {

}

func FindUserByName(name string) User {
	user := User{}
	utils.DB.Where("username = ?", name).First(&user)

	// Encode token
	str := fmt.Sprintf("%d", time.Now().Unix())
	token := utils.Md5Encode(str)

	utils.DB.Model(&user).Where("Id = ? ", user.ID).Update("identity", token)
	return user
}

func UpdateUser(user User) *gorm.DB {
	return utils.DB.Model(&user).Updates(User{
		Username: user.Username,
		Password: user.Password,
		Phone:    user.Phone,
		Email:    user.Email,
		Sex:      user.Sex,
		// Birthday:	user.Birthday,
	})
}

func FindUserByUsernameAndPwd(name string, password string) User {
	user := User{}
	utils.DB.Where("username = ? and password = ? ", name, password).First(&user)
	return user
}
