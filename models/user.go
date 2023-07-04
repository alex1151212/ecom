package models

import (
	"ecom/utils"
	"time"

	"gorm.io/gorm"
)

// 基本使用者資訊
type User struct {
	gorm.Model
	// TODO改成唯一值
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
	Favourite []uint `gorm:"serializer:json;"` // ProductID
}

type UserInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone";valid:"matches(^09\\d{8}$)"`
	Email    string `json:"email";valid:"email"`
	// Birthday  *time.Time
	Sex       int    `json:"sex"`
	AvatarURL string `json:"avatarUrl"`
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

func FindUserByName(username string) User {
	user := User{}
	utils.DB.Where("username = ?", username).First(&user)

	return user
}
func FindUserById(userId uint) User {
	user := User{}
	utils.DB.Where("Id = ?", userId).First(&user)

	return user
}

func FindUserFavouriteProduct(username string) []Product {
	product := []Product{}

	userFavouriteProduct := FindUserByName(username).Favourite
	utils.DB.Find(&product, userFavouriteProduct)
	return product
}

func UpdateUser(user User) *gorm.DB {
	return utils.DB.Model(&user).Where("Id = ?", user.ID).Omit().Updates(User{
		Username:  user.Username,
		Password:  user.Password,
		Phone:     user.Phone,
		Email:     user.Email,
		Sex:       user.Sex,
		AvatarURL: user.AvatarURL,
		// Birthday:	user.Birthday,
	})
}

func AddFavouriteProduct(username string, productId uint) *gorm.DB {
	user := User{}
	utils.DB.Where("username = ?", username).First(&user)
	_, found := utils.SliceFind(user.Favourite, productId)
	if found {
		user.Favourite = utils.SliceRemove(user.Favourite, productId)
		return utils.DB.Save(&user)
	}

	user.Favourite = append(user.Favourite, productId)
	return utils.DB.Save(&user)
}

func FindUserByUsernameAndPwd(name string, password string) User {
	user := User{}
	utils.DB.Where("username = ? and password = ? ", name, password).First(&user)
	return user
}
func UpdateUserAvatar(username string) {

}
