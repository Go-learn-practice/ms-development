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
	log.Println("Project Router Initialized Successfully")
	router.Register(&RouterProject{})
}

// Route 处理具体接口信息
func (rp *RouterProject) Route(r *gin.Engine) {
	// 初始化 grpc 的客户端的连接
	InitGrpcProjectClient()

	h := New()
	group := r.Group("/project")
	// 使用中间件
	group.Use(middle.TokenVerify())
	group.POST("/index", h.index)
	group.POST("/project/selfList", h.myProjectList)
	group.POST("/project", h.myProjectList)
	group.POST("/project_template", h.projectTemplate)
	group.POST("/project/save", h.projectSave)
	group.POST("/project/read", h.readProject)
	group.POST("/project/recycle", h.recycleProject)
	group.POST("/project/recovery", h.recoveryProject)
	group.POST("/project_collect/collect", h.collectProject)
	group.POST("/project/edit", h.editProject)
}
