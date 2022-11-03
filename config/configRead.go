package config

import (
	"fmt"
	"io/ioutil"

	"github.com/skycandyzhe/go-com/mypath"
	"gopkg.in/yaml.v2"
)

// 解析yml文件
type BaseInfo struct {
	Version   string     `yaml:"version"`
	DebugFlag bool       `yaml:"debugFlag"`
	Console   bool       `yaml:"console"`
	Logs      LogsEntity `yaml:"logs"`
}

type LogsEntity struct {
	// Log_level string `yaml:"log_level"`
	LogName  string `yaml:"logname"`
	Log_path string `yaml:"log_path"`
	Err_path string `yaml:"err_path"`
}

func (c *BaseInfo) GetConf(filepath string) *BaseInfo {
	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

// info := config.BaseInfo{}
// conf := info.GetConf("config.yaml")
var Conf *BaseInfo

func GetDefaultConf() *BaseInfo {
	if mypath.FileExists("log_config.yaml") {
		Conf = &BaseInfo{}
		Conf = Conf.GetConf("log_config.yaml")
	} else {
		Conf = nil
	}
	return Conf
}
