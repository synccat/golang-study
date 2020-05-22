package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := make(map[string]interface{}, 4)
	jsonBuf := `{"company":"itcast","subjects":["java","c++","python"],"isok":true,"price":8888.88}`
	fmt.Printf("%T\n", jsonBuf)
	if err := json.Unmarshal([]byte(jsonBuf), &m); err == nil {
		fmt.Println(m)
	}
}
