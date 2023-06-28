package service

import (
	"ecom/utils"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"ecom/models"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

//	 CreateUser
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
	fmt.Println(user.Username, "  >>>>>>>>>>>  ", password, repassword)

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
	user.Username = c.PostForm("name")
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

//	 Login
//	 @Summary 所有用戶
//		@Tags		用戶
//		@Param username formData string false		"使用者名稱"
//		@Param password formData string false		"使用者密碼"
//		@Success	200	{string}	json{"code","message"}
//		@Router		/user/login [post]
func LoginUser(c *gin.Context) {

	data := models.User{}
	name := c.PostForm("username")
	password := c.PostForm("password")

	user := models.FindUserByName(name)
	if user.Username == "" {
		utils.RespFail(c.Writer, "該使用者不存在")
		return
	}
	flag := utils.ValidPassword(password, user.Salt, user.Password)
	if !flag {
		utils.RespFail(c.Writer, "密碼不正確")
		return
	}
	enCodePwd := utils.MakePasssword(password, user.Salt)
	data = models.FindUserByUsernameAndPwd(name, enCodePwd)

	utils.RespOK(c.Writer, data, "登入成功")
}
