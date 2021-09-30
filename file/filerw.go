package file

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func CheckFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
func WriteStr(data string, filename string) error {
	var f *os.File
	var err error
	if CheckFileIsExist(filename) { //如果文件存在
		f, _ = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
	} else {
		f, _ = os.Create(filename) //创建文件
	}
	defer f.Close()
	_, err = io.WriteString(f, data) //写入文件(字符串)
	if err != nil {
		return err

	}
	return nil

}
func WriteByte(data []byte, filename string) {

	ioutil.WriteFile(filename, data, 0664)

}
func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}
func ReadAllString(filePth string) (string, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return "", err
	}
	defer f.Close()
	b, e := ioutil.ReadAll(f)
	if e != nil {
		return "", e
	} else {
		return string(b), e
	}
}
func ReadAllStringRmNull(filePth string) (string, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return "", err
	}
	defer f.Close()
	b, e := ioutil.ReadAll(f)
	var ret []byte
	for _, d := range b {
		if d != 0 {
			ret = append(ret, d)
		}
	}
	if e != nil {
		return "", e
	} else {
		return string(ret), e
	}
}
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

func FileExtraZip(filepath string) {

	// file, err := ioutil.TempFile("", "")
	// if err != nil {
	// 	logger.GetDefaultLogger().Warn("create tempfile failure:", err)
	// 	return ""
	// }
	// defer func() {
	// 	file.Close()
	// 	err = os.Remove(file.Name())
	// 	if err != nil {
	// 		logger.GetDefaultLogger().Warn("file remove failure:", err)
	// 		return ""
	// 	}
	// 	return ""
	// }()

	// if _, err := file.Write([]byte(buf)); err != nil {
	// 	panic(err)
	// }
	// // file.Close()
	// fmt.Println(file.Name())
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func Mkdir(path string) bool {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		fmt.Printf("mkdir failed![%v]\n", err)
	} else {
		fmt.Printf("mkdir %s success!\n", path)
		return true
	}
	return false
}
