package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	a1 := a[3]
	fmt.Println(a1)
	a2 := a[2:5:6]
	fmt.Printf("%d:%d\n", len(a2), cap(a2))
	a3 := a[:6]
	fmt.Printf("%d:%d\n", len(a3), cap(a3))
	a4 := a[3:6]
	fmt.Printf("%d:%d\n", len(a4), cap(a4))
}
