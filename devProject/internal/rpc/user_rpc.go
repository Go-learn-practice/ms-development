package rpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"test.com/devGrpc/user/login"
)

var LoginServiceClient login.LoginServiceClient

func InitGrpcUserClient() {
	// 从 api 导入 etcd 配置
	//etcdRegister := discovery.NewResolver(config.Conf.Etcd.Addrs, logs.LG)
	//resolver.Register(etcdRegister)
	// etcd:///user
	conn, err := grpc.Dial("0.0.0.0:8881", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Println("grpc client connected user successfully")
	LoginServiceClient = login.NewLoginServiceClient(conn)
}
