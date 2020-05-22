package main

import "fmt"

func main() {
	var flag bool
	flag = true
	fmt.Printf("flag = %t\n", flag)
	//bool类型不能转换为int
	//	fmt.Printf("flag = %d\n", int(flag))

	//0就是假，非0就是真
	//整形也不能转换为bool

}
