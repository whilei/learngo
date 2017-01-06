package main

import (
	"bytes"
	"fmt"
)

func main() {
	var slice1 []int = make([]int, 10, 20)
	// load the array/slice:
	for i := 0; i < len(slice1); i++ {
		slice1[i] = slice1[i] + 5*i
	}

	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("\nThe length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
	slice1 = slice1[1:8]
	fmt.Printf("\n--->%d\n", slice1[0])

	var slicestring1 []byte = []byte{'g', 'o', 'o', 'g', 'l', 'e'}
	fmt.Printf("---->%v", string(Append(slicestring1[0:1], slicestring1[1:])))

}

func Append(slice, data []byte) []byte {
	var buffer bytes.Buffer
	var p, i int
	for i = 0; i < len(slice); i++ {
		buffer.WriteByte(slice[i])
		//buffer.WriteString(string(slice[i])) //注意类型
	}
	for p = 0; p < len(data); p++ {
		buffer.WriteByte(data[p])
		//buffer.WriteString(string(data[p]))
	}
	// return buffer.String() //返回string ，不是 []string
	return buffer.Bytes()
}
