package main

import (
	"fmt"
	aaz "fmt"
)

func main() {
	var a = 1 << 10
	var b = 100 & 101
	var c float64 = float64(b) / 298
	fmt.Println("---->hello golang !!-->", a)
	fmt.Println("---->", b)
	fmt.Println("---->", c)
	aaz.Printf("---->%d\n",c)
}
