package app

// 默认的请求常量
const (
	SUCCESS     = 200
	FAILED      = 500
	SUCCESS_MSG = "请求成功"
	ERROR_MSG   = "请求失败"
)

// 返回请求体
type ResponseData struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
