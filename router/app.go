package router

import (
	docs "ecom/docs"
	"ecom/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {

	r := gin.Default()

	//swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//ping pong
	r.GET("/", service.Ping)

	user := r.Group("/user")
	user.POST("/createUser", service.CreateUser)
	user.POST("/deleteUser", service.DeleteUser)
	user.POST("/updateUser", service.UpdateUser)

	r.POST("/product/createProduct", service.CreateProduct)

	r.POST("/login", service.LoginUser)
	auth := r.Group("/auth")
	auth.POST("/logout", service.Auth().LogoutHandler)
	auth.Use(service.Auth().MiddlewareFunc())
	{
		auth.GET("/hello", service.HelloHandler)
	}
	return r
}
