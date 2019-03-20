package main

import (
	"fmt"
	"moqikaka.com/goutil/logUtil"
	"os/user"
	"sync"
	"testing"
)

var (
	wg sync.WaitGroup
)

func TestGetUnixHome(t *testing.T) {
	wg.Add(1)
	us, err := user.Current()
	if err != nil {
		fmt.Println(err)
		return
	}
	logUtil.DebugLog("home:", us.HomeDir)
	fmt.Println(us.HomeDir)
	wg.Wait()
}
