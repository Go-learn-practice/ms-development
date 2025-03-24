package project

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"strconv"
	"test.com/devApi/pkg/model"
	"test.com/devApi/pkg/model/menu"
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
		return
	}
	menus := indexResponse.Menus
	var menusResp []*menu.Menu
	_ = copier.Copy(&menusResp, menus)
	c.JSON(http.StatusOK, resp.Success(menusResp))
}

func (p *HandlerProject) myProjectList(c *gin.Context) {
	resp := &common.Result{}
	// 1.获取参数
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	memberId := c.GetInt64("memberId")
	memberName := c.GetString("memberName")
	page := &model.Page{}
	page.Bind(c)
	selectBy := c.PostForm("selectBy")
	msg := &project.ProjectRpcRequest{
		MemberId:   memberId,
		MemberName: memberName,
		SelectBy:   selectBy,
		Page:       page.Page,
		PageSize:   page.PageSize,
	}
	// 2.调用 rpc 服务
	myProjectResponse, err := ProjectServiceClient.FindProjectByMemId(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, resp.Fail(code, msg))
		return
	}

	var pms []*pro.ProjectAndMember
	_ = copier.Copy(&pms, myProjectResponse.Pm)
	if pms == nil {
		pms = []*pro.ProjectAndMember{}
	}
	// 3.处理结果
	c.JSON(http.StatusOK, resp.Success(gin.H{
		"list":  pms,
		"total": myProjectResponse.Total,
	}))
}

func (p *HandlerProject) projectTemplate(c *gin.Context) {
	result := &common.Result{}
	// 1. 获取参数
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	memberId := c.GetInt64("memberId")
	memberName := c.GetString("memberName")
	page := &model.Page{}
	page.Bind(c)
	viewTypeStr := c.PostForm("viewType")
	viewType, _ := strconv.ParseInt(viewTypeStr, 10, 64)
	// 2. 调用 rpc 服务
	msg := &project.ProjectRpcRequest{
		MemberId:         memberId,
		MemberName:       memberName,
		ViewType:         int32(viewType),
		Page:             page.Page,
		PageSize:         page.PageSize,
		OrganizationCode: c.GetString("organizationCode"),
	}
	templateResponse, err := ProjectServiceClient.FindProjectTemplate(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}

	var pms []*pro.ProjectTemplate
	_ = copier.Copy(&pms, templateResponse.Ptm)
	// 3. 处理结果
	if pms == nil {
		pms = []*pro.ProjectTemplate{}
	}
	for _, pm := range pms {
		if pm.TaskStages == nil {
			pm.TaskStages = []*pro.TaskStagesOnlyName{}
		}
	}
	c.JSON(http.StatusOK, result.Success(gin.H{
		"list":  pms,
		"total": templateResponse.Total,
	}))
}

func (p *HandlerProject) projectSave(c *gin.Context) {
	result := &common.Result{}
	// 1. 获取参数
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	memberId := c.GetInt64("memberId")
	organizationCode := c.GetString("organizationCode")
	var req *pro.SaveProjectRequest
	_ = c.ShouldBind(&req)
	msg := &project.ProjectRpcRequest{
		MemberId:         memberId,
		OrganizationCode: organizationCode,
		TemplateCode:     req.TemplateCode,
		Name:             req.Name,
		Id:               int64(req.Id),
		Description:      req.Description,
	}
	// 2. 调用 rpc 服务
	saveProject, err := ProjectServiceClient.SaveProject(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	// 3. 处理结果
	var rsp *pro.SaveProject
	_ = copier.Copy(&rsp, saveProject)
	c.JSON(http.StatusOK, result.Success(rsp))
}

// readProject 项目详情
func (p *HandlerProject) readProject(c *gin.Context) {
	result := &common.Result{}
	projectCode := c.PostForm("projectCode")
	memberId := c.GetInt64("memberId")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	detail, err := ProjectServiceClient.FindProjectDetail(ctx, &project.ProjectRpcRequest{
		ProjectCode: projectCode,
		MemberId:    memberId,
	})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	pd := &pro.ProjectDetail{}
	_ = copier.Copy(&pd, detail)
	c.JSON(http.StatusOK, result.Success(pd))
}

// recycleProject 移动到回收站
func (p *HandlerProject) recycleProject(c *gin.Context) {
	result := &common.Result{}
	projectCode := c.PostForm("projectCode")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := ProjectServiceClient.UpdateDeletedProject(ctx, &project.ProjectRpcRequest{
		ProjectCode: projectCode,
		Deleted:     true,
	})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
	}
	c.JSON(http.StatusOK, result.Success(nil))
}

// recoveryProject 恢复项目
func (p *HandlerProject) recoveryProject(c *gin.Context) {
	result := &common.Result{}
	projectCode := c.PostForm("projectCode")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err := ProjectServiceClient.UpdateDeletedProject(ctx, &project.ProjectRpcRequest{
		ProjectCode: projectCode,
		Deleted:     false,
	})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	c.JSON(http.StatusOK, result.Success([]int{}))
}

// collectProject 收藏项目
func (p *HandlerProject) collectProject(c *gin.Context) {
	result := &common.Result{}
	projectCode := c.PostForm("projectCode")
	collectType := c.PostForm("type")
	// 1. 从gin获取保存的参数
	memberId := c.GetInt64("memberId")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// 2. 调用 rpc 服务
	_, err := ProjectServiceClient.UpdateCollectProject(ctx, &project.ProjectRpcRequest{
		ProjectCode: projectCode,
		MemberId:    memberId,
		CollectType: collectType,
	})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	c.JSON(http.StatusOK, result.Success([]int{}))
}

func (p *HandlerProject) editProject(c *gin.Context) {
	result := &common.Result{}
	var req *pro.ProjectReq
	_ = c.ShouldBind(&req)
	memberId := c.GetInt64("memberId")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &project.UpdateProjectMessage{}
	_ = copier.Copy(&msg, req)
	msg.MemberId = memberId
	_, err := ProjectServiceClient.UpdateProject(ctx, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		c.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	c.JSON(http.StatusOK, result.Success([]int{}))
}
