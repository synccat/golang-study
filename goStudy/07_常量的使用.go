package main

import "fmt"

func main() {
	//变量：程序运行期间可以改变的量，变量声明需要var
	//常量：程序运行期间不可以改变的量，常量的声明需要const
	const a int = 10
	//a = 20 //err,常量不允许修改
	fmt.Printf("a = %d\n", a)
	const b = 10 //不适用:=
	fmt.Printf("b = %d\n", b)
	fmt.Printf("b type is %T\n", b)

}
