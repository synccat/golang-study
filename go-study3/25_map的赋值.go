package main

import (
	"fmt"
)

func main() {
	var m1 map[int]string
	fmt.Println("m1 = ", m1)
	fmt.Println("len = ", len(m1))
	m2 := make(map[int]string)
	fmt.Println("m2 = ", m2)
	fmt.Println("len = ", len(m2))
	m3 := make(map[int]string, 2)
	fmt.Println("m3 = ", m3)
	fmt.Println("len = ", len(m3))
	m3[1] = "adfsa"
	m3[2] = "2321"
	m3[3] = "3fda"
	fmt.Println("m3 = ", m3)
	fmt.Println("len = ", len(m3))
	

}
