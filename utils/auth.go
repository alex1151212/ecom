package utils

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AuthMiddleware(payloadFunc func(data interface{}) jwt.MapClaims,
	identityHandler func(*gin.Context) interface{},
	authenticator func(c *gin.Context) (interface{}, error),
	authorizator func(data interface{}, c *gin.Context) bool,
) *jwt.GinJWTMiddleware {

	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "test zone",
		SigningAlgorithm: viper.GetString("jwt.algorithm"),
		Key:              []byte(viper.GetString("jwt.secretKey")),
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,
		IdentityKey:      viper.GetString("jwt.identityKey"), //指定cookie的id
		PayloadFunc:      payloadFunc,
		IdentityHandler:  identityHandler,
		Authenticator:    authenticator, //登入驗證邏輯
		Authorizator:     authorizator,
		Unauthorized: func(c *gin.Context, code int, message string) {
			Resp(c.Writer, code, "", message)
		},
		// 指定從哪裡獲得token 格式為："<source>:<name>" 如有多個，用逗號隔開
		// "header: Authorization, query: token, cookie: jwt"
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	return authMiddleware

}
