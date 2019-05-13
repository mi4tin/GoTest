package main

import (
	"fmt"
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	var envValue string
	os.Setenv("wshenv","369")
	envValue = os.Getenv("wshenv")//windows下环境变量更新，要重启电脑才生效。linux只在当前终端生效
	fmt.Println(envValue)
}
