package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	//服务监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	defer listener.Close()
	if err != nil {
		fmt.Println("net.Listen err：", err)
		return
	}
	// 建立连接
	conn, err1 := listener.Accept()
	defer conn.Close()
	if err1 != nil {
		fmt.Println("listener.Accept err：", err1)
		return
	}
	buf := make([]byte, 1024<<3)
	// 接收文件名
	var fileName string
	n, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("conn.Read err：", err2)
		conn.Write([]byte("error"))
		return
	} else {
		conn.Write([]byte("ok"))
		fileName = string(buf[:n])
	}
	// 创建文件
	file, err3 := os.Create(fileName)
	if err3 != nil {
		fmt.Println("os.Create err：", err3)
	}
	// 写入文件
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件传输完成")
				return
			}
		} else {
			file.Write(buf[:n])
		}
		if n == 0 {
			fmt.Println("文件传输完成")
			return
		}
	}
}
