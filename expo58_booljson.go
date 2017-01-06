package main

import (
	"log"
)

type A struct {
	B *bool `json:"b"`
}

func main() {
	var a A

	str := `{"b":false}`
	json.Unmarshal([]byte(str), &a)
	if a.B != nil {
		log.Println(*a.B)
	}
}
