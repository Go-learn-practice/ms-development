package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"test.com/devApi/api/middle"
	"test.com/devApi/api/rpc"
	"test.com/devApi/router"
)

type RouterUser struct {
}

// 注册路由
func init() {
	log.Println("User Router Initialized")
	router.Register(&RouterUser{})
}

// Route 处理具体接口信息
func (routerUser *RouterUser) Route(r *gin.Engine) {
	// 初始化 grpc 的客户端的连接
	rpc.InitGrpcUserClient()

	h := New()
	r.POST("/project/login/getCaptcha", h.getCaptcha)
	r.POST("/project/login/register", h.register)
	r.POST("/project/login", h.login)
	org := r.Group("/project/organization")
	org.Use(middle.TokenVerify())
	org.POST("/_getOrgList", h.myOrgList)
}
