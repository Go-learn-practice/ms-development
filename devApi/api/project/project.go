package project

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"test.com/devApi/pkg/model"
	"test.com/devApi/pkg/model/pro"
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
	var menusResp []*pro.MenusResponse
	_ = copier.Copy(&menusResp, indexResponse.Menus)
	c.JSON(http.StatusOK, resp.Success(menusResp))
}

func (p *HandlerProject) myProjectList(c *gin.Context) {
	resp := &common.Result{}
	// 1.获取参数
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	memberIdStr, _ := c.Get("memberId")
	memberId := memberIdStr.(int64)
	page := &model.Page{}
	page.Bind(c)
	// 2.调用 rpc 服务
	msg := &project.ProjectRpcRequest{MemberId: memberId, Page: page.Page, PageSize: page.PageSize}
	myProjectResponse, err := ProjectServiceClient.FindProjectByMemId(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, resp.Fail(code, msg))
	}
	if myProjectResponse.Pm == nil {
		myProjectResponse.Pm = []*project.ProjectMessage{}
	}
	var pms []*pro.ProjectAndMember
	_ = copier.Copy(&pms, myProjectResponse.Pm)
	// 3.处理结果
	c.JSON(http.StatusOK, resp.Success(gin.H{
		"list":  pms,
		"total": myProjectResponse.Total,
	}))
}
