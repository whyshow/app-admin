package config

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

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
}
