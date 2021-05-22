package app

import "github.com/gin-gonic/gin"

// 默认的请求常量
const (
	SUCCESS = 200
	FAILED = 500
	SUCCESS_MSG = "请求成功"
	ERROR_MSG = "请求失败"
)

// 返回请求体
type ResponseData struct {
	Code		int64			`json:"code"`
	Message		string			`json:"message"`
	Data		interface{}		`json:"data"`
}

// 请求成功返回的参数
func Success(context *gin.Context,data interface{}, message string) {
	if message == "" {
		message = SUCCESS_MSG
	}
	resp := ResponseData{
		Code: SUCCESS,
		Message: message,
		Data: data,
	}
	context.JSON(SUCCESS, resp)
}

// 请求失败返回的参数
func Failed(context *gin.Context, message string) {
	if message == "" {
		message = ERROR_MSG
	}
	resp := ResponseData{
		Code: FAILED,
		Message: message,
		Data: nil,
	}
	context.JSON(SUCCESS, resp)
}

