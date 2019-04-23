package main

import (
	"testing"
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"time"
)

func TestOs(t *testing.T){
	signalChan := make(chan os.Signal, 1)
	// 捕捉 Ctrl+c 和 kill 信号，写入signalChan
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM,syscall.SIGKILL)
	// 此处执行处理逻辑
	//nsqd.Main()
	fmt.Println("执行处理逻辑")

	// signalChan阻塞进程
	go func() {
		for{
			sign:=<-signalChan
			if sign==syscall.SIGTERM||sign==syscall.SIGINT{
				fmt.Println("over")
				os.Exit(0)
			}
		}
	}()

	time.Sleep(100*time.Minute)
	// 捕捉信号后在Exit函数中处理信息，例如内存持久化等信息防止丢失
	//nsqd.Exit()
}