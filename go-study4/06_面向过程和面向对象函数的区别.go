package main

import (
	"fmt"
)

//实现2数相加
//面向过程
func Add01(a, b int) int {
	return a + b
}

//面向对象
type long int

func (temp long) add(a long) long {
	return temp + a
}

func main() {
	result := Add01(1, 1) //普通函数调用方式
	fmt.Println("result = ", result)
	var r long = 3
	r = r.add(5)
	fmt.Println("result = ", r)

}
