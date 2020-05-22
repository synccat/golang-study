package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个定时器，设置时间为2s，2s后，往time通道写内容（当前时间）
	fmt.Println("开始了")
	timer := time.NewTimer(time.Second)
	timer.Reset(time.Second >> 2)
	<-timer.C
	fmt.Println("时间到")

}
