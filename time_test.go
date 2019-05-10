package main

import (
	"testing"
	"time"
	"fmt"
	"moqikaka.com/goutil/logUtil"
)

func TestCostTime(t *testing.T){
	start := time.Now()
	time.Sleep(time.Second*2)
	cost := time.Since(start)
	fmt.Println(cost)
	logUtil.DebugLog("cost:",cost)
}