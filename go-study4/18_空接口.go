package main

import (
	"fmt"
)

func main() {
	var i interface{} = 1
	fmt.Println("i = ", i)
	i = "asdfasdf"
	fmt.Println("i = ", i)
}
