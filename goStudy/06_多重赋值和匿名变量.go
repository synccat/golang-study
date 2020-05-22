package main

import "fmt"

//go函数可以返回多个值
func test() (a, b, c int) {
	return 1, 2, 3
}

func main() {
	a, b := 10, 20
	//交换2个变量的值
	var tmp int
	tmp = a
	a = b
	b = tmp
	fmt.Printf("a = %d, b= %d\n", a, b)

	a, b = b, a
	fmt.Printf("a = %d, b= %d\n", a, b)

	//_匿名变量，丢弃数据不处理，_匿名变量配合函数返回值使用，才有优势
	var d int
	_, d, _ = test()
	fmt.Printf("d = %d\n", d)
}
