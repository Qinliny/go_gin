package controller

import (
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	BaseController
}

// index控制器
func (index IndexController) Index(context *gin.Context) {
	index.Success(context, nil, "index请求成功")
}
