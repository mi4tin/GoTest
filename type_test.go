package main

import (
	"testing"
	"fmt"
	"sync/atomic"
)

func TestIntToString(t *testing.T){
	var BuildInfoData []*int32
	for j:=0;j<5;j++{
		fmt.Println(j,BuildInfoData[j])
	}
	return
	var i int64
	i=3
	atomic.StoreInt64(&i,5)
	fmt.Println("i:",i)
	return
	var a int32=1000000
	fmt.Println(a)
}
