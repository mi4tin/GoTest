package main

import (
	"fmt"
	"runtime"
	"testing"
)

type X struct {
	data [1 << 20]byte
	ptr  *X
}

func TestGc1(t *testing.T) {
	var a, b X
	a.ptr = &b
	b.ptr = &a
	runtime.SetFinalizer(&a, func(*X) { println("b fina.") })

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)

	fmt.Printf("sys:%d Kb\n", m.Sys/1024)

}
