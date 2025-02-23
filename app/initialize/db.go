package initialize

import (
	"app-admin/app/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Gorm = new(_gorm)

var DB *gorm.DB

type _gorm struct{}

type Mysql struct {
	Path        string `structure:"path" json:"path" yaml:"path"`                          // 服务器地址
	Port        string `structure:"port" json:"port" yaml:"port"`                          // 端口
	Config      string `structure:"config" json:"config" yaml:"config"`                    // 高级配置
	Dbname      string `structure:"db-name" json:"dbname" yaml:"db-name"`                  // 数据库名
	Username    string `structure:"username" json:"username" yaml:"username"`              // 数据库用户名
	Password    string `structure:"password" json:"password" yaml:"password"`              // 数据库密码
	MaxIdleCons int    `structure:"max-idle-cons" json:"maxIdleCons" yaml:"max-idle-cons"` // 空闲中的最大连接数
	MaxOpenCons int    `structure:"max-open-cons" json:"maxOpenCons" yaml:"max-open-cons"` // 打开到数据库的最大连接数
	LogMode     string `structure:"log-mode" json:"logMode" yaml:"log-mode"`               // 是否开启Gorm全局日志
	LogZap      bool   `structure:"log-zap" json:"logZap" yaml:"log-zap"`                  // 是否通过zap写入日志文件
}

// InitDB 初始化数据库
func InitDB(dbConfig config.Databases) {
	conf := Mysql{
		Path:        dbConfig.Url,
		Port:        dbConfig.Port,
		Dbname:      dbConfig.Name,
		Username:    dbConfig.Username,
		Password:    dbConfig.Password,
		MaxIdleCons: 10,
		MaxOpenCons: 100,
	}
	DB = GormMysqlByConfig(conf)
	if DB == nil {
		panic("db no init")
	} else {
		fmt.Println("数据库初始化完成")
	}
}

// Dsn 拼接数据库连接
func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
}

// GormConfig DbConfig gorm 自定义配置
func (g *_gorm) GormConfig() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	return config
}

// GormMysqlByConfig 初始化Mysql数据库用过传入配置
func GormMysqlByConfig(m Mysql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if gormDb, err := gorm.Open(mysql.New(mysqlConfig), Gorm.GormConfig()); err != nil {
		panic(err)
	} else {
		sqlDB, _ := gormDb.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleCons)
		sqlDB.SetMaxOpenConns(m.MaxOpenCons)
		return gormDb
	}
}
