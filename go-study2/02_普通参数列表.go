package main

import (
	"fmt"
)

//有参无返回值函数的定义，普通参数列表
func MyFunc01(a int) {
	fmt.Println("a = ", a)
}
func MyFunc02(a int, b int) {
	fmt.Printf("a = %d, b=%d\n", a, b)
}
func MyFunc03(a, b int) {
	fmt.Printf("a = %d, b=%d\n", a, b)
}
func MyFunc04(a int, b string, c float64) {
	fmt.Printf("a = %d, b = %s, c = %f\n", a, b, c)
}
func main() {
	//有参无返回值函数调用：函数名（所需参数）
	MyFunc01(666)
	MyFunc02(6, 6)
	MyFunc03(7, 7)
	MyFunc04(11, "abc", 3.21)
}
