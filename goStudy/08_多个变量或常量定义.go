package main

import "fmt"

func main() {
	//不同类型变量的声明（定义）
	var a int
	var b float64

	a, b = 10, 3.14
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	var (
		c int
		d float64
	)
	c, d = 20, 6.28
	fmt.Println("c=", c)
	fmt.Println("d=", d)
	const (
		i = 10
		j = 3.14
	)
	fmt.Println("i=", i)
	fmt.Println("j=", j)
}
