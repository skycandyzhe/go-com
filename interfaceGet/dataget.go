package interfaceGet

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// var err_parse = errors.New("can't parse")
var err_out_index = errors.New("out of index")
var err_not_key = errors.New("not has key")
var err_convert_map = errors.New("can't convert to map")
var err_convert_list = errors.New("can't convert to list")
var err_convert_str = errors.New("can't convert to str")
var err_convert_int = errors.New("can't convert to int")
var err_convert_error = errors.New("can't convert to need data type")
var arrar_num = regexp.MustCompile(`\[(\d+)\]`)

// 需要动态获取复杂类型的数据
// [5].id,   data[5].pool.aliced.5
func Get(keyname string, data interface{}) (interface{}, error) {
	deal_data := data
	var err error
	keys := strings.Split(keyname, ".")
	for _, key := range keys {
		beforeKey := strings.Split(key, "[")[0]
		if beforeKey != "" {
			deal_data, err = GetMap(beforeKey, deal_data)
			if err != nil {
				return nil, err
			}
		}
		nums := arrar_num.FindStringSubmatch(key)
		if len(nums) >= 2 {
			for i := 1; i < len(nums); i++ {
				array_index, _ := strconv.Atoi(nums[i])
				deal_data, err = GetList(array_index, deal_data)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return deal_data, nil
}
func GetMap(keyname string, data interface{}) (interface{}, error) {
	data_map, ok := data.(map[string]interface{})
	if !ok {
		return nil, err_convert_map
	}
	val, has := data_map[keyname]
	if has {
		return val, nil
	}
	return nil, err_not_key
}

func GetMapStr(keyname string, data interface{}) (string, error) {
	data_map, ok := data.(map[string]interface{})
	if !ok {
		return "", err_convert_map
	}
	val, has := data_map[keyname]
	if has {
		if var_str, ok := val.(string); ok {
			return var_str, nil
		} else {
			return "", err_convert_str
		}

	}
	return "", err_not_key
}
func GetMapInt(keyname string, data interface{}) (int, error) {
	data_map, ok := data.(map[string]interface{})
	if !ok {
		return 0, err_convert_map
	}
	val, has := data_map[keyname]
	if has {
		if var_int, ok := val.(float64); ok {
			return int(var_int), nil
		} else {
			return 0, err_convert_int
		}

	}
	return 0, err_not_key
}
func GetMapFloat(keyname string, data interface{}) (float64, error) {
	data_map, ok := data.(map[string]interface{})
	if !ok {
		return 0, err_convert_map
	}
	val, has := data_map[keyname]
	if has {
		if var_float, ok := val.(float64); ok {
			return var_float, nil
		} else {
			return 0, err_convert_error
		}

	}
	return 0, err_not_key
}
func GetList(index int, data interface{}) (interface{}, error) {
	data_vec, ok := data.([]interface{})
	if !ok {
		return nil, err_convert_list
	}
	if index >= len(data_vec) {
		return nil, err_out_index
	}
	return data_vec[index], nil
}
