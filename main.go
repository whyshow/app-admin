package main

import (
	"app-admin/app/initialize"
	"app-admin/app/router"
)
func main() {
	//gin.SetMode(gin.ReleaseMode)
	initialize.Include(router.Routers)
	r := initialize.Init()
	r.Run(":8090") // 监听并在 0.0.0.0:8080 上启动服务
}
