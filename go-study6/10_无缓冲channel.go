package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个无缓存的channel
	ch := make(chan int)
	//新建协程
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("%s\n", "子协程")
			ch <- i
		}
	}()
	//延时
	time.Sleep(time.Second << 1)
	for i := 0; i < 3; i++ {
		num := <-ch
		fmt.Printf("num = %d\n", num)
	}
}
