package mypath

import (
	"log"
	"os"
)

func Ext(path string) string {
	for i := len(path) - 2; i >= 0 && path[i] != '\\' && path[i] != '/'; i-- {
		if path[i] == '.' {
			return path[i+1:]
		}
	}
	return ""
}
func GetCurrentPath() string {
	p, _ := os.Getwd()
	return p
}

//check file
func PathExist(path string) (bool, error) {
	fi, err := os.Stat(path)
	return (err == nil || os.IsExist(err)) && fi.IsDir(), err
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		// fmt.Println(err)
		return false
	}
	return true
}

/**
判断目录是否存在
*/
func IsDir(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		log.Println(err)
		return false
	}
	return s.IsDir()
}

/**
判断文件是否存在
*/
func IsFile(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		log.Println(err)
		return false
	}
	return !s.IsDir()
}

/**
创建文件夹
*/
func createDir(dirName string) bool {
	err := os.Mkdir(dirName, 755)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
