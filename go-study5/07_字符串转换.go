package main

import (
	"fmt"
	"strconv"
)

func main() {
	//转换字符串后追加到字节数组
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	slice = strconv.AppendInt(slice, 9879, 10)
	slice = strconv.AppendQuote(slice, "dsfwefawfawe")
	fmt.Println("slice = ", string(slice))
	//其他类型转换为字符串
	var str string
	str = strconv.FormatBool(false)
	fmt.Println("str = ", str)
	str = strconv.FormatFloat(3.1415926, 'f', -1, 64)
	fmt.Println("str = ", str)
	str = strconv.FormatInt(3223, 10)
	fmt.Println("str = ", str)

	//整形转字符串，常用
	str = strconv.Itoa(231131)
	fmt.Println("str = ", str)

	//字符串转其他类型
	// var flag bool
	// var err error

	if flag, err := strconv.ParseBool("false"); err == nil {
		fmt.Println("flag = ", flag)
	} else {
		fmt.Println("err = ", err)
	}
	if num, err := strconv.Atoi("32142d41"); err == nil {
		fmt.Println("num = ", num)
	} else {
		fmt.Println("err = ", err)
	}

}
