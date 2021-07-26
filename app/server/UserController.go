package server

import (
	"github.com/gin-gonic/gin"
)

func UserIndex(context *gin.Context) {
	Success(context, nil, "user请求成功")
}
