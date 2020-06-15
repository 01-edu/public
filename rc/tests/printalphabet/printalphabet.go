package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func fillArray(a []rune) {
	for i := 'a'; i <= 'z'; i++ {
		a = append(a, i)
	}
}

func main() {
	a := []rune{'a', 'b', 'c', 'd', 'e', 'f'}
	b := int('a')
	for _, v := range a {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func testingScope() {
	defFun := func(s string) {
		fmt.Println(s)
	}
	defFun("Hello")
}
