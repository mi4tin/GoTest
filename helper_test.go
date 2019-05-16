package main

import (
	"testing"
	"moqikaka.com/X1/PlatformHelper/timeHelper"
	"time"
	"fmt"
)

func TestWeek(t *testing.T){
	fmt.Println("week:",timeHelper.YearWeek(time.Now()))
}