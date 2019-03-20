package main

import (
	"fmt"
	"strings"
	"testing"
	"unsafe"
)

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

var s = strings.Repeat("a", 1024)

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func TestRepat(t *testing.T) {
	return
	fmt.Println(s)
}

func test() {
	b := []byte(s)
	_ = string(b)
}

func test2() {
	b := str2bytes(s)
	_ = bytes2str(b)
}

func BenchmarkByteStrTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}

func BenchmarkByteStrTestBlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test2()
	}
}
