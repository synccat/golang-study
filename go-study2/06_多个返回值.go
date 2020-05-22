package main

import "fmt"

func myFunc01() (int, int, int) {
	return 1, 2, 3
}
func myFunc02() (a int, b int, c int) {
	a, b, c = 111, 222, 333
	return

}
func myFunc03() (a, b, c int) {
	a, b, c = 1121, 2232, 3333
	return

}
func main() {
	a, b, c := myFunc03()
	fmt.Printf("a = %d,a = %d,a = %d\n", a, b, c)
}
