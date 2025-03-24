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
	log.Println("User Router Initialized Successfully")
	router.Register(&RouterUser{})
}

// Route 处理具体接口信息
func (routerUser *RouterUser) Route(r *gin.Engine) {
	// 初始化 grpc 的客户端的连接
	rpc.InitGrpcUserClient()

	h := New()
	group := r.Group("/project")
	group.POST("/login/getCaptcha", h.getCaptcha)
	group.POST("/login/register", h.register)
	group.POST("/login", h.login)
	org := group.Group("/organization")
	org.Use(middle.TokenVerify())
	org.POST("/_getOrgList", h.myOrgList)
}
