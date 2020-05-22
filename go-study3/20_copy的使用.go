package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := []int{11, 22, 33}
	copy(a, b)
	fmt.Println(a)
}
