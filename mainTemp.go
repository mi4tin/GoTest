package main

import "fmt"

func main() {
//	fmt.Println("pre:",runtime.GOMAXPROCS(5))
	go func() {
		fmt.Println("123")
		fmt.Println("456")
	}()

	go func() {
		for  {

		}
	}()

	select {

	}

	return
}