package main

import (
	"fmt"
)

//无参无返回值函数的定义
func MyFunc() {
	a := 666
	fmt.Println("a = ", a)

}
func main() {
	//无参无返回值函数的调用： 函数名()
	MyFunc()
}
