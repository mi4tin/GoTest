package main

import (
	//"fmt"
	"fmt"
	"moqikaka.com/goutil/redisUtil"
	"testing"
	"time"
)

//var f = 1

var redisObj *redisUtil.RedisPool

func init() {
	//f = 2

	fmt.Println("redisInit")

	var err error
	redisConfig, err := redisUtil.NewRedisConfig("ConnectionString=127.0.0.1:6379;Password=123456;Database=4;MaxActive=50;MaxIdle=20;IdleTimeout=300;DialConnectTimeout=10")
	if err != nil {
		panic(fmt.Errorf("redis连接字符串格式有误，错误信息为：%s", err))
	}

	redisObj = redisUtil.NewRedisPool2("testRedis", redisConfig)

	if err := redisObj.TestConnection(); err != nil {
		//panic(fmt.Errorf("redis连接失败：%s", err))
		fmt.Println("redis连接失败")
	} else {
		fmt.Println("redis连接成功")
	}

}

func Benchmark_SetNx(t *testing.B) {
	return
	for i := 0; i < t.N; i++ {
		r, err := SetDetail(redisObj, "test5", time.Now().Unix(), redisUtil.Expire_Seond, 100, redisUtil.Set_NX)
		if err != nil {
			fmt.Println(err)
		}
		if r != nil {
			fmt.Println("re", r)
		}
		//fmt.Println("setnx", i)
	}
}

func TestSexNx(t *testing.T) {
	key := "test3"
	r, err := SetDetail(redisObj, key, time.Now().Unix(), redisUtil.Expire_Seond, 100, redisUtil.Set_NX)
	if err != nil {
		fmt.Println(err)
	}
	if r != nil {
		fmt.Println("aaa", r)
	}
	fmt.Println("SexNx1")
	r, err = SetDetail(redisObj, key, time.Now().Unix(), redisUtil.Expire_Seond, 100, redisUtil.Set_NX)
	if err != nil {
		fmt.Println(err)
	}
	if r != nil {
		fmt.Println("aaa", r)
	}
	fmt.Println("SexNx2")
}

// 设置key和对应的value
// key:key
// value:value
// expireVal:超时时间值
// expireType:设置的超时类型
// setType:值存储类型
// 返回值:
// 错误对象
func SetDetail(this *redisUtil.RedisPool, key string, value interface{}, expireType redisUtil.ExpireType, expireVal int, setType redisUtil.SetType) (r interface{}, err error) {
	conn := this.GetConnection()
	defer conn.Close()

	paramList := []interface{}{key, value}

	// 超时设置组装
	if expireType != redisUtil.Expire_NoExpire {
		paramList = append(paramList, string(expireType), expireVal)
	}

	// 设置组装
	if setType != redisUtil.Set_Write {
		paramList = append(paramList, string(setType))
	}

	r, err = conn.Do("SET", paramList...)
	//fmt.Println("detai", r)

	return
}
