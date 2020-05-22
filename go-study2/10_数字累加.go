package main

import "fmt"

func test01() (sum int) {
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return
}
func test03(i int) (sum int) {
	if i == 1 {
		sum = i
		return
	}
	sum = i + test03(i-1)
	return
}
func main() {
	fmt.Println("for循环计算出来的结果是：", test01())
	//	fmt.Println("递归循环，从1到100的结果是：", test02(1))
	fmt.Println("递归循环，从100到1的结果是：", test03(100))
}
