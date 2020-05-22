package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

func spiderOnePage(url string)(result string,err error){
	resp,err1 :=http.Get(url)
	if err1!=nil{
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte,4096)
	for{
		n,err2 :=resp.Body.Read(buf)
		if n==0{
			return
		}
		if err2==nil&&n!=0{
			err =err2
			return
		}
		result +=string(buf[:n])
	}
	fmt.Println(result)
	return
}
func doSpider(i int){
	url := "https://www.pengfu.com/index_"+strconv.Itoa(i)+".html"
	result,err :=spiderOnePage(url)
	if err!=nil{
		fmt.Println("spiderPage err",err)
	}
	//通过正则表达式，取得每一个笑话的链接
	re := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	urls :=re.FindAllStringSubmatch(result,-1)
	fmt.Println(urls)
}
func startSpider(start,end int){
	fmt.Println("开始爬取数据")
	//循环爬取每一页的数据
	for i:=start;i<=end;i++{
		doSpider(i)
	}

}
func main(){
	var start, end int
	fmt.Println("请输入要爬取的起始页(>=1):")
	fmt.Scan(&start)
	fmt.Println("请输入要爬取的终止页(>=start)：")
	fmt.Scan(&end)
	startSpider(start,end)
}