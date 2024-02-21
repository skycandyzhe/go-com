package tests

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/skycandyzhe/go-com/utils"
)

func Test_Hash(t *testing.T) {

	root := `D:\codes\whl-samples\基础功能测试集\MalConfig`

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".dmp") || strings.HasSuffix(path, ".dat") ||
			strings.HasSuffix(path, ".py") || strings.HasSuffix(path, ".yar") ||
			strings.HasSuffix(path, ".pyc") || strings.HasSuffix(path, ".dump") ||
			strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, "tmp") ||
			strings.HasSuffix(path, ".json") || strings.HasSuffix(path, ".i64") ||
			strings.HasSuffix(path, ".DS_Store") || strings.HasSuffix(path, ".idb") {
			// os.Remove(path)
			return nil
		}

		MD5, _ := utils.Hash_Md5_Str(path)
		pre := `D:\codes\whl-samples\`
		if strings.HasPrefix(path, pre) {

			path = path[len(pre):]
		}
		fmt.Printf("%s\t%s\t \n", MD5, path)
		return nil
	})
}
