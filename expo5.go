package main

import (
	"fmt"
	"math/rand"
	"time"
)

type TZ int

func main() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d / ", a)
	}
	fmt.Println()
	for i := 0; i < 5; i++ {
		r := rand.Intn(8)
		fmt.Printf("%d / ", r)
	}
	fmt.Println()
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	fmt.Println("---->", timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / ", 100*rand.Float32())
	}
	var a, b TZ = 3, 4
	c := a + b
	fmt.Printf("c has the value: %d", c) // 输出：c has the value: 7
	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point
	//判断是否为字母：unicode.IsLetter(ch)
	//判断是否为数字：unicode.IsDigit(ch)
	//判断是否为空白符号：unicode.IsSpace(ch)
	//var ah1, ah2, ah3 = "a", 1, ""
	//fmt.Println(unicode.IsLetter(ah1), unicode.IsDigit(ah2), unicode.IsSpace(ah3))
}
