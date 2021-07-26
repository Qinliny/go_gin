package server

import (
	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {
	Success(context, nil, "index请求成功")
}
