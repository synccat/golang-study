package main

import (
	"fmt"
)

func MyFunc01(a int, b int) {

}

//...int，不定参数类型
//注意：不定参数，一定放在形参中的最后一个参数
func MyFunc02(args ...int) { //传递的参数可以是零个或多个
	fmt.Println("len(args) = ", len(args))
	for i := 0; i < len(args); i++ {
		fmt.Printf("args[%d] = %d\n", i, args[i])
	}
	for i, data := range args {
		fmt.Printf("args[%d] = %d\n", i, data)
	}
	for i := range args {
		fmt.Printf("args[%d]\n", i)
	}
	for data := range args {
		fmt.Printf("%d\n", data)
	}
}
func main01() {
	MyFunc02()
	MyFunc02(1, 2)
	MyFunc02(1, 2, 3)
}
func MyFunc03(a int, args ...int) {

}
func main() {
	MyFunc03(1, 32, 123, 12314)
}
