package main

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"
)

func deadloop() {
	i:=0
	for{
		i++
		//runtime.Gosched()
	}
}

func worker() {
	for {
		fmt.Println("worker is running")
		time.Sleep(time.Second * 1)
	}
}

func main() {
	fmt.Println("default gomax:",runtime.GOMAXPROCS(2))
	fmt.Println("cur gomax:",runtime.GOMAXPROCS(0))

	go func(){
		fmt.Println("hello world")
	}()

	fmt.Println("size:",unsafe.Sizeof(int64(1)))

	go func(){
		for {
			_=make([]int64, 8192)
		}
	}()
	fmt.Println("end")
	select {}


}