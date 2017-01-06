package main

import (
	"fmt"
)

func main() {
	println("In main before calling greeting")
	greeting()
	println("In main after calling greeting")
	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", MultiPly3Nums(2, 5, 6))
	// var i1 int = MultiPly3Nums(2, 5, 6)
	// fmt.Printf("MultiPly 2 * 5 * 6 = %d\n", i1)
}

func greeting() {
	println("In greeting: Hi!!!!!")
}

func MultiPly3Nums(a, b, c int) int {
	// var product int = a * b * c
	// return product
	return a * b * c
}
