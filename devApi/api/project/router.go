package project

import (
	"github.com/gin-gonic/gin"
	"log"
	"test.com/devApi/api/middle"
	"test.com/devApi/router"
)

type RouterProject struct {
}

// 注册路由
func init() {
	log.Println("User Router Initialized")
	router.Register(&RouterProject{})
}

// Route 处理具体接口信息
func (routerProject *RouterProject) Route(r *gin.Engine) {
	// 初始化 grpc 的客户端的连接
	InitGrpcProjectClient()

	h := New()
	group := r.Group("/project")
	group.Use(middle.TokenVerify())
	group.POST("/index", h.index)
}
