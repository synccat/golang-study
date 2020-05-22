package main

import "fmt"

func maxAndMin(a, b int) (max, min int) {
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	return
}
func main() {
	max, min := maxAndMin(10, 20)
	fmt.Printf("max = %d, min = %d\n", max, min)
}
