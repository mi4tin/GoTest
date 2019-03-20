package main

import (
	"fmt"
	"testing"
	"unsafe"
)

const capacity1 = 50000

func TestP(t *testing.T) {
	//var a = 3.141596
	//var p = &a
	//var pp = &p
	//fmt.Println("a = ", a)
	//fmt.Println("chong address a = ", &a)
	//fmt.Println("p = ", p)
	//fmt.Println("chongchong p = ", &p)
	//fmt.Println("pp = ", pp)
	//fmt.Println("chongchong *pp = ", *pp)
	//fmt.Println("chongchong **pp = ", **pp)

	var a = 3014
	var p1 = &a
	var p2 = &a
	if p1 == p2 {
		fmt.Printf("chongchong p1(%x) = p2(%x) value(%d): \n", p1, p2, *p1)
	}
}

func TestUnsafe1(t *testing.T) {
	return
	var n int64 = 5
	var pn = &n
	var x = 0xc04204a160
	fmt.Println("x", x)
	fmt.Println("pn", unsafe.Pointer(pn), &n)
	var pf = (*float64)(unsafe.Pointer(pn))
	fmt.Println("pf", pf)
	fmt.Println(fmt.Sprintf("pfqq:%v", *pf))
	// now, pn and pf are pointing at the same memory address
	fmt.Println(*pf) // 2.5e-323
	*pf = 3.14159
	fmt.Println(n) // 4614256650576692846

}

func TestUnsafe2(t *testing.T) {
	a := [4]int{0, 1, 2, 3}
	fmt.Println("a.adr:", &a, &a[0], &a[1], unsafe.Sizeof(unsafe.Pointer(&a[2])))
	p1 := unsafe.Pointer(&a[0])
	p3 := unsafe.Pointer(uintptr(p1) + unsafe.Sizeof(a[0]))
	*(*int)(p3) = 6
	fmt.Println("a =", a) // a = [0 1 2 6]

	// ...

	type Person struct {
		name   string
		age    int
		gender bool
	}

	who := Person{"John", 30, true}
	pp := unsafe.Pointer(&who)
	fmt.Println("ppx:", uintptr(pp), pp)
	fmt.Println("age:", unsafe.Offsetof(who.name))
	pname := (*string)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.name)))
	page := (*int)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.age)))
	pgender := (*bool)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.gender)))
	*pname = "Alice"
	*page = 28
	*pgender = false
	fmt.Println(who) // {Alice 28 false}
}

// case A: conversions between unsafe.Pointer and uintptr
//         don't appear in the same expression
func TestIllegalUseA(t *testing.T) {
	fmt.Println("===================== illegalUseA")

	pa := new([4]int)
	//paf:=&([4]int)

	// split the legal use
	// p1 := unsafe.Pointer(uintptr(unsafe.Pointer(pa)) + unsafe.Sizeof(pa[0]))
	// into two expressions (illegal use):

	//不非法
	//ptr := unsafe.Pointer(pa)
	//p1 := unsafe.Pointer(uintptr(ptr) + unsafe.Sizeof(pa[0]))

	ptr := uintptr(unsafe.Pointer(pa))
	p1 := unsafe.Pointer(ptr + unsafe.Sizeof(pa[0]))

	// "go vet" will make a warning for the above line:
	// possible misuse of unsafe.Pointer

	// the unsafe package docs, https://golang.org/pkg/unsafe/#Pointer,
	// thinks above splitting is illegal.
	// but the current Go compiler and runtime (1.7.3) can't detect
	// this illegal use.
	// however, to make your program run well for later Go versions,
	// it is best to comply with the unsafe package docs.

	*(*int)(p1) = 123
	fmt.Println("*(*int)(p1)  :", *(*int)(p1)) //
}

// case B: pointers are pointing at unknown addresses
func TestIllegalUseB(t *testing.T) {
	fmt.Println("===================== illegalUseB")

	a := [4]int{2, 1, 2, 3}
	p := unsafe.Pointer(&a)
	p = unsafe.Pointer(uintptr(p) + uintptr(len(a))*unsafe.Sizeof(a[0]))
	// now p is pointing at the end of the memory occupied by value a.
	// up to now, although p is invalid, it is no problem.
	// but it is illegal if we modify the value pointed by p
	*(*int)(p) = 123
	fmt.Println("*(*int)(p)  :", *(*int)(p)) // 123 or not 123
	// the current Go compiler/runtime (1.7.3) and "go vet"
	// will not detect the illegal use here.

	// however, the current Go runtime (1.7.3) will
	// detect the illegal use and panic for the below code.
	p = unsafe.Pointer(&a)
	for i := 0; i < len(a); i++ {
		*(*int)(p) = 123 // Go runtime (1.7.3) never panic here in the tests

		fmt.Println("index", i)
		fmt.Println(i, ":", *(*int)(p))
		// panic at the above line for the last iteration, when i==4.
		// runtime error: invalid memory address or nil pointer dereference

		//fmt.Println("indexf", i, uintptr(p))
		fmt.Println(i, ":", *(*int)(p))
		p = unsafe.Pointer(uintptr(p) + unsafe.Sizeof(a[0]))
	}
}

func TestUnsafePointer(t *testing.T) {
	type MyInt int

	a := []MyInt{0, 1, 2}
	// b := ([]int)(a) // error: cannot convert a (type []MyInt) to type []int
	b := *(*[]int)(unsafe.Pointer(&a))

	b[0] = 3

	fmt.Println("a =", a) // a = [3 1 2]
	fmt.Println("b =", b) // b = [3 1 2]

	a[2] = 9

	fmt.Println("a =", a) // a = [3 1 9]
	fmt.Println("b =", b) // b = [3 1 9]
}

func value() interface{} {
	m := make(map[int]int, capacity1)
	for i := 0; i < capacity1; i++ {
		m[i] = i
	}
	return m
}

func pointer() interface{} {
	m := make(map[int]*int, capacity1)
	for i := 0; i < capacity1; i++ {
		v := i
		m[i] = &v
	}
	return m
}

func BenchmarkVPValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]int, capacity1)
		for i := 0; i < capacity1; i++ {
			m[i] = i
		}
	}
}

func BenchmarkVPPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m := make(map[int]*int, capacity1)
		for i := 0; i < capacity1; i++ {
			b.StopTimer()
			v := i
			b.StartTimer()
			m[i] = &v
		}
	}
}
