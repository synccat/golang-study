package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func DoWork(start int, end int) (err error) {
	for i := start; i <= end; i++ {
		//创建地址
		page := (i - 1) * 50
		url := "https://tieba.baidu.com/f?kw=%E5%89%91%E6%9D%A5&ie=utf-8&pn=" + strconv.Itoa(page)
		//连接url
		resp, err1 := http.Get(url)
		if err1 != nil {
			continue
		}
		defer resp.Body.Close()
		var str string
		for {
			//读取内容
			buf := make([]byte, 1024<<2)
			n, err1 := resp.Body.Read(buf)
			if n == 0 {
				err = err1
				break
			}
			str += string(buf[:n])
		}
		//创建文件
		filename := strconv.Itoa(i) + ".html"
		file, _ := os.Create(filename)
		defer file.Close()
		file.WriteString(str)
	}
	return
}
func main() {
	var start, end int
	fmt.Println("请输入要爬取的起始页（>=1）：")
	fmt.Scan(&start)
	fmt.Println("请输入要爬取的终止页（>=起始页）：")
	fmt.Scan(&end)
	DoWork(start, end)
}
