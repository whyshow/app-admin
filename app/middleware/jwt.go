package middleware

import (
	"github.com/gin-gonic/gin"
)

// JWT 鉴权
// 只有经过验证的才可以请求路由
func JWT()gin.HandlerFunc {
	return func(context *gin.Context) {
		if  context.Request.Header.Get("sec-ch-ua-platform") != "\"macOS\""  {
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