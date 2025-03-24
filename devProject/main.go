package main

import (
	"github.com/gin-gonic/gin"
	srv "test.com/devCommon"
	"test.com/devProject/config"
	"test.com/devProject/router"
)

func main() {
	r := gin.Default()

	// 注册 grpc 服务
	gc := router.RegisterGrpc()
	// grpc 服务注册到 etcd
	//router.RegisterEtcdServer()

	stop := func() {
		gc.Stop()
	}

	// 初始化 user grpc 客户端
	router.InitUserGrpc()

	// 启动服务
	srv.Run(r, config.Conf.SC.Name, config.Conf.SC.Addr, stop)
}
