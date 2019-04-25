package main

import (
	"testing"
	"errors"
	"fmt"
)

type ITest2Obj interface {
	Test()
	Test1()
}

type Obj2 struct {
	Y string
}

func (this *Obj2) Test(){
	fmt.Println("test1",this.Y)
}

func (this *Obj2) Test1(){
	fmt.Println("test2",this.Y)
}

//todo:还未通过测试
func TestConvertInterface(t *testing.T){
	obj1:=errors.New("123")
	r,ok:=obj1.(ITest2Obj)
	fmt.Println(r,ok)
}