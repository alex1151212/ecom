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

//	 UpdateProduct
//	 @Summary 更新產品資訊
//		@Tags		商品
//		@Param product body models.ProductInfo true " "
//		@Success	200	{string}	json " "
//		@Router		/product/updateProduct [post]
func UpdateProduct(c *gin.Context) {
	product := models.Product{}

	err := c.BindJSON(&product)

	if err != nil {
		utils.RespFail(c.Writer, "輸入內容有誤")
		return
	}

	v := models.FindProduct(product.ID)
	if v == nil {
		utils.RespFail(c.Writer, "查無此產品")
		return
	}

	models.UpdateProduct(product)
	if err != nil {
		utils.RespFail(c.Writer, "更新產品資訊錯誤")
		return
	}

	utils.RespOK(c.Writer, product, "測試成功")
}

//	 GetProduct
//	 @Summary 查詢產品資訊
//		@Tags		商品
//		@Param productId query int true " "
//		@Success	200	{string}	json " "
//		@Router		/product/getProduct [get]
func GetProduct(c *gin.Context) {

	productId := c.Query("productId")
	intProductId, _ := strconv.Atoi(productId)

	v := models.FindProduct(uint(intProductId))

	if v == nil {
		utils.RespFail(c.Writer, "查無此產品")
		return
	}

	utils.RespOK(c.Writer, v, "測試成功")
}
