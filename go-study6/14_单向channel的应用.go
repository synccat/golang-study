package main

import (
	"fmt"
)

func producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i * i
	}
	close(ch)
}
func comsumer(ch <-chan int) {
	for num := range ch {
		fmt.Printf("chnum = %d\n", num)
	}
}
func main() {
	//创建一个双向通道
	ch := make(chan int)
	//生产者，生产数组，写入channel
	go producer(ch)
	//消费者，从channel读取内容，打印
	comsumer(ch)
}
