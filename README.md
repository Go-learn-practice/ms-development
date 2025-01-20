# msDevelopment 开发

## Gin Web 框架
> 安装 `go get -u github.com/gin-gonic/gin`

**Gin 优雅地重启或停止**
```go
package main

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	// SIGINT 用户发送 Ctrl + C 触发
	// SIGTERM 结束程序（可以被捕获、阻塞或忽略）
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Project web Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
```

## Zap 日志库
> - Zap github [zap](https://github.com/uber-go/zap)
> - file rotate [file rotate](https://github.com/natefinch/lumberjack)

** Zap 日志库配置**

## Viper 文件配置
> viper

## grpc 服务
> 服务转发

## etcd 服务发现
> 在 api 中直接调用的地址，这种方式明显是不合适的，需要引入服务发现，etcd可以帮助我们完成这个过程

**地址：[ectd](https://github.com/etcd-io/etcd)**

安装库：
```shell
go get go.etcd.io/etcd/client/v3
```

## Gorm 数据库
> gorm