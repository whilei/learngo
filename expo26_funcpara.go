package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	callback(1, Add)
	println(strings.IndexFunc("abc xyz", unicode.IsSpace))
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2) ，个人观点尽量避免使用这种混淆用法
}
