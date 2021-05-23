package controller

import (
	"Qinly/library"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type BaseController struct {
	*gin.Context
}

// 返回请求体
type ResponseData struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 默认的请求常量
const (
	SUCCESS     = 200
	FAILED      = 500
	SUCCESS_MSG = "请求成功"
	ERROR_MSG   = "请求失败"
)

// 请求成功返回的参数
func (base BaseController) Success(context *gin.Context, data interface{}, message string) {
	if message == "" {
		message = SUCCESS_MSG
	}
	resp := ResponseData{
		Code:    SUCCESS,
		Message: message,
		Data:    data,
	}
	context.JSON(SUCCESS, resp)
}

// 请求失败返回的参数
func (base BaseController) Failed(context *gin.Context, message string) {
	if message == "" {
		message = ERROR_MSG
	}
	resp := ResponseData{
		Code:    FAILED,
		Message: message,
		Data:    nil,
	}
	context.JSON(SUCCESS, resp)
}

// 加密函数形成Token
func (base BaseController) EncodeToken(param map[string]string) string {
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
	config, err := library.ParseConfig()
	if err != nil {
		panic(err.Error())
	}
	if tokenString, err := token.SignedString([]byte(config.App.PrivaeKey)); err != nil {
		return ""
	} else {
		return tokenString
	}
}
