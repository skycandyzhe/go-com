package stringUtil

import (
	"github.com/dlclark/regexp2"
)

// import os.path
// import re

type StringSet struct {
	// filetype string //可能有普通文件 与压缩文件如zip文件以及office文件
	Content []string
}

// var reg_str_utf8 = regexp2.MustCompile(`[\x1f-\x7e]{6,}`, 0)
// 测试字符串查找 {6，} 找到9103条数据  5：11602   7 8332  8  7320   9 6954   10 6233
var reg_str_utf16le = regexp2.MustCompile(`[\u4e00-\u9fa5\x1f-\x7e]{6,}`, 0)

var MAX_STRINGLEN = 1024
var MAX_STRSIZE = 2048

/*
string 字符串提取
*/
func GetContentStringSet(content string) []string {
	var strSet []string
	// var con = ""

	// reg_str_utf16le.FindRunesMatch([]rune(content))
	var m *regexp2.Match
	// m, _ = reg_str_utf8.FindStringMatch(content)
	// for m != nil {
	// 	fmt.Println(m.String())
	// 	strSet.Content = append(strSet.Content, m.String())
	// 	m, _ = reg_str_utf8.FindNextMatch(m)
	// }
	// fmt.Println("---------")
	m, _ = reg_str_utf16le.FindStringMatch(content)
	var index = 0
	for m != nil && index < MAX_STRSIZE {
		index++
		// fmt.Println(m.String())
		if len(m.String()) > MAX_STRINGLEN {
			strSet = append(strSet, m.String()[:MAX_STRINGLEN])
		} else {
			strSet = append(strSet, m.String())
		}
		m, _ = reg_str_utf16le.FindNextMatch(m)
	}

	return strSet

}

// class Strings():
//     """Extract strings from analyzed file."""
//     MAX_FILESIZE = 16 * 1024 * 1024
//     MAX_STRINGCNT = 2048
//     MAX_STRINGLEN = 1024

//     def __init__(self, file_path):
//         self.file_path = file_path

//     def run(self):
//         """Run extract of printable strings.
//         @return: list of printable strings.
//         """
//         self.key = "strings"
//         strings = []

//         for s in re.findall(b"[\x1f-\x7e]{6,}", data):
//             strings.append(s.decode("utf-8"))
//         for s in re.findall(b"(?:[\x1f-\x7e][\x00]){6,}", data):
//             strings.append(s.decode("utf-16le"))

//         # Now limit the amount & length of the strings.
//         strings = strings[:self.MAX_STRINGCNT]
//         for idx, s in enumerate(strings):
//             strings[idx] = s[:self.MAX_STRINGLEN]

//         return strings
