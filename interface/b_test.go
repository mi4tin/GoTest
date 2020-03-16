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

func a4(obj []int){
	obj[0]=3
}

func addV(a []int32){
	a=append(a,2)
	fmt.Println(a)
}

func a5(obj map[string]string){
	obj["b"]="2"
	fmt.Println(obj)
}

func addV1(a *[]int32){
	//(*a)[0]=0
	//(*a)[1]=1
	//(*a)[2]=2
	*a=append(*a,3)
	*a=append(*a,4)
//	a=&b
	fmt.Println(a)
}

func TestReferences(t *testing.T){
	//测试make
	slice4:=make([]*int32,5)
	fmt.Println(slice4)

	//测试切片指针是引用类型。切片增加元素有用
	var slice2 * []int32
	slice2 = new([]int32)//给指针初始化
	//fmt.Println("a:",slice2)
	//*slice2 =make([]int32,3)//给切片初始化，分配一个切片内存
	*slice2=append(*slice2,1)
	fmt.Println("a:",slice2)
	addV1(slice2)
	fmt.Println(slice2)
	return

	//测试切片整体不是引用类型，测试切片增加元素无用
	slice1 :=[]int32{}
	slice1=make([]int32,2)
	slice1=append(slice1,1)
	fmt.Println(slice1)
	addV(slice1)
	fmt.Println(slice1)
	return

	//测试整个map是引用类型
	var obj3 map[string]string
	obj3=map[string]string{}
	obj3["a"]="1"
	fmt.Println(obj3)
	a5(obj3)
	fmt.Println(obj3)
	return

	//测试切片整体不是引用类型，测试切片增加元素无用
	slice3 :=[]int32{}
	slice1=append(slice3,1)
	fmt.Println(slice3)
	addV(slice3)
	fmt.Println(slice3)
	return

	//测试切片为更改值时，为引用类型
	var slice []int
	slice=[]int{}
	//slice[0]=1
	slice=append(slice,1)
	a4(slice)
	fmt.Println(slice[0])
	return

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
