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



func TestMapLen(t *testing.T) {
	maps:=make(map[int32]string,7)
	maps[1]="a"
	maps[2]="b"
	maps[3]="c"
	maps[4]="d"
	maps[5]="e"
	maps[6]="f"

	 //x:= [len]int32{1,2,3}
	 var y []int32
	 y=append(y,1)
	 fmt.Println("wsh:",y)
}


func BenchmarkASArray(b *testing.B) {
	maps:=make(map[int32]string,7)
	maps[1]="a"
	maps[2]="b"
	maps[3]="c"
	maps[4]="d"
	maps[5]="e"
	maps[6]="f"

	len:=len(maps)
	fmt.Println("map.len:",len)
	var x [7]int32
	for i := 0; i < b.N; i++ {
		index:=0
		for k,_:=range maps{
			x[index]=k
			index++
		}

		//_ = array()
	}
}

func BenchmarkASSlice(b *testing.B) {
	maps:=make(map[int32]string,10)
	maps[1]="a"
	maps[2]="b"
	maps[3]="c"
	maps[4]="d"
	maps[5]="e"
	maps[6]="f"

	for i := 0; i < b.N; i++ {
		var x []int32
		for k,_:=range maps{
			x=append(x,k)
		}
		//_ = array()
	}

}


func BenchmarkASSlice1(b *testing.B) {
	maps:=make(map[int32]string,10)
	maps[1]="a"
	maps[2]="b"
	maps[3]="c"
	maps[4]="d"
	maps[5]="e"
	maps[6]="f"


	for i := 0; i < b.N; i++ {
		len:=len(maps)
		x:=make([]int32,len)
		for k,_:=range maps{
			x=append(x,k)
		}
		//_ = array()
	}


}


func BenchmarkASSlice2(b *testing.B) {
	maps:=make(map[int32]string,10)
	maps[1]="a"
	maps[2]="b"
	maps[3]="c"
	maps[4]="d"
	maps[5]="e"
	maps[6]="f"


	for i := 0; i < b.N; i++ {
		len:=len(maps)
		x:=make([]int32,len,len)
		for k,_:=range maps{
			x=append(x,k)
		}
		//_ = array()
	}

}