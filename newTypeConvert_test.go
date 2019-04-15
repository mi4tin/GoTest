package main

import (
	"encoding/json"
	"fmt"
	"moqikaka.com/goutil/mathUtil"
	"testing"
)

func TestNew3(t *testing.T) {
	return
	out := make(chan int, 1)
	fmt.Println("TestNew3.start")
	out <- 2
	fmt.Println("testnew3")
	fmt.Println(len(out))
	out <- 3
	fmt.Println("testnew3.re")
}

func TestNew2(t *testing.T) {
	return
	fmt.Println(1<<3 - 1)
	return
	rint := mathUtil.GetRand().GetRandRangeInt(3, 2)
	fmt.Println(rint)
}

func TestNew1(t *testing.T) {
	return
	f := 2.3
	fmt.Println("new1", int(f))
}

func TestJsonUnmarshal(t *testing.T) {
	return
	var result map[int32][]int32
	err := json.Unmarshal([]byte(`{"1":[2],"10":[1]}`), &result)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(result, "res")
}

func TestTime(t *testing.T) {
	return
	a := 100
	b := 200
	count := 0
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			count++
		}
	}
	fmt.Println("count:", count)
}
