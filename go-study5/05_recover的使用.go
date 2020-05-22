package main

import (
	"fmt"
)

func funca(num int) {
	var a [10]int
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	a[num] = 10
}
func funcb() {
	fmt.Println("bbbbbbbbbb")
}
func funcc() {
	fmt.Println("ccccccccc")
}
func main() {
	funca(20)
	funcb()
	funcc()
}
