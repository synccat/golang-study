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
	*Person     //只有类型，没有名字，匿名字段，继承了Person里面的成员
	id      int //基础类型的匿名字段
	addr    string
}

func main() {
	s2 := new(Student)
	s2.Person = new(Person)
	s2.name = "ha"
	s2.sex = 'm'
	s2.age = 18
	fmt.Println(s2.name, s2.age, s2.sex)
}
