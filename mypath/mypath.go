package mypath

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func Ext(path string) string {
	for i := len(path) - 2; i >= 0 && path[i] != '\\' && path[i] != '/'; i-- {
		if path[i] == '.' {
			return path[i+1:]
		}
	}
	return ""
}

// check file
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

/*
*
判断目录是否存在
*/
func DirExists(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		log.Println(err)
		return false
	}
	return s.IsDir()
}

/*
*
判断文件是否存在
*/
func FileExists(fileAddr string) bool {

	if s, err := os.Stat(fileAddr); os.IsNotExist(err) {
		return false
	} else {
		return !s.IsDir()
	}

}

/*
*
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

/*
  - @data: 2022-08-31 09:53:35
  - @athor: byy
  - @descript: 创建特定路径 创建失败返回error
    如果这个路径是文件返回error
  - @return:
*/
func CreateDirPath(path string) error {
	if IsExist(path) {
		if DirExists(path) {
			return nil
		} else {
			return errors.New("path is file")
		}
	}
	err := os.MkdirAll(path, 0660)
	return err
}

func NewTempFile() string {

	return ""
}

/*
* @data: 2022-08-31 10:34:31
* @athor: byy
* @descript: 删除路径 确保指定的路径不存在
* @return:
 */
func DelDirs(path string) error {
	abspath, err := filepath.Abs(path)
	if err != nil {
		return nil
	}
	return os.RemoveAll(abspath)
}

func WriteByte(data []byte, filename string) {
	ioutil.WriteFile(filename, data, 0664)

}

func ReadFileString(filePth string) (string, error) {
	data, err := ioutil.ReadFile(filePth)
	return string(data), err
}
func ReadFileStringRmNull(filePth string) (string, error) {
	data, err := ioutil.ReadFile(filePth)
	var ret []byte
	for _, d := range data {
		if d != 0 {
			ret = append(ret, d)
		}
	}
	if err != nil {
		return "", err
	} else {
		return string(ret), err
	}
}

// 读特定大小的区域
func ReadTargetSizeFileByte(filepath string, targetSize int64) ([]byte, error) {
	file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	var size = stat.Size()
	if size > targetSize {
		size = targetSize
	}
	// define read block size = 2
	buf := make([]byte, size)

	length, err := file.Read(buf)
	if err != nil {
		if err != io.EOF {
			return nil, err
		}
	}
	_ = length

	return buf, err
}

func ListDir(dirpath string) (dirs []string, files []string) {
	infos, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return
	}
	for _, info := range infos {
		if info.IsDir() {
			dirs = append(dirs, info.Name())
		} else {
			files = append(files, info.Name())
		}
	}
	return
}
