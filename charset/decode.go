package charset

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"mime/quotedprintable"

	"strings"

	"github.com/axgle/mahonia"
	"github.com/djimenez/iconv-go"
)

func DataConvert(data []byte, orgCharset string) (string, error) {
	orgCharset = strings.ToLower(orgCharset)
	if orgCharset == "utf-8" || orgCharset == "utf8" {
		return string(data), nil
	} else {
		data, err := iconv.ConvertString(string(data), orgCharset, "utf-8")
		if err != nil {
			return data, err
		}
		ret, err := Convert(string(data), orgCharset, "utf-8")
		return ret, err
	}
}

func Convert(src string, srcCode string, tagCode string) (string, error) {
	srcCoder := mahonia.NewDecoder(srcCode)
	if srcCoder == nil {
		return src, fmt.Errorf("not found decode %s", srcCode)
	}
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	if tagCoder == nil {
		return src, fmt.Errorf("not found decode %s", srcCode)
	}
	_, cdata, err := tagCoder.Translate([]byte(srcResult), true)
	return string(cdata), err
}

func DecodeStr(str string, EncodeType string) ([]byte, error) {
	EncodeType = strings.ToLower(EncodeType)
	if EncodeType == "base64" {
		return base64.StdEncoding.DecodeString(str)
	} else if EncodeType == "quoted-printable" {
		reader := strings.NewReader(str)
		decode := quotedprintable.NewReader(reader)
		return ioutil.ReadAll(decode)
	} else {
		return nil, fmt.Errorf("unkown encode format %s", EncodeType)
	}

}
