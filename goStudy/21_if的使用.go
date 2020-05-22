package main

import "fmt"

func main() {
	s := "王思聪"
	if s == "王思聪" { //左括号和if在同一行
		fmt.Println("左手一个妹子，右手一个妹子")
	} else {
		fmt.Println("撸代码")
	}

	//if支持1个初始化语句,初始化语句和判断条件以";"分隔
	if a := 10; a == 10 {
		fmt.Println("a == 10")
	}

}
