package utils

import (
	"crypto/md5"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/skycandyzhe/go-com/mypath"
)

//下载文件到临时路径
func DownLoadFileToTemp(DownloadUrl string) (string, error) {
	res, err := http.Get(DownloadUrl)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	filename := filepath.Base(DownloadUrl)
	tempdir := os.TempDir()
	new_file := path.Join(tempdir, filename)

	if mypath.IsExist(new_file) {
		os.Remove(new_file)
	}

	f, err := os.Create(new_file)
	if err != nil {
		return "", err
	}
	io.Copy(f, res.Body)
	f.Close()
	return f.Name(), nil
}

//获取文件hash
func Hash_Md5(path string) ([]byte, error) {
	filein, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	hs := md5.New()
	io.Copy(hs, filein)
	return hs.Sum(nil), nil
}
