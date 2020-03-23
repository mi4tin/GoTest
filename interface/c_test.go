package _interface

import (
	"fmt"
	"moqikaka.com/X1/PlatformHelper/timeHelper"
	"moqikaka.com/goutil/timeUtil"
	"testing"
)

func TestA(t *testing.T){
	fmt.Println(timeUtil.ToDateTimeString2(timeHelper.GetNowAddMinTime(10)))
}
