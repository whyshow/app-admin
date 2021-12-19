package main

import (
	"app-admin/app/initialize"
	"app-admin/app/router"
)
func main() {

	//gin.SetMode(gin.ReleaseMode) // 设置 release模式
	//initialize.Include(router.Routers) // 设置路由
	r := initialize.Init(router.Routers) // 初始化 gin
	r.Run(":8090") // 监听并在 0.0.0.0:8080 上启动服务
}
