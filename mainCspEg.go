package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
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
	fmt.Printf("There are %d cores.\n", runtime.NumCPU())

	go worker()

	go deadloop()

	i := 3
	for {
		fmt.Printf("main is running, i=%d\n", i)
		i--
		if i == 0 {
			//debug.PrintStack()
			runtime.GC()

			debug.PrintStack()
		}

		time.Sleep(time.Second * 1)
	}
}