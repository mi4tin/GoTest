package main

import (
	"testing"
	"fmt"
)

func TestChan(t *testing.T) {
	//close()
	//var b [5]int
	cha := make(chan int)
	fmt.Println("go000")
	cha<-1
	fmt.Println("go111")
	go func() {
		a,ok:=<-cha
		fmt.Println("chan:",a,ok)
	}()


	//close(cha)
}