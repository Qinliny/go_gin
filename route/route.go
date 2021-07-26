package route

import (
	"Qinly/app/server"
	"Qinly/library"
	"github.com/gin-gonic/gin"
)

// 初始化路由
func initRoute() *gin.Engine {
	route := gin.Default()
	// 加载路由
	loadRoute(route)
	return route
}

// 加载路由
func loadRoute(route *gin.Engine) {
	// index路由
	route.GET("/index", Middleware, server.Index)

	// user路由
	route.GET("/user", Middleware, server.UserIndex)

	// 检验登录的路由
	route.POST("/checkLogin", Middleware, server.CheckLogin)
}

func Run() {
	route := initRoute()
	// 从配置信息中获取请求的端口
	route.Run(":" + library.ConfigParam.App.Port)
}
