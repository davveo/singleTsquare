package middleware

import "github.com/gin-gonic/gin"

func JwtMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {

		context.Next()
	}
}
