package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second >> 1)
	}
	fmt.Printf("\n")
}

func main() {
	go func() {
		printer("hello")
		ch <- -999
	}()
	go func() {
		<-ch
		printer("world")
	}()
	time.Sleep(time.Second << 4)
}
