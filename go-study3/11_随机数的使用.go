package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//设置种子
	rand.Seed(time.Now().UnixNano())
	//产生随机数
	for i := 0; i < 4; i++ {
		fmt.Println("rand = ", rand.Int())
		fmt.Println("rand = ", rand.Intn(100))

	}
}
