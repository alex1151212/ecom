package main

import (
	"ecom/router"
	"ecom/utils"
)

func main() {

	utils.InitConfig()
	utils.InitMySQL()

	r := router.Router()

	r.Run(":8080")
}
