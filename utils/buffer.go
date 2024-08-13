package utils

import (
	"bytes"
	"encoding/json"
)

func ValueToBuffer(v interface{}) (*bytes.Buffer, error) {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	// 这里要设置禁止html转义
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	return &buffer, err
}
