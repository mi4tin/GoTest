package main

import (
	"testing"
	"fmt"
)


type Obj2 struct {
	Y string
}

func (this *Obj2) Test(){
	fmt.Println("test1",this.Y)
}

func (this *Obj2) Test1(){
	fmt.Println("test2",this.Y)
}

func TestConvertInterface(t *testing.T){

	type ITest2Obj interface {
		Test()
	}

	var te1 ITest2Obj=&Obj2{}
	te1.Test()
	return

	var a interface{}=&Obj2{}
	v,ok:=a.(ITest2Obj)

	fmt.Println("v:",v,"ok:",ok)

}

func TestType(t *testing.T){
	var  s  interface{}="123456"

	switch x:=s.(type) {
	case string:
		fmt.Println("string1")
		fmt.Println(x)
	default:
		fmt.Println("Not Type")
	}
}