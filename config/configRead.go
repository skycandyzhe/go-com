package config

import (
	"fmt"
	"os"

	"github.com/skycandyzhe/go-com/mypath"
	"gopkg.in/yaml.v2"
)

// 解析yml文件
type BaseInfo struct {
	Version   string `yaml:"version"`
	DebugFlag bool   `yaml:"debugFlag"`
	Console   bool   `yaml:"console"`
	LogPath   string `yaml:"log_path"`
}

func (c *BaseInfo) saveConf(filepath string) error {

	output, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath, output, 0660)
}
func GetDefaultConf() *BaseInfo {
	return GetConf("log_config.yaml")
}
func GetConf(configPath string) *BaseInfo {

	if mypath.FileExists(configPath) {
		yamlFile, err := os.ReadFile(configPath)
		if err != nil {
			fmt.Println(err.Error())
		}
		var conf BaseInfo
		err = yaml.Unmarshal(yamlFile, &conf)
		if err != nil {
			fmt.Println(err.Error())
		}
		return &conf
	} else {
		Conf := &BaseInfo{
			Version:   "v1.0",
			DebugFlag: true,
			Console:   false,
			LogPath:   "logs/",
		}
		Conf.saveConf(configPath)
		return Conf
	}
}
