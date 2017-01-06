package main

import (
	"io"
	"log"
)

func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err) //在return之后执行，所以获取到所有值
	}()
	return 7, io.EOF
}

func main() {
	func1("Go")
}
