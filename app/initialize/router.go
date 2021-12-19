package initialize

import (
	"app-admin/app/middleware"
	"github.com/gin-gonic/gin"
)
type Option func(*gin.Engine)

var options []Option
// Include 注册app的路由配置
//func Include(opts ...Option) {
//	options = append(options, opts...)
//}

// Init 初始化
func Init(opts ...Option) *gin.Engine {
	r := gin.New()
	r.Use(middleware.Validate())
	options = append(options, opts...) // 注册app的路由配置
	for _, opt := range options {
		opt(r)
	}
	return r
}