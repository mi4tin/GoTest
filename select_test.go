package main

import (
	"testing"
	"fmt"
	"sync"
)

func TestCloseChannel(t *testing.T) {
	c1 := make(chan interface{},1);c1<-"c1a";close(c1)
	c2 := make(chan interface{}); close(c2)
	c3 := make(chan interface{})//; close(c3)

	var c1Count, c2Count, c3Count int
	for i := 10; i >= 0; i-- {
		select {
		case c1v:=<-c1:
			fmt.Println(c1v)
			c1Count++
		case <-c2:
			c2Count++
		case <-c3:
			c3Count++
		}
	}
	fmt.Printf("c1Count: %d\nc2Count: %d\nc3Count: %d\n", c1Count, c2Count, c3Count)
}

//代表此channel只能写，不能读
func producer(out chan <- int)  {
	for i:=0; i<=10; i++ {
		out<-i*i
	}
	close(out)
}

func consumer(in <-chan int)  {
	for num := range in{
		fmt.Println("num = ", num)
	}
}

//单向chan应用
func TestSingleChan(t *testing.T){
	//创建一个双向通道ch
	ch  := make(chan int)

	//生产者，生产数字，写入channel
	go producer(ch) //chennel传参，引用传递

	//消费者，从channel里读取数字、然后打印
	consumer(ch)
}

func TestSelect(t *testing.T){
	wg:=sync.WaitGroup{}
	wg.Add(1)
	//var c2 <-chan interface{}
	c0:=make(chan  interface{},10)
	read:=make(<-chan interface{},10)//单向读
	write:=make(chan<- interface{},10)//单向写

	//var c3 chan<- interface{}
	fmt.Println("aa0")
	go func() {
		fmt.Println("write chan")
		c0<-"chan123"
		write<-"chanWrite123"
		read=c0//双向转单向读
		//fmt.Println("c1:",<- c1)
	}()

	go func() {
		for{
			select {
			case c1v:=<- read:
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