package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个定时器，设置时间为2s，2s后，往time通道写内容（当前时间）
	timer := time.NewTimer(time.Second << 1)
	fmt.Println(time.Now())
	<-timer.C
	fmt.Println(time.Now())

}
