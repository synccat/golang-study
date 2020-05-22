package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http.Get err: ", err)
	}
	fmt.Println("Status=", resp.Status)
	fmt.Println("StatusCode=", resp.StatusCode)
	fmt.Println("Header=", resp.Header)
	fmt.Println("Body=", resp.Body)
	buf := make([]byte, 1024<<2)
	var temp string
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		temp += string(buf[:n])
	}
	fmt.Println("body = ", temp)
	defer resp.Body.Close()
}
