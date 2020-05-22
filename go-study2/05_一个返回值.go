package main

import "fmt"

//无参有返回值，只有一个返回值
func myFunc01() int {
	return 666
}

//给返回值起一个变量名，go推荐写法
func myFunc02() (result int) {
	result = 888
	return
}
func main() {
	fmt.Println(myFunc01())
	fmt.Println(myFunc02())
}
