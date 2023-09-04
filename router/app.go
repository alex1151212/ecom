package router

import (
	docs "ecom/docs"
	"ecom/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {

	r := gin.Default()
	r.Use(cors.New(CorsConfig()))
	//swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//ping pong
	r.GET("/", service.Ping)

	user := r.Group("/user")
	user.POST("/createUser", service.CreateUser)
	user.POST("/deleteUser", service.DeleteUser)
	user.POST("/updateUser", service.UpdateUser)
	user.POST("/uploadUserAvatar", service.UploadUserAvatar)

	product := r.Group("/product")
	product.GET("/getProduct", service.GetProduct)
	product.POST("/createProduct", service.CreateProduct)
	product.POST("/updateProduct", service.UpdateProduct)

	r.POST("/login", service.LoginUser)
	auth := r.Group("/auth")

	auth.Use(service.Auth().MiddlewareFunc())
	{
		auth.GET("/hello", service.HelloHandler)
		auth.POST("/addFavouriteProduct", service.AddFavouriteProduct)
		auth.GET("/getFavouriteProduct", service.GetFavouriteProduct)
		auth.POST("/logout", service.Auth().LogoutHandler)
	}

	return r
}
