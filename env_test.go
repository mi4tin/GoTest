package main

import (
	"fmt"
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	var envValue string
	envValue = os.Getenv("GOPATH")
	fmt.Println(envValue)
}
