package global

import (
	"app-admin/app/initialize"
	"gorm.io/gorm"
)

func MySql() *gorm.DB {
	return initialize.DB
}