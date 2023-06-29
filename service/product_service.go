package service

import (
	"ecom/models"

	"github.com/gin-gonic/gin"
)

//	 CreateProduct
//	 @Summary 新增商品
//		@Tags		商品
//		@Param product body models.CreateProduct true " "
//		@Success	200	{string}	json " "
//		@schemes
//		@Router		/product/createProduct [post]
func CreateProduct(c *gin.Context) {
	product := models.Product{}

	c.BindJSON(&product)

}
