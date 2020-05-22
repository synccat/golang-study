package main

import "fmt"

func main() {
	defer fmt.Println("bbbbbb")
	fmt.Println("aaaaaa")
	//defer延迟调用，main函数结束前调用
}
