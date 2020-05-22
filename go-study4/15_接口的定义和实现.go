package main

import (
	"fmt"
)

type Person interface {
	sayhi()
}
type Teacher struct {
	name string
	age  int
}
type Student struct {
	name string
	age  int
}

func WhoSayHi(p Person) {
	p.sayhi()
}
func (t *Teacher) sayhi() {
	fmt.Println("hello Teacher!")
}
func (s *Student) sayhi() {
	fmt.Println("hello Student!")
}
func main() {
	s := &Student{"tom", 22}
	t := &Teacher{"jerry", 3}
	p := make([]Person, 2)
	p[0] = s
	p[1] = t
	for _, pp := range p {
		WhoSayHi(pp)
	}
}
