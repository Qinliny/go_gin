package controller

import (
	"Qinly/app"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (user UserController) Index(context *gin.Context) {
	app.Success(context, nil, "user请求成功")
}
