package main

import (
	"fmt"
	"moqikaka.com/goutil/logUtil"
	"runtime"
	"sync"
	"testing"
)


 type  a struct {


}

func TestM(t *testing.T){
	fmt.Println(runtime.NumCPU())

	runtime.GOMAXPROCS(2)
	fmt.Println(runtime.NumCPU())

}

func TestGoRoutine(t *testing.T){
	fmt.Println("pre:",runtime.GOMAXPROCS(8))
	var wg sync.WaitGroup
	wg.Add(1)
fmt.Println("start")
	//return
	//j:=0
	//go func() {
	//	for{
	//		j++
	//		//fmt.Println(11)
	//	}
	//}()
	//
	//j1:=0
	//go func() {
	//	for{
	//		j1++
	//		//fmt.Println(11)
	//	}
	//}()
	////
	//j2:=0
	//go func() {
	//	for{
	//		j2++
	//		//fmt.Println("j2:",11)
	//	}
	//}()
	////
	//j3:=0
	//go func() {
	//	for{
	//		j3++
	//		//time.Sleep(10)
	//		//fmt.Println("j3:",j3)
	//	}
	//}()
	//

	//f, _ := os.OpenFile("TestGoRoutine.log", os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND,0755)
	//
	//os.Stdout = f


	i:=0
	logUtil.InfoLog("i:%v",i)

	go func() {
		for{
			i++
			fmt.Println("i:",i)
			logUtil.InfoLog("i:%v",i)
		}
	}()

	wg.Wait()
}