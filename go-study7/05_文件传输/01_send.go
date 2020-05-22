package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func sendFile(path string, conn net.Conn) {
	//以只读方式打开文件
	file, err := os.Open(path)
	fmt.Println(path)
	if err != nil {
		fmt.Println("os.Open err: ", err)
		return
	}
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件传输完成")
				return
			}
			fmt.Println("file.Read err: ", err)
			return
		}
		//向服务器发送数据
		n, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("conn.Write err: ", err)
			return
		}

	}
}
func main() {
	//提示输入文件
	fmt.Println("请输入文件名称：")
	var path string
	fmt.Scan(&path)
	//获取文件名
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat error: ", err)
		return
	}
	//连接服务器
	conn, err1 := net.Dial("tcp", "127.0.0.1:8000")
	if err1 != nil {
		fmt.Println("net.Dial err: ", err1)
		return
	}
	defer conn.Close()
	//发送文件名到服务器
	_, err2 := conn.Write([]byte(info.Name()))
	if err2 != nil {
		fmt.Println("conn.Write err: ", err2)
	}
	//接收服务器回复，回复ok则执行上传
	buf := make([]byte, 1024)
	n3, err3 := conn.Read(buf)
	if err3 != nil {
		fmt.Println("conn.Read err: ", err3)
		return
	}
	if string(buf[:n3]) == "ok" {
		//发送文件内容
		sendFile(path, conn)
	}
}
