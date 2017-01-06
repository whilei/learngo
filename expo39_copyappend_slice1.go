package main

import "fmt"

func main() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)

	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	/*
		append 会分配新的切片来保证已有切片元素和新增元素的存储。
		因此，返回的切片可能已经指向一个不同的相关数组了。append 方法总是返回成功，除非系统内存耗尽了。
	*/
	fmt.Println(sl3)
	sL4 := make([]int, 10)
	sL4 = InsertStringSlice(sl3, sl3, 23)
	fmt.Println(sL4)
}

func InsertStringSlice(slice_source []int, slice_insert []int, insertIndex int) []int {
	var m = len(slice_source)
	var n = len(slice_insert)
	var newSlice []int
	fmt.Println("--->%d,%d,%v,%v", m, n, slice_source, slice_insert)

	if m < insertIndex {
		fmt.Println("Insert index point is larger than length of slice_source , ERROR")
		//return null
		insertIndex = m
	}

	// newSlice = make([]int, (m+n)*2)

	fmt.Println("0--->%v", newSlice)
	for v := 0; v < insertIndex; v++ {
		newSlice = append(newSlice, slice_insert[v])
		fmt.Println("AT o--->%v", newSlice)
	}
	fmt.Println("1--->%v", newSlice)
	for i := 0; i < n; i++ {
		newSlice = append(newSlice, slice_insert[i])
		fmt.Println("AT i--->%v", newSlice)
	}
	fmt.Println("2--->%v", newSlice)
	for p := insertIndex; p < m; p++ {
		newSlice = append(newSlice, slice_source[p])
		fmt.Println("AT p--->%v", newSlice)
	}
	fmt.Println("3--->%v", newSlice)
	return newSlice[0 : m+n]
}
