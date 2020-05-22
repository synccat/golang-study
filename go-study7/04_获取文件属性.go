package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args
	if len(list) != 2 {
		fmt.Println("usage: xxxx.xx")
		return
	}
	fileName := list[1]
	info, err := os.Stat(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(info.Name())
	fmt.Println(info.Size())
}
