package _interface

import (
	"fmt"
	"testing"
)

type A struct {
	Name string
}


func a1(obj A){
	obj.Name="2"
}


func a2(obj *A){
	obj.Name="2"
}

func a3(obj map[string]string){
	obj["a"]="2"
}

func TestReferences(t *testing.T){
	//测试map是引用类型
	var obj2 map[string]string
	obj2=map[string]string{}
	obj2["a"]="1"
	a3(obj2)
	fmt.Println(obj2["a"])
	return

	//测试指针是引用类型
	var obj1 *A
	obj1=&A{Name:"1"}
	a2(obj1)
	fmt.Println(obj1.Name)
	return

	//测试struct不是引用类型
	var obj0 A
	obj0=A{Name:"1"}
	a1(obj0)
	fmt.Println(obj0.Name)
	return

	//测试*，&的使用
	var obj *A
	obj=&A{Name:"a"}

	fmt.Println(obj.Name)
	fmt.Println((*obj).Name)

}
