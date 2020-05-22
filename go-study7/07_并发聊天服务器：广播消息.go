package main

import (
	"fmt"
	"net"
	"strings"
	"time"
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

	cli := client{make(chan string), addr, addr}
	onlineMap[addr] = cli
	msg := "[" + addr + "]" + addr + "login"

	go sendMsg(msg)
	go castMsg(conn, cli)
	//从客户端接收消息并广播
	buf := make([]byte, 1024)
	go func() {
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				fmt.Println(" conn.Read err: ", err)
				msg := "---[" + cli.Addr + "]" + cli.Name + "logout"
				go sendMsg(msg)
				delete(onlineMap, addr)
				return
			} else {
				//查询当前用户
				if string(buf[:n-1]) == "who" {
					conn.Write([]byte("online list:\n"))
					for _, cli := range onlineMap {
						msg := "---[" + cli.Addr + "]" + cli.Name + "\n"
						conn.Write([]byte(msg))
					}
				} else if len(string(buf[:n-1])) > 8 && string(buf[:6]) == "rename" {
					name := strings.Split(string(buf[:n-1]), " ")[1]
					cli.Name = name
					onlineMap[addr] = cli
				} else if string(buf[:n-1]) == "exit" {
					msg := "---[" + cli.Addr + "]" + cli.Name + "logout"
					go sendMsg(msg)
					delete(onlineMap, addr)
					return

				} else {
					msg := "[" + addr + "]" + cli.Name + ": " + string(buf[:n-1])
					go sendMsg(msg)

				}
			}
		}
	}()

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
