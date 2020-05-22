package main

import "fmt"

func main() {
	var ch byte
	ch = 97
	fmt.Println("ch = ", ch)
	//格式化输出，%c以字符方式打印，%d以整形方式打印
	fmt.Printf("%c, %d\n", ch, ch)
	ch = 'a'
	fmt.Printf("%c, %d", ch, ch)

	//大写转小写，小写转大写，大小写相差32，小写大
	fmt.Printf("大写：%d, 小写：%d\n", 'A', 'a')
	fmt.Printf("大写转小写：%c\n", 'A'+32)
	fmt.Printf("小写转大写：%c\n", 'a'-32)
}
