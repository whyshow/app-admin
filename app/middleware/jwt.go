package middleware

import (
	"github.com/gin-gonic/gin"
)

// Validate 鉴权
// 只有经过验证的才可以请求路由
func Validate()gin.HandlerFunc {
	return func(context *gin.Context) {
		//for k,v :=range context.Request.Header {
		//	fmt.Println(k,v)
		//}
		//if  context.Request.Header.Get("User-Agent") != "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.2 Safari/605.1.15"   {
		//	context.JSON(401,gin.H{
		//		"code": 1,
		//		"message": "Unauthorized",
		//	})
		//	context.Abort()
		//	return
		//}else {
		//	context.Next()
		//}
		context.Next()
	}
}