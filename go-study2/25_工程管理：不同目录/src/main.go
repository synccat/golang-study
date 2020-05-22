package main

import (
	"calc"
	"fmt"
)

func main() {
	a := calc.Add(10, 20)
	fmt.Println("a = ", a)
	fmt.Println("a = ", calc.Minus(10, 20))
}
