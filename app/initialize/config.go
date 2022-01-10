package initialize

import (
	"encoding/json"
	"fmt"
	"os"
)

var CONF *Config

type Config struct {
	Http Http `json:"http"`
	Databases Databases `json:"databases"`
}

type Http struct {
	Port string `json:"port"`
	Mode string `json:"mode"`
}

type Databases struct {
	Databases string `json:"databases"`
	Url string `json:"url"`
	Port string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetConfig(path string) (Config, error) {
	// 打开文件
	file, _ := os.Open(path)
	// 关闭文件
	defer file.Close()
	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(file)
	conf := Config{}
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}
	CONF = &conf
	return conf, nil
}


