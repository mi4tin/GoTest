package main

import (
	"fmt"
	"testing"
	"moqikaka.com/X1/PlatformHelper/jsonHelper"
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


func TestSlice1(t *testing.T) {
	m := make([]string, 1)
	m[0] = "2342"
	//m[1] = "2342"
	fmt.Println(jsonHelper.JsonString(m))
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
