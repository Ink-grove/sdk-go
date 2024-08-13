package utils

import "encoding/json"

func MapToJson(param map[string]interface{}) []byte {
	dataType, _ := json.Marshal(param)
	return dataType
}
