package main

import (
	"fmt"
	"testing"
)

func testMap(m map[int]int) {
	for i := 0; i < 10000; i++ {
		m[i] = i
	}
}

func TestMapSpace(t *testing.T) {
	m := make(map[string]string, 1)
	m[""] = "2342"
	m["2"] = "a"
	m["3"] = "b"
	fmt.Println("mp:",m["2-a"],m["3-a"])
}

func BenchmarkCapMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m := make(map[int]int, 10000)
		b.StartTimer()
		testMap(m)
	}
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m := make(map[int]int)
		b.StartTimer()
		testMap(m)
	}
}

const n = 100

func TestCapMap(b *testing.T) {
	fmt.Println("TestCapMap")
	for i := 0; i < n; i++ {
		m := make(map[int]int, 100)
		testMap(m)
	}
	fmt.Println("TestCapMapend")
}

func TestMap(b *testing.T) {
	fmt.Println("TestMap")
	for i := 0; i < n; i++ {
		m := make(map[int]int)
		testMap(m)
	}
}

type Te struct {
	A string
}

const Width, Height = 640, 480

type Cursor struct {
	X, Y int
}

func Center(c *Cursor) {
	c.X += Width / 2
	c.Y += Height / 2
}

func CenterCursor() {

}

func array() [capacity]int {
	var d [capacity]int

	for i := 0; i < len(d); i++ {
		d[i] = 1
	}

	return d
}

func slice() []int {
	d := make([]int, capacity)

	for i := 0; i < len(d); i++ {
		d[i] = 1
	}

	return d
}

const capacity = 1024
