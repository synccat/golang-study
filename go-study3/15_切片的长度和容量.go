package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 0, 0}
	s := a[1:3:4]
	fmt.Println("s = ", s)
	fmt.Println("len(s) = ", len(s))
	fmt.Println("cap(s) = ", cap(s))
}
