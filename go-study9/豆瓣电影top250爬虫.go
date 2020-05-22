package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

var page chan int = make(chan int)

func saveFile(i int, titleArray, scoreArray, numArray [][]string) {
	path := "C:/GOPATH/src/go-study9/douban/" + strconv.Itoa(i) + ".txt"
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("os.Create err = ", err)
		return
	}
	file.WriteString("电影名称\t\t\t评分\t\t\t评分人数\r\n")
	for idx, data := range titleArray {
		file.WriteString(data[1] + "\t\t\t" + scoreArray[idx][1] + "\t\t\t" + numArray[idx][1] + "\r\n")
	}
	file.Close()
}

func spiderData(i int) {
	srcUrl := "https://movie.douban.com/top250?start=" + string(strconv.Itoa((i-1)*25)) + "&filter="
	result, err := httpGet(srcUrl)
	if err != nil {
		fmt.Println("httpGet err", err)
	}
	//正则表达式
	titleRe := regexp.MustCompile(`<img width="100" alt="(?s:(.*?))"`)
	scoreRe := regexp.MustCompile(`<span class="rating_num" property="v:average">(?s:(.*?))</span>`)
	numRe := regexp.MustCompile(`<span>(.*?)人评价</span>`)
	titleArray := titleRe.FindAllStringSubmatch(result, -1)
	scoreArray := scoreRe.FindAllStringSubmatch(result, -1)
	numArray := numRe.FindAllStringSubmatch(result, -1)
	saveFile(i, titleArray, scoreArray, numArray)
	page <- i

}
func httpGet(srcUrl string) (result string, err error) {
	resp, err1 := http.Get(srcUrl)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)

	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 && err2 == io.EOF {
			break
		}
		if err2 != nil && err2 != io.EOF {
			fmt.Println("resp.Body.Read err=", err2)
			return
		}
		result += string(buf[:n])
	}
	return
}
func startSpider(start, end int) {
	for i := start; i <= end; i++ {
		go spiderData(i)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d页已经抓取完成\n", <-page)
	}
}
func main() {
	startSpider(1, 10)
}
