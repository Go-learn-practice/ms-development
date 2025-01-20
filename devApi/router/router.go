package router

import (
	"github.com/gin-gonic/gin"
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

var routers []Router

func InitRouter(r *gin.Engine) {
	//registerRouter := New()
	//registerRouter.Route(&user.RouterUser{}, r)

	for _, router := range routers {
		router.Route(r)
	}
}

// Register 注册路由
func Register(ros ...Router) {
	routers = append(routers, ros...)
}
