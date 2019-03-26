package main_test

import (
	"flag"
	"fmt"
	"log"
	"moqikaka.com/X1/PlatformHelper/jsonHelper"
	"moqikaka.com/X1/PlatformHelper/timeHelper"
	"moqikaka.com/goutil/typeUtil"
	"net/http"
	_ "net/http/pprof"
	"net/url"
	"strconv"
	"sync"
	"testing"
	"time"
)

func getSequence(i int) func() int {
	return func() int {
		i += 1
		return i
	}
}

func Counter(wg *sync.WaitGroup) {
	time.Sleep(time.Second)

	var counter int
	for i := 0; i < 1000000; i++ {
		time.Sleep(time.Millisecond * 200)
		counter++
	}
	wg.Done()
}

type name1 struct {
	int
	string
}

func (this *name1) String() string {
	return fmt.Sprintf("This is String(%s).", this.string)
}

type Car interface {
	NameGet() string
	Run(n int)
	Stop()
}

type BMW struct {
	Name string
}

func (this *BMW) NameGet() string {
	return this.Name
}
func (this *BMW) Run(n int) {
	fmt.Printf("BMW is running of num is %d \n", n)
}

func (this *BMW) Stop() {
	fmt.Printf("BMW is stop \n")
}

type Test11 struct {
	Name string
}

func TestA(t *testing.T) {
	int32V, errf := typeUtil.Int32("")
	fmt.Println("int32v:", int32V, errf)
	return
	fmt.Println("week", timeHelper.YearWeek(time.Now()))
	return
	ba := 1
	fmt.Println(float64(ba) * 0.01)
	return
	fmt.Println(jsonHelper.JsonString(Test11{Name: "<sdf"}))
	return
	t24 := 5
	fmt.Println(int32(t24))
	return
	var t23 string
	var x11 interface{}
	x11 = t23

	//f := x11.(int) //转成int
	//fmt.Println(f)
	f, ok := x11.(int) //转成int，不报错
	fmt.Println(f, ok)
	return
	var car Car
	fmt.Println(car) // <nil>

	var bmw BMW = BMW{Name: "宝马"}
	car = &bmw
	fmt.Println(car.NameGet()) //宝马
	car.Run(1)                 //BMW is running of num is 1
	car.Stop()                 //BMW is stop
	return
	n := new(name1)
	fmt.Println(n) //This is String().
	n.string = "suoning"
	d25 := fmt.Sprintf("%s", n) //This is String(suoning).
	fmt.Println(d25)
	return
	fmt.Println(fmt.Sprintf("%v", 234324))
	return
	urlObj, _ := url.Parse("http://mingame.sy1314.cn/?dfdsf")
	fmt.Println(urlObj.RawQuery)
	return
	//t11221 := time.Now()
	//t1122 := time.Date(t11221.Year(), t11221.Month(), t11221.Day(), 0, 0, 0, 0, time.Local)
	//intOut, _ := strconv.Atoi(strconv.FormatFloat(math.Floor(1.6), 'f', 0, 64))

	disDay := time.Now().Sub(time.Date(2018, 7, 30, 0, 0, 0, 0, time.Local)).Hours() / 24
	fmt.Println(disDay)

	intOut, _ := strconv.Atoi(strconv.FormatFloat(disDay, 'f', 0, 64)) //转成int
	fmt.Println(intOut)
	//fmt.Println(intOut)
	return
	var tf float64
	var tA int64
	tA = 206
	tf = float64(tA) * 0.01
	fmt.Println(tf)
	return
	chb := make(chan int)
	go func() {
		fmt.Println("01")
		<-chb
		fmt.Println("1")
	}()
	fmt.Println(0)
	time.Sleep(time.Second * 2)
	chb <- 1
	fmt.Println("2")

	return
	const (
		a = 1 << iota
		b
		c
		d
	)
	const (
		c1 = 1 << iota
		c2
	)
	fmt.Println(a, b, c, d)
	fmt.Println(c1, c2)
	return
	flag.Parse()

	fmt.Println("启动")
	//远程获取pprof数据
	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Counter(&wg)
	}
	wg.Wait()

	// sleep 10mins, 在程序退出之前可以查看性能参数.
	time.Sleep(60 * time.Second)
}
