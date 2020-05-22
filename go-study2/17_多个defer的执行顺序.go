package main

import "fmt"

func test(x int) {
	result := 100 / x
	fmt.Println(result)
}
func main() {
	defer fmt.Println("aaaaaa")
	defer fmt.Println("bbbbbb")
	//调用一个函数，导致内存出问题
	defer test(0)
	defer fmt.Println("cccccc")
	//defer延迟调用，main函数结束前调用
}
