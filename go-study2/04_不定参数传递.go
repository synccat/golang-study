package main

import (
	"fmt"
)

func myFunc01(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("data = ", data)
	}
}
func myFunc02(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("data = ", data)
	}
}
func test(args ...int) {
	//全部元素传递给myFunc01
	//	myFunc01(args...)
	//只想把后2个参数传递给另外一个函数使用
	myFunc02(args[2:]...) //从下标2开始，把后面所有元素传递过去
	myFunc02(args[:2]...) //前2个
}

func main() {
	test(1, 2, 3, 4)
}
