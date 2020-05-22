package main

import (
	"fmt"
	"net"
	"strings"
)

func handleRequest(conn net.Conn) {
	defer conn.Close()
	addr := conn.RemoteAddr().String()
	fmt.Printf("%s is connected", addr)
	buf := make([]byte, 1024)
	var i int = 0
	for {
		fmt.Printf("【%s】第%d次发送请求数据！\n", addr, i)
		i++
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		if string(buf[:n]) == "exit" {
			fmt.Printf("【%s】请求退出连接\n", addr)
			fmt.Printf("【%s】已断开连接\n", addr)
			return
		}

		//转换为大写
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))

	}
}
func main() {
	//建立监听服务
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
	}
	defer listener.Close()
	//建立连接，多用户
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleRequest(conn)
	}

}
