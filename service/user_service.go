package service

import (
	"ecom/utils"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"ecom/models"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//	 Register
//	 @Summary 新增用戶
//		@Tags		用戶
//		@Param username formData string false		"使用者名稱"
//		@Param password formData string false		"使用者密碼"
//		@Param repassword formData string false		"使用者確認密碼"
//		@Success	200	{string}	json{"code","message"}
//		@Router		/user/createUser [post]
func CreateUser(c *gin.Context) {
	user := models.User{}

	user.Username = c.PostForm("username")
	password := c.PostForm("password")
	repassword := c.PostForm("repassword")

	salt := fmt.Sprintf("%06d", rand.Int31())

	data := models.FindUserByName(user.Username)
	if user.Username == "" || password == "" || repassword == "" {

		utils.RespFail(c.Writer, "使用者名稱或密碼不能為空")
		return
	}
	if data.Username != "" {
		utils.RespFail(c.Writer, "使用者名稱已註冊")
		return
	}
	if password != repassword {
		utils.RespFail(c.Writer, "兩次密碼不一致")
		return
	}

	user.Password = utils.MakePasssword(password, salt)
	user.Salt = salt
	models.CreateUser(user)

	utils.RespOK(c.Writer, data, "創建用戶成功")
}

//	 DeleteUser
//	 @Summary 刪除用戶
//		@Tags		用戶
//		@Param id formData string false		"使用者ID"
//		@Success	200	{string}	json{"code","message"}
//		@Router		/user/deleteUser [post]
func DeleteUser(c *gin.Context) {
	user := models.User{}

	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)

	models.DeleteUser(user)
	utils.RespOK(c.Writer, id, "刪除用戶成功")
}

//	 UpdateUser
//	 @Summary 編輯用戶
//		@Tags		用戶
//		@Param id formData string false		"使用者ID"
//		@Param name formData string false		"使用者名稱"
//		@Param password formData string false		"使用者密碼"
//		@Param phone formData string false		"電話號碼"
//		@Param email formData string false		"電子信箱"
//		@Success	200	{string}	json{"code","message"}
//		@Router		/user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.User{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code":    -1, //0成功 -1失敗
			"message": "編輯用戶失敗",
			"data":    user,
		})
		return
	}
	models.UpdateUser(user)
	utils.RespOK(c.Writer, user, "編輯用戶成功")

}

//	 LoginUser
//	 @Summary 用戶登入
//		@Tags		用戶
//		@Param username formData string false		"使用者名稱"
//		@Param password formData string false		"使用者密碼"
//		@Success	200	{string}	json "{"code","message"}"
//		@Router		/login [post]
func LoginUser(c *gin.Context) {
	Auth().LoginHandler(c)
}

// LogoutUser
// @Summary 用戶登出
// @Security BearerAuth
// @Tags		用戶
// @Success	200	{string}	json "{"code","message"}"
// @Router		/auth/logout [post]
func LogoutUser(c *gin.Context) {
	Auth().LogoutHandler(c)
}

// Auth Test
// @Summary 驗證功能測試路由
// @Security BearerAuth
// @Tags		用戶
// @Success	200	{string}	json "{"code","message"}"
// @Router		/auth/hello [get]
func HelloHandler(c *gin.Context) {
	identityKey := viper.GetString("jwt.identityKey")
	fmt.Println(">>>>>>>>>>>>>>", identityKey)
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(viper.GetString("jwt.identityKey"))
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*models.User).Username,
		"text":     "Hello World.",
	})
}

func Auth() *jwt.GinJWTMiddleware {

	payloadFunc := func(data interface{}) jwt.MapClaims {
		if v, ok := data.(*models.User); ok {

			return jwt.MapClaims{
				viper.GetString("identityKey"): v.Username,
			}

		}
		return jwt.MapClaims{}
	}

	identityHandler := func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		return &models.User{
			Username: claims[viper.GetString("identityKey")].(string),
		}
	}

	authenticator := func(c *gin.Context) (interface{}, error) {

		data := models.User{}

		name, isNameEmpty := c.GetPostForm("username")
		password, isPassword := c.GetPostForm("password")
		if !isNameEmpty || !isPassword {
			return "", jwt.ErrMissingLoginValues
		}

		user := models.FindUserByName(name)
		if user.Username == "" {
			return "", errors.New("user doesn't exist")
		}
		flag := utils.ValidPassword(password, user.Salt, user.Password)
		if !flag {
			return "", jwt.ErrMissingLoginValues
		}
		enCodePwd := utils.MakePasssword(password, user.Salt)
		data = models.FindUserByUsernameAndPwd(name, enCodePwd)
		return &models.User{
			Username: data.Username,
		}, nil

	}

	authorizator := func(data interface{}, c *gin.Context) bool {
		// TODO 驗證方式重寫
		if v, ok := data.(*models.User); ok && v.Username == "admin" {
			return true
		}

		return false
	}

	return utils.AuthMiddleware(payloadFunc, identityHandler, authenticator, authorizator)
}
