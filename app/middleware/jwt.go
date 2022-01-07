package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// JWT 鉴权
// 只有经过验证的才可以请求路由
func JWT()gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("Validate")
		if  context.Request.Header.Get("token") != "token"  {
			context.JSON(401,gin.H{
				"code": 1,
				"message": "Unauthorized",
			})
			context.Abort()
			return
		}else {
			context.Next()
		}
		context.Next()
	}
}