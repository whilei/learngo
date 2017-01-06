package main

import "fmt"

func main() {
	// Version A:
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
		fmt.Println("items[%v] --->", items[i])
	}
	fmt.Printf("Version A: Value of items: %v\n", items)

	// Version B: NOT GOOD!
	items2 := make([]map[int]int, 5)
	items3 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1) // item is only a copy of the slice element.
		item[1] = 2                 // This 'item' will be lost on the next iteration.丢失
		fmt.Printf("Version A: Value of item: %v\n", item[1])
	}
	fmt.Printf("Version B: Value of items: %v\n", items2)
	fmt.Printf("Version C: Value of items: %v\n", items3)
}
