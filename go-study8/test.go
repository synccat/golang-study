package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 50
	var s string = "sdas"
	s = s + strconv.Itoa(i)
	fmt.Println(s)
}
