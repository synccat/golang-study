package main

import (
	"fmt"
)

type Student struct {
	name string
	age  int
}

func main() {
	s := make([]interface{}, 3)
	s[0] = 1233
	s[1] = "tomandjerry"
	s[2] = Student{"tom", 22}
	for i, data := range s {
		if value, ok := data.(int); ok == true {
			fmt.Printf("s[%d]的值是：%d,类型是:int\n", i, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("s[%d]的值是：%s,类型是:string\n", i, value)

		} else if value, ok := data.(Student); ok == true {
			fmt.Printf("s[%d]的值是：name = %s,age = %d,类型是:Student\n", i, value.name, value.age)

		}
	}
}
