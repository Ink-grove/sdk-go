package constant

import (
	"errors"
	"fmt"
)

var (
	ERR_INIT_CLIENT  = errors.New("client init error")
	ERR_PARAMS_INPUT = errors.New("params input error")
)

var (
	SUCCESS = 0
)

var errMsgDic = map[int]string{
	SUCCESS: "ok",
}

// GetErrMsg 获取错误描述
func GetErrMsg(code int) string {
	if msg, ok := errMsgDic[code]; ok {
		return msg
	}
	return fmt.Sprintf("unknown error code %d", code)
}
