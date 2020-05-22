package main

import "fmt"

func swap(a, b int) {
	a, b = b, a
	fmt.Printf("swap: a=%d,b=%d\n", a, b)
}

func main() {
	a, b := 10, 20
	//通过一个函数交换a和b的内容
	swap(a, b)
	fmt.Printf("main: a=%d,b=%d\n", a, b)

}
