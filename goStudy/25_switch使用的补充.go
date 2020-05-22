package main

import "fmt"

func main() {
	//支持一个初始化语句，初始化语句和变量本身，以分号分隔
	switch num := 1; num { //switch后面写的是变量本身
	case 1:
		fmt.Println("按下的是1楼")
		fallthrough //不跳出switch语句，无条件继续执行后面的语句
	case 2:
		fmt.Println("按下的是2楼")
	case 3:
		fmt.Println("按下的是3楼")
	default:
		fmt.Println("按下的是XXXX楼")

	}

	score := 85
	switch { //可以没有条件
	case score > 90:
		fmt.Println("优秀")
	case score > 80:
		fmt.Println("良好")
	case score > 70:
		fmt.Println("普通")
	}
}
