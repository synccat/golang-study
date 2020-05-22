package main

import (
	"fmt"
)

type Person struct {
	name string
	sex  byte
	age  int
}

func (Person) printInfo() {
	fmt.Println("name = ")
}
func (p *Person) addInfo(name string, sex byte, age int) {
	p.name = name
	p.sex = sex
	p.age = age
}
func main() {
	p := new(Person)
	p.addInfo("jerry", 'm', 21)
	p.printInfo()
}
