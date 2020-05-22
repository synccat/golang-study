package main

import (
	"fmt"
	"time"
)

func main() {
	tiker := time.NewTicker(time.Second >> 1)
	var i int = 0
	for {
		if i == 5 {
			tiker.Stop()
			fmt.Println("执行完成")
			break
		}
		<-tiker.C
		i++
		fmt.Println("又执行了一次")

	}
}
