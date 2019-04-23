package main

import (
	"testing"
	"fmt"
	"sync"
)

func TestSelect(t *testing.T){

	wg:=sync.WaitGroup{}
	wg.Add(1)
	//var c2 <-chan interface{}
	c1:=make(<-chan interface{},10)
	c0:=make(chan  interface{},10)

	//var c3 chan<- interface{}
	fmt.Println("aa0")
	go func() {
		fmt.Println("write chan")
		c0<-"chan123"
		c1=c0
		//fmt.Println("c1:",<- c1)
	}()

	go func() {
		for{
			select {
			case c1v:=<- c1:
				fmt.Println("c1v:",c1v)
				wg.Done()
			default:
				//println("default")
			}
		}
	}()

	wg.Wait()

	fmt.Println("over")
}