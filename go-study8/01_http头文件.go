package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("listener.Accept err = ", err1)
		return
	}
	buf := make([]byte, 1024<<2)
	n, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("conn.Read err = ", err2)
		return
	}
	fmt.Println(string(buf[:n]))
}
