package main

import (
	"fmt"
)

type Person struct {
	name string
	sex  byte
	age  int
}

func (p *Person) printInfo(height int) {
	//	fmt.Printf("name = %s, sex = %c, age = %d\n", p.name, p.sex, p.age)
	fmt.Println("height = ", height)
}
func (p Person) printValue() {
	fmt.Println("2222222222")
}
func main() {
	var s1 Person = Person{"mike", 'm', 18}
	//方法表达式
	f1 := (*Person).printInfo
	f2 := Person.printValue
	f1(&s1, 171)
	f2(s1)
}
