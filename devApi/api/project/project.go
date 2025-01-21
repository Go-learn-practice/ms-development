package project

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	menus "test.com/devApi/pkg/model/project"
	common "test.com/devCommon"
	"test.com/devCommon/errs"
	"test.com/devGrpc/project"
	"time"
)

type HandlerProject struct {
}

func New() *HandlerProject {
	return &HandlerProject{}
}

func (p *HandlerProject) index(c *gin.Context) {
	resp := &common.Result{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// 调用 rpc 服务
	msg := &project.IndexRequest{}
	indexResponse, err := ProjectServiceClient.Index(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, resp.Fail(code, msg))
	}
	var menusResp []*menus.MenusResponse
	_ = copier.Copy(&menusResp, indexResponse.Menus)
	c.JSON(http.StatusOK, resp.Success(menusResp))
}
