package main

import "fmt"

func main() {
	var a int = 10
	//每个变量有两层含义：变量的内存，变量的地址
	fmt.Printf("a = %d\n", a)
	fmt.Printf("a的地址 = %v\n", &a)

	//保存某个变量的地址，需要指针类型   *int 保存int的地址，**int 保存 *int地址
	var p *int
	p = &a
	fmt.Printf("p = %v, &a = %v\n", *p, &a)
	fmt.Printf("*p = %v, a = %v\n", *p, a)
}
