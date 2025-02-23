package config

import (
	"encoding/json"
	"os"
)

var CONF *Config

type Config struct {
	Http      Http      `json:"http"`
	Databases Databases `json:"databases"`
}
type Databases struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Http struct {
	Port string `json:"port"`
	Mode string `json:"mode"`
}

func OpenHttpConfig(path string) Http {
	// 打开文件
	file, _ := os.Open(path)
	// 关闭文件
	defer file.Close()
	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(file)
	conf := Http{}
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	err := decoder.Decode(&conf)
	if err != nil {
		panic(err)
	}
	return conf
}
func OpenConfig(path string) Config {
	// 打开文件
	file, _ := os.Open(path)
	if file == nil {
		panic("config file is nil")
	}
	// 关闭文件
	defer file.Close()
	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(file)
	conf := Config{}
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	err := decoder.Decode(&conf)
	if err != nil {
		panic(err)
	}
	CONF = &conf
	return conf
}

// OpenDBConfig 打开数据库配置
func OpenDBConfig(path string) Databases {
	// 打开文件
	file, _ := os.Open(path)
	if file == nil {
		panic("config file is nil")
	}
	// 关闭文件
	defer file.Close()
	decoder := json.NewDecoder(file)
	dbConf := Databases{}
	err := decoder.Decode(&dbConf)
	if err != nil {
		panic(err)
	}
	return dbConf
}
