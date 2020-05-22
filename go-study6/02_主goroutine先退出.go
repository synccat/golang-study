package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		i := 0
		for {
			fmt.Println("child i = ", i)
			i++
			time.Sleep(time.Second << 1)
		}
	}()
	// i := 0
	// for {
	// 	fmt.Println("main i = ", i)
	// 	i++
	// 	time.Sleep(time.Second << 1)
	// 	if i == 3 {
	// 		break
	// 	}
	// }
}
