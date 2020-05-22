package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := make(map[string]interface{}, 4)
	m["company"] = "itcast"
	m["subjects"] = []string{"java", "c++", "python"}
	m["isok"] = true
	m["price"] = 888.888
	if buf, err := json.Marshal(m); err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Println("buf = ", string(buf))
	}
	if buf, err := json.MarshalIndent(m, "", "	"); err != nil {
		fmt.Println("err = ", err)

	} else {
		fmt.Println("buf = ", string(buf))

	}
}
