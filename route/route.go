package route

import (
	"GinDemo/app/controller"
	"GinDemo/library"
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
	route.GET("/index", middleware, new(controller.IndexController).Index)

	// user路由
	route.GET("/user", middleware, new(controller.UserController).Index)

	// 检验登录的路由
	route.POST("/checkLogin", middleware, new(controller.LoginController).CheckLogin)
}

func Run() {
	route := initRoute()

	// 获取配置信息
	config, err := library.ParseConfig()
	if err != nil {
		panic(err.Error())
	}

	// 从配置信息中获取请求的端口
	route.Run(":" + config.App.Prot)
}
