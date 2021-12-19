package router

import (
	"app-admin/app/api"
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine)  {
	e.GET("/",api.TestApi)
}
