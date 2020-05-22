package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.GOMAXPROCS(4)
	fmt.Print(n)
	for {
		go fmt.Print(3)
		go fmt.Print(2)
		go fmt.Print(1)
		fmt.Print(0)

	}
}
