package pyfeature

import "strings"

// data T mp[T] bool
// func PyIn(type T)(data T, mp map[T]bool) bool {
// 	_, ok := mp[data]
// 	return ok
// }

//TODO 以后有泛型 将string 转为泛型
func PyStrIn(data string, mp map[string]bool) bool {
	_, ok := mp[data]
	return ok
}

// 添加泛型断言实现
func PyStrInT(data string, mp interface{}) bool {
	switch mp := mp.(type) {
	case map[string]interface{}:
		_, ok := mp[data]
		return ok
	case map[string]bool:
		_, ok := mp[data]
		return ok
	case map[string]int:
		_, ok := mp[data]
		return ok
	case map[string]float32:
		_, ok := mp[data]
		return ok
	case map[string]float64:
		_, ok := mp[data]
		return ok
	default:
		return false
	}

}

//  strAll in
func PyStrAllinContent(content string, key []string) bool {
	for _, data := range key {
		if !strings.Contains(content, data) {
			return false
		}
	}
	return true

}

//  str Any in
func PyStrAnyinContent(content string, key []string) bool {
	for _, data := range key {
		if strings.Contains(content, data) {
			return true
		}
	}
	return false

}
