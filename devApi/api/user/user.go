package user

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
	"test.com/devApi/pkg/model/user"
	common "test.com/devCommon"
	"test.com/devCommon/errs"
	"test.com/devGrpc/user/login"
	"time"
)

type HandlerUser struct {
}

func New() *HandlerUser {
	return &HandlerUser{}
}

// 获取验证码处理
func (h *HandlerUser) getCaptcha(ctx *gin.Context) {
	resp := &common.Result{}
	// 1. 获取参数
	mobile := ctx.PostForm("mobile")
	fmt.Printf("当前获取到的手机号是：%s \n", mobile)
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	loginCaptchaResponse, err := LoginServiceClient.Captcha(c, &login.CaptchaRequest{Mobile: mobile})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, resp.Fail(code, msg))
		return
	}
	ctx.JSON(http.StatusOK, resp.Success(loginCaptchaResponse.Code))
}

func (h *HandlerUser) register(ctx *gin.Context) {
	// 1. 获取参数 参数模型
	resp := &common.Result{}
	var req user.RegisterReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(http.StatusBadRequest, "参数错误"))
		return
	}
	// 2. 参数校验 判断参数是否合法
	if err := req.Verify(); err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(http.StatusBadRequest, err.Error()))
		return
	}
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// 3. 调用 user grpc 服务 获取响应
	msg := &login.RegisterRequest{}
	err = copier.Copy(msg, req)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(http.StatusBadRequest, "参数解析错误"))
		return
	}
	_, err = LoginServiceClient.Register(c, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, resp.Fail(code, msg))
		return
	}
	// 4. 返回响应
	ctx.JSON(http.StatusOK, resp.Success(""))
}

func (h *HandlerUser) login(ctx *gin.Context) {
	// 1. 获取参数 参数模型
	resp := &common.Result{}
	var req user.LoginReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(http.StatusBadRequest, "参数错误"))
		return
	}
	// 2. 调用user grpc 完成登录
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	msg := &login.LoginRequest{}
	err = copier.Copy(msg, req)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(http.StatusBadRequest, err.Error()))
		return
	}
	loginResponse, err := LoginServiceClient.Login(c, msg)
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, resp.Fail(code, msg))
		return
	}
	rsp := &user.LoginRsp{}
	err = copier.Copy(rsp, loginResponse)
	if err != nil {
		ctx.JSON(http.StatusOK, resp.Fail(http.StatusBadRequest, err.Error()))
		return
	}
	// 3. 返回响应
	ctx.JSON(http.StatusOK, resp.Success(rsp))
}

func (h *HandlerUser) myOrgList(ctx *gin.Context) {
	result := &common.Result{}
	memberIdStr, _ := ctx.Get("memberId")
	memberId := memberIdStr.(int64)
	list, err := LoginServiceClient.MyOrgList(context.Background(), &login.UserRequest{MemId: memberId})
	if err != nil {
		code, msg := errs.ParseGrpcError(err)
		ctx.JSON(http.StatusOK, result.Fail(code, msg))
		return
	}
	if list.OrganizationList == nil {
		ctx.JSON(http.StatusOK, result.Success([]*user.OrganizationList{}))
		return
	}
	var orgs []*user.OrganizationList
	_ = copier.Copy(&orgs, list.OrganizationList)
	ctx.JSON(http.StatusOK, result.Success(orgs))
}
