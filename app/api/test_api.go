package api

import (
	"app-admin/app/global"
	"app-admin/app/middleware"
	"app-admin/app/model/common/response"
	"app-admin/app/model/system"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func TestApi(c* gin.Context)  {
	//response.Ok(c)
	var appUsers []system.AppUsers
	err := global.MySql().Find(&appUsers).Error
	if err == nil {
		response.OkWithDetailed(&appUsers,"操作成功",c)
		middleware.LOGGER.Log(logrus.InfoLevel,"操作成功")
	}else {
		response.FailWithDetailed(&appUsers,"操作失败",c)
		fmt.Println("操作失败:" +err.Error())
	}

}
