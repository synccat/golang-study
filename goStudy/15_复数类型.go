package main

import "fmt"

func main() {
	var t complex128 //声明
	t = 2.1 + 3.14i  //赋值
	fmt.Println("t=", t)

	//自动推导类型
	t2 := 3.3 + 6.28i
	fmt.Println("t2=", t2)

	//通过内建函数，取实部和虚部
	fmt.Println("real(t2) = ", real(t2), ",imag(t2) = ", imag(t2))
}
