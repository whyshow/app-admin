package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Validate 鉴权
// 只有经过验证的才可以请求路由
func Validate()gin.HandlerFunc  {
	return func(context *gin.Context) {
		for k,v :=range context.Request.Header {
			fmt.Println(k,v)
		}
		if  context.Request.Header.Get("User-Agent") != "okhttp/3.14.9"   {
			context.JSON(401, gin.H{
				"code": 1,
				"message": "Unauthorized",
			})
			context.Abort()
			return
		}else {
			context.Next()
		}
	}
}