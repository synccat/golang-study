package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {

		for i := 0; i < 5; i++ {
			fmt.Println("go = ", i)
		}
	}()

	for i := 0; i < 3; i++ {
		fmt.Println("hello = ", i)
	}
	runtime.Gosched()
}
