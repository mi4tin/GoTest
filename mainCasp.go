package main

import (
	"fmt"
	"runtime"
)

func at(){
	i:=0
	i++
}
func main() {
	fmt.Println("pre:",runtime.GOMAXPROCS(3))
	//var wg sync.WaitGroup
	//wg.Add(1)


	j:=0
	go func() {
		for{
			j++
			if(j%1000==0){
				fmt.Println("j:",j)
			}
		}
	}()

	i:=0
	go func() {
		for{
			i++
			fmt.Println("NumGoroutine:",runtime.NumGoroutine(),"C:",i)
		}
	}()

	select {

	}

	//wg.Wait()
	return
}