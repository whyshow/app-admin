package main

import (
	"app-admin/app/initialize"
	"app-admin/app/router"
	"fmt"
	"github.com/gin-gonic/gin"
)
func main() {
	conf,_ := initialize.GetConfig("./app/config/config.json")
	fmt.Println(conf.Http.Port)
	initialize.InitDB("mysql")
	gin.SetMode(gin.ReleaseMode) // 设置 release模式
	r := initialize.Init(router.Routers) // 初始化 gin
	r.Run(":"+initialize.CONF.Http.Port) // 监听并在 0.0.0.0:8080 上启动服务
}


