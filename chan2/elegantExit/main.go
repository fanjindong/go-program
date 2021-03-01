package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			fmt.Println("running")
		}
	}()

	sig := <-ch
	fmt.Println("接收到信号: " + sig.String())
	time.Sleep(3 * time.Second)
	fmt.Println("完成资源清理，优雅退出")
}
