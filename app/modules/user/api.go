package user

import (
	"app-admin/app/common"
	"app-admin/app/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ApiTest(c *gin.Context) {
	request, _ := GetUserByID(100000)
	common.OkWithDetailed(request, "操作成功", c)
	middleware.LOGGER.Log(logrus.InfoLevel, "操作成功")
}
