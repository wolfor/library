// Convert project Convert.go
package Convert

import (
	"encoding/json"
	"errors"
	//	"reflect"
	"strconv"
	"strings"
)

//convert json format string to map[string]interface{}
func ConvertJsonToMap(jsonStr string) map[string]interface{} {
	//json convert to map
	var mapResult map[string]interface{}

	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	err := json.Unmarshal([]byte(jsonStr), &mapResult)

	if err != nil {
		return nil
	}

	return mapResult
}

//convert map[string]interface{} to json format string
func ConvertMapToJson(mapInstance map[string]interface{}) string {
	if mapInstance == nil {
		return ""
	}

	jsonStr, err := json.Marshal(mapInstance)

	if err != nil {
		return ""
	}

	return string(jsonStr)
}

//convert interface{} to float64
func ConvertInterfaceToFloat64(val interface{}) float64 {
	var (
		result float64
	)

	switch v := val.(type) {

	case string:
		result, _ = strconv.ParseFloat(v, 64)
	case int:
	case float32:
		result = float64(v)
	case float64:
		result = v
	case nil:
	default:
		panic(errors.New(strings.Join([]string{val.(string), "val not numeric"}, "")))
	}

	return result
}
