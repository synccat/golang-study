package main

import "fmt"
import "os"

func main() {
	//接收用户传递的参数，字符串的形式
	list := os.Args
	n := len(list)
	fmt.Println("n = ", n)

	for i := 0; i < len(list); i++ {
		fmt.Printf("list[%d] = %s\n", i, list[i])
	}
	for i, data := range list {
		fmt.Printf("list[%d] = %s\n", i, data)
	}

}
