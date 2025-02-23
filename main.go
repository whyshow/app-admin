package main

import (
	"app-admin/app/config"
	"app-admin/app/initialize"
	"app-admin/app/modules/user"
	"github.com/gin-gonic/gin"
)

func main() {
	config.OpenConfig("./app/config/config.json")
	initialize.InitDB(config.CONF.Databases)
	gin.SetMode(config.CONF.Http.Mode)     // 设置 release模式
	r := initialize.Init(user.RoutersUser) // 初始化 gin
	r.Run(":" + config.CONF.Http.Port)     // 监听并在 0.0.0.0:8080 上启动服务
}
