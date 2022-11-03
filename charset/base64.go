package charset

import (
	"encoding/base64"
	"regexp"
)

func JudgeBase64(str string) bool {
	pattern := "^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{4}|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)$"
	matched, err := regexp.MatchString(pattern, str)
	if err != nil {
		return false
	}
	if !(len(str)%4 == 0 && matched) {
		return false
	}
	unCodeStr, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return false
	}
	tranStr := base64.StdEncoding.EncodeToString(unCodeStr)
	//return str==base64.StdEncoding.EncodeToString(unCodeStr)
	return str == tranStr
	// if str == tranStr {
	// 	return true
	// }
	// return false
}
func Base64Decode(encodeStr string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(encodeStr)
}

func Base64Encode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}
