package main

import "fmt"

func main() {
	//1
	//	a := 10
	//	if a == 10 {
	//		fmt.Println("a==10")
	//	} else {
	//		fmt.Println("a!=10")
	//	}

	//2
	if a := 10; a == 10 {
		fmt.Println("a==10")
	} else if a > 10 {
		fmt.Println("a>10")
	} else if a < 10 {
		fmt.Println("a<10")
	} else {
		fmt.Println("这是不可能的")
	}
}
