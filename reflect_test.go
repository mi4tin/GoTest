package main

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

var offset uintptr = 0XFFFF

type RestFulHello struct {
}

func (ff *RestFulHello)Get(){

}

func TestReflect1(t *testing.T) {
h:=&RestFulHello{}
	fmt.Println( reflect.TypeOf(h).String())
	mh,_:=reflect.TypeOf(h).MethodByName("Get")
	fmt.Println(mh.Name)
}

func TestReflect(t *testing.T) {
	fmt.Println("offset", offset)
}

func intCX(d interface{}) int64 {
	return 0
	v := reflect.ValueOf(d).Elem()
	f := v.FieldByName("X")
	x := f.Int()
	x++
	f.SetInt(x)
	return x
}

var offset1 uintptr = 0xFFFF

func unsafeIncx(d interface{}) {

	if offset1 == 0XFFFF {
		v := reflect.TypeOf(d)
		f, _ := v.FieldByName("X")
		offset1 = f.Offset
		//offset1 = unsafe.Offsetof(f)
		fmt.Println("offset1", offset1)
	}
	//p := unsafe.Pointer(&d)
	p2 := unsafe.Pointer(&d)
	//fmt.Println("p:", p)
	fmt.Println("p2:", p2)
	//fmt.Println("pp", p)
	//fmt.Println("offset2", unsafe.Offsetof(d.X))
	px := (*int64)(unsafe.Pointer((*[2]uintptr)(p2)[1] + offset1))
	//	px := (*int64)(unsafe.Pointer(p2 + offset1))
	//fmt.Println("px", px)
	//*px++
	*px = 12
	fmt.Println("dyyy", d)
	//	return *px
}

type sss struct {
	X int
	Y int
}

func TestT(t *testing.T) {
	str := "hello world!"
	fmt.Printf("%v\n", str)
}

func TestIncx(b *testing.T) {
	d := struct {
		X int
	}{100}
	intCX(d)
}

func TestUnsafe(t *testing.T) {
	d := struct {
		X int
		Y int
		Z int
	}{100, 101, 102}
	p0 := unsafe.Pointer(&d)
	//fmt.Println("p:", p)
	fmt.Println("p0:", p0)
	unsafeIncx(d)
}

func BenchmarkIncx(b *testing.B) {
	d := struct {
		X int
	}{100}
	for i := 0; i < b.N; i++ {
		intCX(&d)
	}
}

func BenchmarkUnsafeIncx(b *testing.B) {
	d := struct {
		X int
		Y int
		Z int
	}{100, 101, 102}
	for i := 0; i < b.N; i++ {
		unsafeIncx(d)
	}
}
