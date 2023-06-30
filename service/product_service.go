package service

import (
	"ecom/models"
	"ecom/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

//	 CreateProduct
//	 @Summary 新增商品
//		@Tags		商品
//		@Param product body models.CreateProductType true " "
//		@Success	200	{string}	json " "
//		@Router		/product/createProduct [post]
func CreateProduct(c *gin.Context) {
	product := models.Product{}

	err := c.BindJSON(&product)

	if err != nil {
		utils.RespFail(c.Writer, "輸入內容有誤")
	}

	models.CreateProduct(product)

	utils.RespOK(c.Writer, product, "產品新增成功")
}

//	 GetProduct
//	 @Summary 取得產品資訊
//		@Tags		商品
//		@Param id query uint true " "
//		@Success	200	{string}	json " "
//		@Router		/product/getProduct [get]
func GetProduct(c *gin.Context) {
	id := c.Query("id")
	uintId, _ := strconv.ParseUint(id, 10, 32)
	v := models.FindProduct(uint(uintId))

	utils.RespOK(c.Writer, v, "測試成功")
}

//	 UpdateProduct
//	 @Summary 更新產品資訊
//		@Tags		商品
//		@Param product body models.UpdateProductType true " "
//		@Success	200	{string}	json " "
//		@Router		/product/updateProduct [post]
func UpdateProduct(c *gin.Context) {
	product := models.Product{}

	err := c.BindJSON(&product)

	if err != nil {
		utils.RespFail(c.Writer, "輸入內容有誤")
	}

	// v := models.FindProduct(product.ID)

	models.UpdateProduct(product)

	utils.RespOK(c.Writer, product, "測試成功")
}
