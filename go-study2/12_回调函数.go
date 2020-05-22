package main

import "fmt"

type FuncType func(int, int) int

//实现加法
func add(a, b int) int {
	return a + b
}

//实现减法
func minus(a, b int) int {
	return a - b
}

//回调函数，函数参数是函数类型，这个函数就是回调函数
//计算器，可以进行四则运算
//多态，多种形态，调用同一个接口，不同的表现，可以实现不同表现，加减乘除
//先有想法，后面再实现功能
func calc(a, b int, fTest FuncType) (result int) {
	fmt.Println("cacl")
	result = fTest(a, b) //这个函数还没有实现
	return
}
func main() {
	a := calc(1, 1, add)
	b := calc(3, 1, minus)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
}
