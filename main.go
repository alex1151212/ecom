package main

import "ecom/router"

func main() {
	r := router.Router()

	r.Run(":8080")
}
