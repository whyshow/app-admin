package api

import (
	"app-admin/app/global"
	"app-admin/app/model/common/response"
	"app-admin/app/model/system"
	"github.com/gin-gonic/gin"
)

func TestApi(c* gin.Context)  {
	var appUsers []system.AppUsers
	global.MySql().Find(&appUsers)
	response.OkWithDetailed(&appUsers,"操作成功",c)
}
