package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func minus(a, b int) int {
	return a - b
}

//函数也是一种数据类型，通过type给一个函数类型起名
//FuncType他是一个函数类型
type FuncType func(int, int) int //没有函数名字，没有{}

func main() {
	var result int
	result = add(1, 1)
	fmt.Println("result1 = ", result)
	var fTest FuncType
	fTest = add //是变量就可以赋值
	result = fTest(10, 20)
	fmt.Println("result2 = ", result)
	fTest = minus
	result = fTest(10, 20)
	fmt.Println("result3 = ", result)
}
