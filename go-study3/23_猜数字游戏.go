package main

import (
	"fmt"
	"math/rand"
	"time"
)

//进行游戏
func doGame(randSlice []int) {
	var inNum int
	keySlice := make([]int, 4)
	fmt.Println("请输入一个四位数：")
	for {
		fmt.Scan(&inNum)
		if inNum < 1000 || inNum > 9999 {
			fmt.Println("请输入正确的4位数！")
		} else {
			n := 0
			//随机数与键盘输入的比较
			getSlice(keySlice, inNum)
			for i := 0; i < 4; i++ {
				if keySlice[i] < randSlice[i] {
					fmt.Printf("第%d位小了\n", i+1)
				} else if keySlice[i] > randSlice[i] {
					fmt.Printf("第%d位大了\n", i+1)
				} else {
					n++
				}
			}
			if n == 4 {
				fmt.Println("恭喜你，答对了！")
				break
			} else {
				fmt.Println("请重新输入：")
			}
		}
	}
}

//数字转切片
func getSlice(srcSlice []int, num int) {
	srcSlice[0] = num / 1000
	srcSlice[1] = num % 1000 / 100
	srcSlice[2] = num % 100 / 10
	srcSlice[3] = num % 10
}

//生成随机数
func getRandNum() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(9999-999) + 999
}

func main() {
	randNum := getRandNum()
	randSlice := make([]int, 4)
	getSlice(randSlice, randNum)
	fmt.Printf("%v", randSlice)
	doGame(randSlice)
}
