package controller

import (
	"Qinly/app"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
	BaseController
}



// index控制器
func (index IndexController) Index(context *gin.Context) {
	app.Success(context, nil, "index请求成功")
}
