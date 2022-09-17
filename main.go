package main

import (
	"base_frame/core"
	"base_frame/global"
	"context"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 初始化日志库
	global.GLOBAL_LOG = core.InitLogger()
	// 初始化mysql
	global.GLOBAL_DB = core.InitDB()
	// 初始化redis
	if err := core.InitCache(); err != nil {
		log.Fatalf("redis conn failed, err=%v\n", zap.Error(err))
		return
	}
	log.Println("redis conn succeed")
	r := core.InitRouters()

	srv := &http.Server{
		Addr:    ":8090",
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("listen: %v\n", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown:%v", err)
	}
	log.Println("Server exiting")
}
