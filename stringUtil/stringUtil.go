package stringUtil

import (
	"encoding/json"

	"github.com/axgle/mahonia"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func ConverUtf16leToStr(src string) string {
	bs_UTF8LE, _, _ := transform.Bytes(unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder(), []byte(src))
	return string(bs_UTF8LE)
}
func ConverUtf16beToStr(src string) string {
	bs_UTF8LE, _, _ := transform.Bytes(unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM).NewDecoder(), []byte(src))
	return string(bs_UTF8LE)
}

func ValidUTF8(buf []byte) bool {
	nBytes := 0
	for i := 0; i < len(buf); i++ {
		if nBytes == 0 {
			if (buf[i] & 0x80) != 0 { //与操作之后不为0，说明首位为1
				for (buf[i] & 0x80) != 0 {
					buf[i] <<= 1 //左移一位
					nBytes++     //记录字符共占几个字节
				}

				if nBytes < 2 || nBytes > 6 { //因为UTF8编码单字符最多不超过6个字节
					return false
				}

				nBytes-- //减掉首字节的一个计数
			}
		} else { //处理多字节字符
			if buf[i]&0xc0 != 0x80 { //判断多字节后面的字节是否是10开头
				return false
			}
			nBytes--
		}
	}
	return nBytes == 0
}
func ConvertToString(src string, srcCode string, tagCode string) string {

	srcCoder := mahonia.NewDecoder(srcCode)

	srcResult := srcCoder.ConvertString(src)

	tagCoder := mahonia.NewDecoder(tagCode)

	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)

	result := string(cdata)

	return result

}

/*
boom split
EF BB BF  UTF-8
FE FF  UTF-16 (big-endian)
FF FE  UTF-16 (little-endian)
00 00 FE FF  UTF-32 (big-endian)
FF FE 00 00  UTF-32 (little-endian)
*/
func BoomSpit(src []byte) (data []byte, strtype string) {
	var len = len(src)
	if len < 4 {
		return src, ""
	}
	if src[0] == 0x00 && src[1] == 0x00 && src[2] == 0xfe && src[3] == 0xff {
		return src[4:], "utf32-be"
	} else if src[0] == 0xff && src[1] == 0xfe && src[2] == 0x00 && src[3] == 0x00 {
		return src[4:], "utf32-le"
	} else if src[0] == 0xef && src[1] == 0xbb && src[2] == 0xbf {
		return src[3:], "utf8"
	} else if src[0] == 0xff && src[1] == 0xfe {
		return src[2:], "utf16-le"
	} else if src[0] == 0xfe && src[1] == 0xff {
		return src[2:], "utf16-be"
	} else {
		return src, ""
	}
}

// 字符串拼接
func Str_Link(str ...string) string {

	return Str_LinkBySpecialChar(rune('_'), str...)
}

/*
 */
func Str_LinkBySpecialChar(joinstr rune, str ...string) string {

	var ret []rune
	// _str1:=strings.tr
	for _, tstr := range str {
		t := 0
		for _, b_ := range tstr {
			if b_ != joinstr {
				t++
				ret = append(ret, b_)
			}
		}
		ret = append(ret, joinstr)
	}
	if len(ret) > 0 {
		return string(ret[:len(ret)-1])
	}
	return ""
}

func InterfaceToJson(data interface{}) string {
	var ret []byte
	ret, err := json.Marshal(data)
	if err != nil {
		return err.Error()
	}
	return string(ret)
}
