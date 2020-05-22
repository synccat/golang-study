package main

import (
	"fmt"
)

//数组做函数参数，是值传递
func modify(p *[5]int) {
	(*p)[0] = 666
	fmt.Println("modify a = ", *p)
}
func main() {
	a := [5]int{3, 213, 4, 2, 13} //初始化
	modify(&a)                    //数组传递过去
	fmt.Println("main: a=", a)
}
