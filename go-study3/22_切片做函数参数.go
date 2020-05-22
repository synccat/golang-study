package main

import (
	"fmt"
	"math/rand"
	"time"
)

func initData(s []int) {
	//设置种子
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100) //100以内的随机数
	}
}
func bubbleSort(s []int) {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}
func main() {
	var n int = 10
	//创建切片，len为n
	s := make([]int, n)
	initData(s) //初始化数组
	fmt.Println("排序前：", s)
	bubbleSort(s)
	fmt.Println("排序后：", s)
}
