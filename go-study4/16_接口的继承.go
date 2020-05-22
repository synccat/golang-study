package main

import (
	"fmt"
)

type Person interface {
	sayhi()
}
type People interface {
	Person
	sing(lrc string)
}
type Teacher struct {
	name string
	age  int
}
type Student struct {
	name string
	age  int
}

func (t *Teacher) sayhi() {
	fmt.Println("hello Teacher!")
}
func (s *Student) sayhi() {
	fmt.Println("hello Student!")
}
func (s *Student) sing(lrc string) {
	fmt.Println("Student is singing ", lrc)
}
func main() {
	s := &Student{"jerry", 22}
	var p People
	p = s
	p.sayhi()
	p.sing("窗外的麻雀......")

}
