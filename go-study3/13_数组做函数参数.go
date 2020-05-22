package main

import (
	"fmt"
)

//数组做函数参数，是值传递
func modify(a [5]int) {
	a[0] = 666
	fmt.Println("modify a = ", a)
}
func main() {
	arrayA := []int{3, 213, 4, 2, 13, 4, 99, 5, 6} //初始化
	modify(a)                                      //数组传递过去
	fmt.Println("main: a=", a)
}
