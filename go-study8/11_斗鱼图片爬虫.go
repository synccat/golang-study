package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func HttpGets(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			return
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return
}
func saveImg(idx int, url string, page chan int) {
	fmt.Printf("开始下载第%d张图片\n", idx)
	f, err := os.Create("C:/GOPATH/src/go-study8/img/" + strconv.Itoa(idx) + ".jpg")
	if err != nil {
		fmt.Println("os.Create err=", err)
		return
	}
	defer f.Close()
	resp, err1 := http.Get(url)
	if err1 != nil {
		fmt.Println("http.Get err=", err)
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		f.Write(buf[:n])
	}
	page <- idx
}
func main() {
	url := "https://www.douyu.com/g_yz"
	result, err := HttpGets(url)
	if err != nil {
		fmt.Println("HttpGet err=", err)
	}
	re := regexp.MustCompile(`"rs1":"(?s:(.*?))","`)
	allLinks := re.FindAllStringSubmatch(result, -1)
	page := make(chan int)
	n := len(allLinks)
	for idx, imgUrl := range allLinks {
		go saveImg(idx, imgUrl[1], page)
	}
	for i := 0; i < n; i++ {
		fmt.Printf("下载第 %d 张图片完成\n", <-page)
	}
}
