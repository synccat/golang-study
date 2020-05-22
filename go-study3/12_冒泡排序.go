package main

import (
	"fmt"
)

func mpSort(sortIn *[]int) {
	sort := *sortIn
	for i := 0; i < len(*sortIn); i++ {
		for j := 0; j < len(sort)-i-1; j++ {
			if sort[j] > sort[j+1] {
				sort[j], sort[j+1] = sort[j+1], sort[j]
			}
		}
	}
	return
}

func main() {
	arrayA := []int{3, 213, 4, 2, 13, 4, 99, 5, 6}
	fmt.Println(arrayA)
	mpSort(&arrayA)
	fmt.Println(arrayA)
}
