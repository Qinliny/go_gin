package controller

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (user UserController) Index(context *gin.Context) {
	user.Success(context, nil, "user请求成功")
}
