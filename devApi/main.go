package main

import (
	"github.com/gin-gonic/gin"
	_ "test.com/devApi/api"
	"test.com/devApi/config"
	"test.com/devApi/router"
	srv "test.com/devCommon"
)

func main() {
	// gin 初始化配置
	r := gin.Default()

	// 注册路由
	router.InitRouter(r)

	// 启动服务
	srv.Run(r, config.Conf.SC.Name, config.Conf.SC.Addr, nil)
}
