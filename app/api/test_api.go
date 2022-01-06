package api

import (
	"app-admin/app/global"
	"app-admin/app/model/common/response"
	"app-admin/app/model/system"
	"fmt"
	"github.com/gin-gonic/gin"
)

func TestApi(c* gin.Context)  {
	var appUsers []system.AppUsers
	err := global.MySql().Find(&appUsers).Error
	if err == nil {
		response.OkWithDetailed(&appUsers,"操作成功",c)
	}else {
		response.FailWithDetailed(&appUsers,"操作失败",c)
		fmt.Println("操作失败:" +err.Error())
	}

}
