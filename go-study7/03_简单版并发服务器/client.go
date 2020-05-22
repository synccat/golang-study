package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	//连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("1", err)
		return
	}

	//监听键盘输入，并传送到服务器
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil {
				fmt.Println("3", err)
				continue
			}
			if _, err := conn.Write(buf[:n-2]); err != nil {
				break
			}

		}
	}()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				conn.Close()
				break
			}
			fmt.Println("2", err)
		}
		fmt.Println(string(buf[:n]))
	}

	defer conn.Close()

}
