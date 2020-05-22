package main

import (
	"fmt"
)

func main() {
	//创建一个无缓存的channel
	ch := make(chan int)
	//新建一个gorountine
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		//不需要再写数据时，关闭channel
		close(ch)
	}()
	for {
		if num, ok := <-ch; ok {
			fmt.Println("num = ", num)
		} else {
			break
		}
	}
}
