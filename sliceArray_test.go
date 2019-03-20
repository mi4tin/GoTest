package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	d := make([]int, 3)
	d[0] = 0
	d[1] = 1
	d[2] = 2
	d = append(d, 3)
	d = append(d, 4)
	d = append(d, 5)
	d = append(d, 6)
	fmt.Println("sliceTest%v", d, cap(d))
}

func TestArray(t *testing.T) {
	var d []int
	d = append(d, 1)
	d[0] = 2
	//d[1] = 3
	fmt.Println(d)
}

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//_ = array()
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//_ = slice()
	}

}
