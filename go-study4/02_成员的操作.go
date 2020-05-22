package main

import (
	"fmt"
)

type Person struct {
	name string
	sex  byte
	age  int
}
type Student struct {
	Person //只有类型，没有名字，匿名字段，继承了Person里面的成员
	id     int
	addr   string
}

func main() {
	s1 := Student{Person{"mike", 'm', 18}, 1, "beijing"}
	fmt.Println(s1.name, s1.sex, s1.age, s1.id, s1.addr)
}
