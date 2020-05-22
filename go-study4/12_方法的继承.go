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

func (p *Person) printInfo() {
	fmt.Printf("name = %s, sex = %c, age = %d", p.name, p.sex, p.age)
}

func main() {
	var s1 Student = Student{Person{"mike", 'm', 18}, 1, "beijing"}
	s1.printInfo()
}
