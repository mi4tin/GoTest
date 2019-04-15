package main

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	var i *int
	i = new(int)

	fmt.Println("i:", *i)
}
