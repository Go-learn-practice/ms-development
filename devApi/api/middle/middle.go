package middle

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"test.com/devApi/api/rpc"
	common "test.com/devCommon"
	"test.com/devCommon/errs"
	"test.com/devGrpc/user/login"
	"time"
)

func TokenVerify() func(*gin.Context) {
	return func(c *gin.Context) {
		result := &common.Result{}
		// 1. 从header中获取token
		token := c.GetHeader("Authorization")
		// 2. 调用rpc服务进行token人证
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		response, err := rpc.LoginServiceClient.TokenVerify(ctx, &login.LoginRequest{Token: token})
		if err != nil {
			code, msg := errs.ParseGrpcError(err)
			c.JSON(http.StatusOK, result.Fail(code, msg))
			c.Abort()
			return
		}
		// 3. 处理结果 认证通过 将信息放入 gin 的上下文 失败返回未登录
		c.Set("memberId", response.Member.Id)
		c.Next()
	}
}
