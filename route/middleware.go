package route

import (
	"github.com/gin-gonic/gin"
)

// 请求中间件
func Middleware(c *gin.Context) {
	c.Next()
}
