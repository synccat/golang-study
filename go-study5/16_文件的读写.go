package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func writeFile(path string) {
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	var buf string
	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("i = %d\n", i)
		if _, err := file.Write([]byte(buf)); err != nil {
			fmt.Println(err)
		}
	}

}
func readFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	buf := make([]byte, 1024*2) //2k大小
	if n, err := f.Read(buf); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(buf[:n]))
	}

}
func readFileLine(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
		}
		fmt.Printf("buf = [%s]\n", string(buf))
	}

}
func main() {
	path := "./demo.txt"
	// writeFile(path)
	// readFile(path)
	readFileLine(path)
}
