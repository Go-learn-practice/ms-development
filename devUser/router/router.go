package router

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"net"
	"test.com/devCommon/discovery"
	"test.com/devCommon/logs"
	"test.com/devGrpc/user/login"
	"test.com/devUser/config"
	loginServiceV1 "test.com/devUser/pkg/service/login.service.v1"
)

// Router 接口
type Router interface {
	Route(r *gin.Engine)
}

type RegisterRouter struct {
}

func New() *RegisterRouter {
	return &RegisterRouter{}
}

func (registerRouter *RegisterRouter) Route(ro Router, r *gin.Engine) {
	ro.Route(r)
}

type grpcConf struct {
	Addr         string
	RegisterFunc func(s *grpc.Server)
}

// RegisterGrpc 注册 grpc 服务
func RegisterGrpc() *grpc.Server {
	c := grpcConf{
		Addr: config.Conf.GC.Addr,
		RegisterFunc: func(s *grpc.Server) {
			// 注册 grpc 服务
			login.RegisterLoginServiceServer(s, loginServiceV1.NewLoginService())
		},
	}
	s := grpc.NewServer()
	c.RegisterFunc(s)

	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		log.Fatalf("failed to listen: %v \n", err)
	}

	go func() {
		// 启动 grpc 服务
		log.Printf("grpc server listening at %v \n", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v \n", err)
		}
	}()

	return s
}

// RegisterEtcdServer 注册 etcd 服务
func RegisterEtcdServer() {
	etcdRegister := discovery.NewResolver(config.Conf.Etcd.Addrs, logs.LG)
	resolver.Register(etcdRegister)
	// grpc 配置
	info := discovery.Server{
		Name:    config.Conf.GC.Name,
		Addr:    config.Conf.GC.Addr,
		Version: config.Conf.GC.Version,
		Weight:  config.Conf.GC.Weight,
	}
	r := discovery.NewRegister(config.Conf.Etcd.Addrs, logs.LG)
	_, err := r.Register(info, 2)
	if err != nil {
		log.Fatal(err)
	}
}
