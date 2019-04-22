package main

import (
	"testing"
	"fmt"
)

type ITest1Obj interface {
	Test()
	Test1()
}

type Obj1 struct {
	Y int
}

func (this *Obj1) Test(){
	fmt.Println("test7788",this.Y)
}

func (this *Obj1) Test1(){
	fmt.Println("test7788",this.Y)
}



func TestInterface(t *testing.T){
	obj1:=&Obj1{Y:100,}
	var iobj ITest1Obj
	iobj=obj1
	iobj.Test()
}



type ITestObj interface {
	GetPropX() int
	SetPropX(x int)
}

type Obj struct {
	X int
}

func (this *Obj) GetPropX() int {
	return this.X
}

func (this *Obj) SetPropX(x int) {
	this.X = x
}

func testFunc1(obj *Obj) {
	obj.X = obj.X + 1
}

func testFunc2(obj *Obj) {
	obj.SetPropX(obj.GetPropX() + 1)
}

func testFunc3(obj ITestObj) {
	obj.SetPropX(obj.GetPropX() + 1)
}

var myobj1 Obj

func Benchmark_Ptr(t *testing.B) {
	return
	for i := 0; i < t.N; i++ {
		testFunc1(&myobj1)
	}
}

func Benchmark_PtrAndCallF(t *testing.B) {
	return
	for i := 0; i < t.N; i++ {
		testFunc2(&myobj1)
	}
}

func Benchmark_Interface(t *testing.B) {
	return
	for i := 0; i < t.N; i++ {
		testFunc3(&myobj1)
	}
}
