package main

import "fmt"

//函数的返回值是一个匿名函数，返回一个函数类型
func test02() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
func main() {
	//返回值为一个匿名函数，返回一个函数类型，通过f来调用返回的匿名函数，f来调用闭包函数
	f := test02()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
func test01() int {
	var x int
	x++
	return x * x
}
func main01() {
	fmt.Println(test01())
	fmt.Println(test01())
	fmt.Println(test01())
	fmt.Println(test01())
}
