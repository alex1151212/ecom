package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping
//
//	@Tags
//	@Success	200	{string}	pong
//	@Router		/ [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}
