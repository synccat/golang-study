package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
	}
	defer resp.Body.Close()
	//循环读取网页数据，传出给调用者
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			// fmt.Println("读取网页完成")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		//累加每一次循环读取的数据，存入result
		result += string(buf[:n])
	}
	return
}

func working(start, end int) {
	fmt.Printf("正在爬取第%d页到%d页... \n", start, end)
	//循环爬取每一页的数据
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E5%89%91%E6%9D%A5&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		result, err := HttpGet(url)
		if err != nil {
			fmt.Println("HttpGet err:", err)
			continue
		}
		// fmt.Println("result=", result)
		//将读到的整网页数据保存成一个文件
		f, err1 := os.Create(strconv.Itoa(i) + ".html")
		if err1 != nil {
			fmt.Println("os.Create err", err1)
			continue
		}
		f.WriteString(result)
		f.Close()

	}
}
func main() {
	//指定爬虫起始、终止页
	var start, end int
	fmt.Println("请输入起始页码(>=1)：")
	fmt.Scan(&start)
	fmt.Println("请输入终止页面(>=start)：")
	fmt.Scan(&end)
	working(start, end)
}
