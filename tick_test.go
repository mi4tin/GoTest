package main

import (
	"testing"
	"time"
	"fmt"
)

func TestTick(t *testing.T){
	//c := time.Tick(5 * time.Second)
	//for {
	//	x:=<- c
	//	fmt.Println("x:",x)
	//}
}



func TestTick1(t *testing.T){
	var ch chan int
	//定时任务
	ticker := time.NewTicker(time.Second * 2)
	go func() {
		for range ticker.C {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		}
		fmt.Println("gofuncover")
		ch <- 1
	}()
	<-ch
}