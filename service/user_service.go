package service

import (
	"ecom/utils"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"

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
//		@Param user body models.UserInfo true " "
//		@Success	200	{string}	json{"code","message"}
//		@Router		/user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.User{}

	err := c.BindJSON(&user)

	if err != nil {

		utils.RespFail(c.Writer, "編輯用戶失敗")
		return
	}

	_, err = govalidator.ValidateStruct(user)
	if err != nil {

		utils.RespFail(c.Writer, "編輯用戶失敗")
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

// AddFavouriteProduct
// @Summary 添加商品到我的最愛
// @Security BearerAuth
// @Tags		用戶
// @Param productId formData string true " "
// @Success	200	{string}	json{"code","message"}
// @Router		/auth/addFavouriteProduct [post]
func AddFavouriteProduct(c *gin.Context) {
	identityKey := viper.GetString("jwt.identityKey")
	user, _ := c.Get(identityKey)

	productId := c.PostForm("productId")
	intProductId, err := strconv.Atoi(productId)

	if err != nil {

		utils.RespFail(c.Writer, "編輯用戶失敗")
		return
	}

	_, err = govalidator.ValidateStruct(user)
	if err != nil {

		utils.RespFail(c.Writer, "編輯用戶失敗")
		return
	}

	models.AddFavouriteProduct(user.(*models.User).Username, uint(intProductId))

	utils.RespOK(c.Writer, user, "編輯用戶成功")
}

// GetFavouriteProduct
// @Summary 添加商品到我的最愛
// @Security BearerAuth
// @Tags		用戶
// @Success	200	{string}	json{"code","message"}
// @Router		/auth/getFavouriteProduct [post]
func GetFavouriteProduct(c *gin.Context) {
	identityKey := viper.GetString("jwt.identityKey")
	user, _ := c.Get(identityKey)

	product := models.FindUserFavouriteProduct(user.(*models.User).Username)

	utils.RespOK(c.Writer, product, "編輯用戶成功")
}

// Auth Test
// @Summary 驗證功能測試路由
// @Security BearerAuth
// @Tags		用戶
// @Success	200	{string}	json "{"code","message"}"
// @Router		/auth/hello [get]
func HelloHandler(c *gin.Context) {
	identityKey := viper.GetString("jwt.identityKey")
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*models.User).Username,
		"text":     "Hello World.",
	})
}

// UploadUserAvatar
// @Summary 用戶照片上傳
// @Tags		用戶
// @Param userId formData int false		"使用這ID"
// @Param file formData file false		"1:1圖片"
// @Success	200	{string}	json "{"code","message"}"
// @Router		/user/uploadUserAvatar [post]
func UploadUserAvatar(c *gin.Context) {
	w := c.Writer
	req := c.Request
	userId := c.PostForm("userId")

	if userId == "" {
		utils.RespFail(w, "無使用者")
		return
	}

	srcFile, head, err := req.FormFile("file")
	if err != nil {
		utils.RespFail(w, err.Error())
		return
	}

	suffix := ".png"
	oFileName := head.Filename
	tem := strings.Split(oFileName, ".")
	if len(tem) > 1 {
		suffix = "." + tem[len(tem)-1]
	}

	fileName := fmt.Sprintf("%s%s", userId, suffix)
	dstFile, err := os.Create("./assets/user_avatar/" + fileName)
	if err != nil {
		utils.RespFail(w, err.Error())
		return
	}
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		utils.RespFail(w, err.Error())
		return
	}
	url := "./asset/user_avatar/" + fileName

	user := models.User{}
	id, _ := strconv.Atoi(userId)
	user.ID = uint(id)
	user.AvatarURL = url
	models.UpdateUser(user)

	utils.RespOK(w, url, "發送圖片成功")
}

func Auth() *jwt.GinJWTMiddleware {

	payloadFunc := func(data interface{}) jwt.MapClaims {
		if v, ok := data.(*models.User); ok {

			return jwt.MapClaims{
				viper.GetString("jwt.identityKey"): v.Username,
			}

		}
		return jwt.MapClaims{}
	}

	identityHandler := func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		return &models.User{
			Username: claims[viper.GetString("jwt.identityKey")].(string),
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
		if v, ok := data.(*models.User); ok && v.Username != "" {
			return true
		}
		return false
	}

	return utils.AuthMiddleware(payloadFunc, identityHandler, authenticator, authorizator)
}
