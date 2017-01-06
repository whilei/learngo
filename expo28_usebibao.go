package main

import (
	"fmt"
	"strings"
)

func main() {
	// make an Add2 function, give it a name p2, and call it:
	p2 := Add2()
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")

	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
	// make a special Adder function, a gets value 3:
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))
	fmt.Printf("--->addBmp is :%v\n", addBmp("file"))   // returns: file.bmp
	fmt.Printf("--->addJpeg is :%v\n", addJpeg("file")) // returns: file.jpeg
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func MakeAddSuffix(suffix string) func(string) string { //工厂函数，也可以叫构造函数，构造其他函数的母体函数
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

//感觉这个例子就是func嵌套，为何不把功能写到一个func中，而是嵌套在多个func中，可能有些场景会比较适合，但不一定所有都适合
