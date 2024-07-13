package main

import (
	"fmt"
	"giligili/conf"
	"giligili/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()

	// 优雅退出...
	go func() {
		r.Run(":8088")
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutdown Server ...: 需要补充自己的处理逻辑...")
}
