package main

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestDefer(t *testing.T) {
	x := exa8()
	fmt.Println("TestDefer:", *x)
}

func TestPointer1(t *testing.T) {
	a := [4]int{0, 1, 2, 3}
	p1 := unsafe.Pointer(&a[1])
	p3 := unsafe.Pointer(uintptr(p1) + 2*unsafe.Sizeof(a[0]))
	*(*int)(p3) = 6
	fmt.Println("a =", a) // a = [0 1 2 6]

}

func exa8() (res *int) {
	var i = 1
	res = &i
	defer func(res *int) {
		*res = *res + 3
	}(res)
	return
}

func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f1() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f2() (r int) {
	defer func(x int) {
		x = x + 5
	}(r)
	return 1
}
func test4() (s [4]int) {
	s = [4]int{1, 2, 3, 4}
	defer func(s [4]int) {
		s[0] = 4
	}(s)
	defer func() {
		s[0] = 3
	}()
	s[0] = 2
	return s
}

func test5() (s interface{}) {
	s = []int{1, 2, 3, 4}
	defer func(s interface{}) {
		s.([]int)[0] = 4
	}(s)
	defer func() {
		s.([]int)[0] = 3
	}()
	s.([]int)[0] = 2
	return s
}

func test6() (s interface{}) {
	s = 1
	defer func(s interface{}) {
		s = 4
	}(s)
	defer func() {
		s = 3
	}()
	s = 2
	return s
}

func f3() (s []*int) {
	f := 1
	s = []*int{&f}
	defer func(s []*int) {
		f1 := 2
		s[0] = &f1
	}(s)
	defer func() {
		f2 := 3
		s[0] = &f2
	}()
	f = 4
	s[0] = &f
	return s
}

func TestIfelse(t *testing.T) {
	return
	var f64 float64
	f64 = 1 * 0.01
	fmt.Println(f64)
	return
	a := 0
	b := 2
	defer func() {
		fmt.Println(111)
	}()
	if a == 0 {

		defer func() {
			if b == 3 {
				fmt.Println(123)
			}
		}()
	}
	b = 3
	fmt.Println(234)
}

func TestIfelse1(t *testing.T) {
	return
	var a = 1
	if a == 1 {
		b := 2
		defer func() {
			if b == 2 {
				fmt.Println("bbb", b)
			}
		}()
		//b = 3
	}
}
