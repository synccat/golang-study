package main

import (
	"fmt"
	"net"
)

//用户结构体
type client struct {
	C    chan string
	Name string
	Addr string
}

//用一个map缓存在线用户
var onlineMap map[string]client

func sendMsg(msg string) {
	for _, cli := range onlineMap {
		cli.C <- msg
	}
}

func castMsg(conn net.Conn, cli client) {
	defer conn.Close()
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}

}

func handleConn(conn net.Conn) {
	//获取用户信息
	addr := conn.RemoteAddr().String()
	// buf := make([]byte, 1024)
	// _, err := conn.Read(buf)
	// if err != nil {
	// 	return
	// }
	cli := client{make(chan string), addr, addr}
	onlineMap[addr] = cli
	msg := "[" + addr + "]" + addr + "login"
	go sendMsg(msg)
	go castMsg(conn, cli)

}

func main() {
	onlineMap = make(map[string]client, 10)
	//监听服务
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Listen err: ", err)
	}
	defer listener.Close()
	//新开协程，转发消息
	// go manager()

	//循环接收用户请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err: ")
			continue
		}
		//处理用户请求
		go handleConn(conn)
	}
}
