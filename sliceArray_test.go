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

	nums := []int{21, 31, 14}
	sum := 0
	for num := range nums {
		//如果将for后面的_去掉，结果是6
		//这是因为 for _ 表示遍历数组的下标，从nums[0]，nums[1]，nums[2]依次开始遍历，
		//所以最后的值为sum=2+3+4=9；但是如果把 for _ 去掉，就变成了遍历0 1 2，所以sum=0+1+2=3。
		//sum += num
		fmt.Println("TestSlice:",num)
	}
	fmt.Println("TestSlice:",sum)
}


func TestSlice1(t *testing.T) {
	m := []int32{}
//	m=append(m,1)
	//m[0] =
	//m[1] = "2342"
	fmt.Println(m)
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