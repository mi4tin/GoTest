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
	fmt.Println("pre:",runtime.GOMAXPROCS(5))

	j:=0
	j1:=0
	j2:=0
	go func() {
		for{
			j++
			j1++
			j2++
			j=j1
			if(j%1000==0){
				fmt.Println("x",runtime.NumGoroutine())
			}
		}
	}()

	i:=0
	go func() {
		for{
			i++
			fmt.Println("x:",runtime.NumGoroutine())
		}
	}()

	i1:=0
	go func() {
		for{
			i1++
			fmt.Println("x:",runtime.NumGoroutine())
		}
	}()

	i2:=0
	go func() {
		for{
			i2++
			fmt.Println("x:",runtime.NumGoroutine())
		}
	}()

	i3:=0
	go func() {
		for{
			i3++
			fmt.Println("x:",runtime.NumGoroutine())
		}
	}()

	i4:=0
	go func() {
		for{
			i4++
			fmt.Println("x:",runtime.NumGoroutine())
		}
	}()

	i5:=0
	go func() {
		for{
			i5++
			fmt.Println("x:",runtime.NumGoroutine())
		}
	}()

	select {

	}

	//wg.Wait()
	return
}