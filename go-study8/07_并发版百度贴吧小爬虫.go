package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var ch chan int = make(chan int)

func GetHtml(i int) {
	//连接url
	url := "https://tieba.baidu.com/f?kw=%E5%89%91%E6%9D%A5&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	fmt.Printf("开始爬取第%d页内容\n", i)
	resp, err1 := http.Get(url)
	if err1 != nil {
		fmt.Println("http.Get err=", err1)
		return
	}
	defer resp.Body.Close()
	var str string
	for {
		//读取内容
		buf := make([]byte, 1024<<2)
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		str += string(buf[:n])
	}
	//创建文件
	filename := strconv.Itoa(i) + ".html"
	file, _ := os.Create(filename)
	file.WriteString(str)
	file.Close()
	ch <- i
}

func DoWork(start int, end int) {
	for i := start; i <= end; i++ {
		go GetHtml(i)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("完成爬取第%d页内容\n", <-ch)
	}

}
func main() {
	var start, end int
	fmt.Println("请输入要爬取的起始页（>=1）：")
	fmt.Scan(&start)
	fmt.Println("请输入要爬取的终止页（>=起始页）：")
	fmt.Scan(&end)
	DoWork(start, end)
}
