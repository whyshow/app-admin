package initialize

import (
	"app-admin/app/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var (
	DB *gorm.DB
)
var Gorm = new(_gorm)

type _gorm struct{}

func InitDB(dbname string) {
	switch dbname {
	case "mysql":
		conf := config.Mysql{
			Path:        "127.0.0.1",
			Port:        "3306",
			Dbname:      "app_admin",
			Username:    "root",
			Password:    "swzhang3",
			MaxIdleCons: 2,
			MaxOpenCons: 100,
		}
		DB = GormMysqlByConfig(conf)
		if DB == nil {
			panic("db no init")
		}else {
			fmt.Println("数据库初始化完成")
		}
	}
}

// Config gorm 自定义配置
func (g *_gorm) Config() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	return config
}

// GormMysqlByConfig 初始化Mysql数据库用过传入配置
func GormMysqlByConfig(m config.Mysql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), Gorm.Config()); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleCons)
		sqlDB.SetMaxOpenConns(m.MaxOpenCons)
		return db
	}
}

