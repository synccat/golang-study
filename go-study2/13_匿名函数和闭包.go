package main

import "fmt"

func main() {
	a := 10
	str := "mike"
	//匿名函数，没有函数名字,函数定义，还没有调用
	f1 := func() {
		fmt.Println("a = ", a)
		fmt.Println("str = ", str)
	}
	f1()

	//给一个函数类型起别名
	type FuncType func() //函数没有参数，没有返回值
	//声明变量
	var f2 FuncType
	f2 = f1
	f2()

	//定义匿名函数，同时调用
	func() {
		fmt.Printf("a = %d, str = %s\n", a, str)
	}() //后面的圆括号代表调用此匿名函数
	//带参数的匿名函数
	func(i, j int) {
		fmt.Printf("i = %d, j = %d\n", i, j)

	}(1, 2)
	//匿名函数，有参有返回值
	x, y := func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(10, 20)
	fmt.Printf("max = %d,min = %d", x, y)
}
