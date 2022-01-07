package initialize

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var CONF *Config

type Config struct {
	Port string `yaml:"port"`
	Mode string `yaml:"mode"`
	Databases Databases `yaml:"databases"`
}

type Databases struct {
	Databases string `yaml:"databases"`
	Url string `yaml:"url"`
	Port string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func  InitConfig(path string) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	CONF = c
}


