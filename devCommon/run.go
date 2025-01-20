package common

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

func Run(router *gin.Engine, srvName, addr string, stop func()) {
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		log.Printf("%s running in %s \n", srvName, srv.Addr)
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	// SIGINT 用户发送 Ctrl + C 触发
	// SIGTERM 结束程序（可以被捕获、阻塞或忽略）
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Printf("Shutdown Project %s \n", srvName)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if stop != nil {
		stop()
	}

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("%s Shutdown cause by %v ", srvName, err)
	}
	select {
	case <-ctx.Done():
		log.Println("wait timeout...")
	}
	log.Printf("%s stopped successfully \n", srvName)
}
