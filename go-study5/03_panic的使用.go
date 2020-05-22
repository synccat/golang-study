package main

import (
	"fmt"
)

func funca() {
	fmt.Println("aaaaaaaa")
}
func funcb() {
	panic("this is a panic test")
}
func funcc() {
	fmt.Println("ccccccc")
}
func main() {
	funca()
	funcb()
	funcc()
}
