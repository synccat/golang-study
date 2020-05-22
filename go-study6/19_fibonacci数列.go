package main

import (
	"fmt"
)

func producer(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println("flag = ", flag)
			return
		}

	}
}
func comsumer(ch <-chan int, quit chan<- bool) {
	for i := 0; i < 100; i++ {
		fmt.Println(<-ch)
	}
	quit <- true

}
func main() {
	ch := make(chan int)    //数字通信
	quit := make(chan bool) //程序是否结束
	//生产者
	go producer(ch, quit)
	//消费者
	comsumer(ch, quit)
}
