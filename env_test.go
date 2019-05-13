package main

import (
	"fmt"
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	var envValue string
	envValue = os.Getenv("wshenv")//windows下环境变量更新，要重启电脑才生效。linux可即时更新
	fmt.Println(envValue)
}
