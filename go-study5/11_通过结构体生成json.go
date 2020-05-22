package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company  string   `json:"company"`
	Subjects []string `json:"subjects"`
	IsOk     bool     `json:"isok"`
	Price    float64  `json:"price"`
}

func main() {
	s := IT{"itcast", []string{"java", "c++", "python"}, true, 8888.88}
	if buf, err := json.Marshal(s); err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Println("buf = ", string(buf))
	}
	if buf, err := json.MarshalIndent(s, "", "	"); err != nil {
		fmt.Println("err = ", err)

	} else {
		fmt.Println("buf = ", string(buf))

	}
}
