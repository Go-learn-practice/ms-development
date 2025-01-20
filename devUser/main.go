package main

import (
	"github.com/gin-gonic/gin"
	srv "test.com/devCommon"
	"test.com/devUser/config"
	"test.com/devUser/router"
)

func main() {
	// gin 初始化配置
	r := gin.Default()

	// 注册 grpc 服务
	gc := router.RegisterGrpc()
	// grpc 服务注册到 etcd
	//router.RegisterEtcdServer()

	stop := func() {
		gc.Stop()
	}

	// 启动服务
	srv.Run(r, config.Conf.SC.Name, config.Conf.SC.Addr, stop)
}
