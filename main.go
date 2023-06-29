package main

import (
	"ecom/router"
	"ecom/utils"
)

func main() {

	// @securityDefinitions.apikey BearerAuth
	// @in header
	// @name Authorization

	utils.InitConfig()
	utils.InitMySQL()

	r := router.Router()

	r.Run(":8080")
}
