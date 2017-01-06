package main

import (
	"fmt"
	"time"
)

const LIM = 1000

var fibs [LIM]uint64

func main() {
	var result, s1, s2 uint64 = 0, 0, 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = s1 + s2
		fibs[i] = result
		switch result {
		case 0:
			s2 = 1
			result = 1
		case 1:
			s1 = 1
		default:
			s1 = s2
			s2 = result
		}

		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
