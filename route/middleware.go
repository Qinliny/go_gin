package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
}

func (m *Middleware) Middleware(c *gin.Context) {
	fmt.Println("我是一个请求中间件...")
	c.Next()
}
