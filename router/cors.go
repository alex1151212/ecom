package router

import (
	"time"

	"github.com/gin-contrib/cors"
)

func CorsConfig() cors.Config {

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://192.168.127.38:3000"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT"}
	corsConfig.AllowHeaders = []string{"Authorization", "Content-Type", "Upgrade", "Origin",
		"Connection", "Accept-Encoding", "Accept-Language", "Host", "Access-Control-Request-Method", "Access-Control-Request-Headers"}
	corsConfig.AllowCredentials = true
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.MaxAge = 12 * time.Hour

	return corsConfig
}
