package main

import "fmt"

func main() {
	function1()
	f()
}

func function1() {
	defer function2() //无论放在任何位置，只会在最后return的时候执行
	fmt.Printf("In function1 at the top\n")
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("function2: Deferred until the end of the calling function!")
}

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("\n%d", i)
	}
}
