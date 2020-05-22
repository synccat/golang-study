package main

import "fmt"

func test(a int) {
	if a == 1 {
		fmt.Println("a = ", a)
		return
	}
	//函数调用自身
	test(a - 1)
	fmt.Println("a = ", a)
	return
}
func main() {
	test(3)
	fmt.Println("main")
}
