package main

import "fmt"

func main() {
	var num int
	fmt.Printf("请按下楼层：")
	fmt.Scan(&num)
	switch num { //switch后面写的是变量本身
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
	//	num := 111
	//	switch num { //switch后面写的是变量本身
	//	case 1:
	//		fmt.Println("按下的是1楼")
	//		break //go语言保留了break关键字，跳出switch语言，不写默认包含
	//	case 2:
	//		fmt.Println("按下的是2楼")
	//	case 3:
	//		fmt.Println("按下的是3楼")
	//	default:
	//		fmt.Println("按下的是XXXX楼")

	//	}
}
