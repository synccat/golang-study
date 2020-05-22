package main

import "fmt"

func main() {
	str := "abc"
	//通过for循环打印每个字符
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}

	//循环打印妹子元素，默认返回2个值,一个是元素位置，一个是元素本身
	for i, data := range str {
		fmt.Printf("str[%d]=%c\n", i, data)
	}

	for i := range str { //第二个元素默认丢弃
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}
}
