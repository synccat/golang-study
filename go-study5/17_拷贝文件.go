package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readFileLine(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
		}
		fmt.Printf("buf = [%s]\n", string(buf))
	}

}
func main() {
	attr := os.Args
	if len(attr) != 3 {
		fmt.Println("请输入正确的命令：xxx.exe sourceDir destDir")
		return
	}
	sourcePath := attr[1]
	destPath := attr[2]
	//打开源文件
	sF, err := os.Open(sourcePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sF.Close()
	//创建新文件
	dF, err := os.Create(destPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dF.Close()
	//设置合适大小的缓冲，循环写入文件
	buf := make([]byte, 1024*1024)
	var count int = 0
	for {
		count++
		fmt.Println(count)
		n, err := sF.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("拷贝完成")
				break
			}
		} else {
			_, err1 := dF.Write(buf[:n])
			if err1 != nil {
				fmt.Println(err1)
			}
		}
	}

}
