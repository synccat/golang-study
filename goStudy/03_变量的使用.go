package main

import "fmt" //导入包必须要使用
func main() {
	//变量，程序运行期间，可以改变的量
	//1、声明格式 var 变量名 类型
	//2、只是声明没有初始化的整形变量，默认值为0
	//3、在同一个{}里，声明的变量名是唯一的
	var a int
	fmt.Println("a=", a)
	//4、可以同时声明多个变量
	//var b, c int
	a = 10 //变量的赋值
	fmt.Println("a=", a)

	//2、变量的初始化，声明变量时，同时赋值
	var b int = 10
	fmt.Println("b=", b)
	b = 20
	fmt.Println("b=", b)

	//3、自动推导类型
	c := 30
	//%T打印变量所属类型
	fmt.Printf("c的type是 %T", c)
}
