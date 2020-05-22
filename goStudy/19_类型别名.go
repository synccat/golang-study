package main

import "fmt"

func main() {
	//给int64起一个别名叫bigint
	type bigint int64
	var a bigint
	fmt.Printf("a type is %T", a)

}
