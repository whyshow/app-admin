package global

import (
	"app-admin/app/initialize"
	"app-admin/app/middleware"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func MySql() *gorm.DB {
	return initialize.DB
}

func Log() *logrus.Logger {
	return middleware.LOGGER
}