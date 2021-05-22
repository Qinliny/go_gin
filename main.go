package main

import (
	"GinDemo/route"
)

var configPath = "./config/config.json"

func main() {
	// 启动加载路由
	route.Run()
}