package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company  string
	Subjects []string
	IsOk     bool
	Price    float64
}

func main() {
	jsonBuf := `{"company":"itcast","subjects":["java","c++","python"],"isok":true,"price":8888.88}`
	var temp IT
	if err := json.Unmarshal([]byte(jsonBuf), &temp); err == nil {
		fmt.Println("temp = ", temp)
	}

}
