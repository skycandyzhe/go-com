package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//解析yml文件
type BaseInfo struct {
	Version   string     `yaml:"version"`
	DebugFlag bool       `yaml:"debugFlag"`
	Timelimit int        `yaml:"timelimit"`
	Logs      LogsEntity `yaml:"logs"`
}

type LogsEntity struct {
	Log_level string `yaml:"log_level"`
	Log_path  string `yaml:"log_path"`
	Err_path  string `yaml:"err_path"`
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

func init() {
	Conf = &BaseInfo{}
	Conf = Conf.GetConf("config.yaml")
	fmt.Println("read config.yaml :", Conf)
}
