package user

import (
	"github.com/gin-gonic/gin"
)

func RoutersUser(e *gin.Engine) {
	e.GET("/", ApiTest)
}
