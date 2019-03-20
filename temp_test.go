package main

import (
	"fmt"
	"moqikaka.com/goutil/mathUtil"
	"testing"
	"time"
)

func TestTemp1(t *testing.T) {
	//time.Duration(time.Second)
}

func TestTemp2(t *testing.T) {
	randObj := mathUtil.GetRand()
	rewardNum := randObj.GetRandRangeInt(2, 4)
	fmt.Println("ran", fmt.Sprintf("*%v", rewardNum))
}
