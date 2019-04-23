package main

import (
	"fmt"
	"testing"
	"moqikaka.com/goutil/mathUtil"
)

func TestGetRandRangeInt(t *testing.T) {
	lower, upper := 1, 2
	rand := mathUtil.GetRand().GetRandRangeInt(lower, upper)
	fmt.Println("ran:",rand)
}
