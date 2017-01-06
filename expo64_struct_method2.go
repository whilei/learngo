package main

import "fmt"

type IntVector []int

type Employee struct {
	name   string
	salary float32
}

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

func (p Employee) Upsalary() (w float32) {
	if p.salary > 0 {
		return p.salary * 1.18
	}
	return 0
}

func main() {
	fmt.Println(IntVector{1, 2, 3}.Sum()) // 输出是6
	emp1 := Employee{"Joe", 28292.22}
	fmt.Println("--->", emp1.name, emp1.Upsalary())
}
