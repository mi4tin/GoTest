package main

import (
	"encoding/json"
	"fmt"
	"github.com/json-iterator/go"
	"testing"
)

type jsonT struct {
	A string
	B float64
	C string
	D string
	F string
}

func TestJson(t *testing.T) {
	s := `{"A":"sdfsfdsfdsfsdfsdfdsfdsfdsfdsfsdfdsfdsddddddddddddddddddddd","B":2}`
	jsont := &jsonT{}
	err := json.Unmarshal([]byte(s), jsont)
	fmt.Println(err)
	fmt.Println(jsont)
}

func BenchmarkGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := `{"A":"sdfsfdsfdsfsdfsdfdsfdsfdsfdsfsdfdsfdsddddddddddddddddddddd","B":2}`
		jsont := &jsonT{}
		json.Unmarshal([]byte(s), jsont)
		//if i == b.N-1 {
		//	fmt.Println(err)
		//	fmt.Println(jsont)
		//}
	}
}

func BenchmarkJsoniter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := `{"A":"sdfsfdsfdsfsdfsdfdsfdsfdsfdsfsdfdsfdsddddddddddddddddddddd","B":2}`
		jsont := &jsonT{}
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		json.Unmarshal([]byte(s), &jsont)
		//if i == b.N-1 {
		//	fmt.Println(err)
		//	fmt.Println(jsont)
		//}

	}
}
