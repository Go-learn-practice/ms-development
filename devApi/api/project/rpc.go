package project

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"test.com/devGrpc/project"
)

var ProjectServiceClient project.ProjectServiceClient

func InitGrpcProjectClient() {
	// 从 api 导入 etcd 配置
	//etcdRegister := discovery.NewResolver(config.Conf.Etcd.Addrs, logs.LG)
	//resolver.Register(etcdRegister)

	// etcd:///pro
	conn, err := grpc.Dial("0.0.0.0:8882", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	log.Println("grpc client connected project successfully")
	ProjectServiceClient = project.NewProjectServiceClient(conn)
}
