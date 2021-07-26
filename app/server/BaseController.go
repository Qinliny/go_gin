package server

import (
	"Qinly/library"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

// 返回请求体
type ResponseData struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 默认的请求常量
const (
	SuccessCode = 200
	FailedCode  = 500
	SuccessMsg  = "请求成功"
	ErrorMsg    = "请求失败"
)

// 请求成功返回的参数
func Success(context *gin.Context, data interface{}, message string) {
	if message == "" {
		message = SuccessMsg
	}
	resp := ResponseData{
		Code:    SuccessCode,
		Message: message,
		Data:    data,
	}
	context.JSON(SuccessCode, resp)
}

// 请求失败返回的参数
func Failed(context *gin.Context, message string) {
	if message == "" {
		message = ErrorMsg
	}
	resp := ResponseData{
		Code:    FailedCode,
		Message: message,
		Data:    nil,
	}
	context.JSON(FailedCode, resp)
}

// 加密函数形成Token
func EncodeToken(param map[string]string) string {
	token := jwt.New(jwt.SigningMethodES256)
	// 加密的参数
	claims := make(jwt.MapClaims)
	// token过期的时间
	claims["expirationTime"] = time.Now().Add(time.Hour * time.Duration(2)).Unix()
	// 当前时间
	claims["currentTime"] = time.Now().Unix()
	for key, val := range param {
		claims[key] = val
	}
	// 加密的参数
	token.Claims = claims
	// 获取私密钥匙
	if tokenString, err := token.SignedString([]byte(library.ConfigParam.App.PrivateKey)); err != nil {
		return ""
	} else {
		return tokenString
	}
}
