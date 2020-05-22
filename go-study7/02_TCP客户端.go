package main

import (
	"fmt"
	"net"
)

func main() {
	//主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	conn.Write([]byte("are you ok?"))
}
