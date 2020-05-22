package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := fmt.Errorf("%s", "this is a normal error!")

	fmt.Println("err1 = ", err1)
	err2 := errors.New("this is another normal error!")
	fmt.Println("err2 = ", err2)

}
