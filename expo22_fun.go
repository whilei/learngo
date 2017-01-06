package main

import "fmt"

var num int = 10
var numx2, numx3 int

func main() {
	numx2, numx3 = getX2AndX3(num)
	PrintValues()
	numx2, numx3 = getX2AndX3_2(num)
	PrintValues()
	var i1 int
	var f1 float32
	i1, _, f1 = ThreeValues()
	fmt.Printf("The int: %d, the float: %f \n", i1, f1)
}

func PrintValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}

func getX2AndX3(input int) (int, int) { //无命名的返回值 缺省
	return 2 * input, 3 * input
}

func getX2AndX3_2(input int) (x2 int, x3 int) { //命名返回值
	x2 = 2 * input
	x3 = 3 * input
	// return x2, x3
	return
}

func ThreeValues() (int, int, float32) {
	return 5, 6, 7.5
}
