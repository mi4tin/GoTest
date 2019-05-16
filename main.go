package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func ExitFunc()  {
	fmt.Println("开始退出...")
	fmt.Println("执行清理...")
	fmt.Println("结束退出...")
	os.Exit(0)
}



func main() {
	defertest1()
	defertest2()
	return
	signalChan := make(chan os.Signal, 10)
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

	//模拟业务程序运行
	time.Sleep(100*time.Minute)
	return

	//创建监听退出chan
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("退出", s)
				ExitFunc()
			default:
				fmt.Println("other", s)
			}
		}
	}()

	fmt.Println("进程启动...")
	sum := 0
	for {
		sum++
		fmt.Println("sum:", sum)
		time.Sleep(time.Second)
	}



	//fmt.Println("main1")
	//select {
	//
	//}
	//fmt.Println("main2")


	router := gin.Default()
	//v

	router.Use(func(f *gin.Context) {
		//fmt.Println("jsfsf:", f.)
		//f.en
		fmt.Println("ip", f.ClientIP())
		fmt.Println("relatevePath1:", f.RelativePath)
		//f.Params
	})
	v1 := router.Group("/api/v1/todos")
	{
		v1.GET("/:id", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"id": c.Param("id"),
			})
		})
	}

	router.GET("/test/:x/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": c.Param("x"),
		})
	})

	router.GET("/a/:a/b/:b", func(c *gin.Context) {
		fmt.Println("relatevePath2:", c.RelativePath)
		c.JSON(200, gin.H{
			"message": c.Param("a"),
		})
	})
	router.Run()
}

func test() {
	x := 100
	go func(a int) {
		a++
		println(a)
	}(x)

	x++
}

func closureTest() {
	y := 200
	go func() {
		y++
		println(y)
	}()
	y++
}
