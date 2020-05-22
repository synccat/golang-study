package main

import (
	"errors"
	"fmt"
)

func devNum(a int, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("除数不能为0")
	} else {
		result = a / b
	}
	return
}
func main() {
	result, err := devNum(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
