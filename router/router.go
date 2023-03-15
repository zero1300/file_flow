package router

import (
	"context"
	"file_flow/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type RegRouteFn = func(public *gin.RouterGroup, auth *gin.RouterGroup)

var regRouteFns []RegRouteFn

func RegRoute(fn RegRouteFn) {
	if fn == nil {
		return
	}
	regRouteFns = append(regRouteFns, fn)
}

// InitRoute 入口函数
func InitRoute() {
	notifyContext, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r := gin.Default()
	r.Use(middleware.Cors())

	public := r.Group("/api/v1/public")
	auth := r.Group("/api/v1/")

	auth.Use(middleware.Auth())

	InitBaseRoutes()
	for _, fn := range regRouteFns {
		fn(public, auth)
	}

	serverPort := viper.GetString("server.port")
	if serverPort == "" {
		serverPort = "8899"
	}

	server := &http.Server{Addr: fmt.Sprintf(":%s", serverPort), Handler: r}

	// http 服务器协程, 等待主协程的信号
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Println("Server error: " + err.Error())
			return
		}
	}()

	// 阻塞主协程, 直到收到停止信号量
	<-notifyContext.Done()

	// 给服务5秒钟清理现场
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	err := server.Shutdown(timeoutCtx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Stopped success.....")
}

func InitBaseRoutes() {
	userSetup()
	fileSetup()
}
