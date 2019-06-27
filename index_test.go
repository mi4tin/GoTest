package main

import (
	"testing"
	"fmt"
	"strings"
)

func BenchmarkT1(b *testing.B) {
	s:="117.139.247.210,117.139.247.212,117.139.247.212,117.139.247.212,117.139.247.212,117.139.247.212,117.139.247.212,117.139.247.212,117.139.247.212,117.139.247.215,"
	for i:=0;i<b.N;i++{
		strings.Index(s,"117.139.247.215,")
	}
}

func TestT1(t *testing.T) {
	s:="117.139.247.210,117.139.247.212"
	fmt.Println(strings.Index(s,"117.139.247.214"))
}