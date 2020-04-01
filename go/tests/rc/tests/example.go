package solutions

import (
	"fmt"
)

func thisIsAFunc() {
	fmt.Println("This is a function")
}

var ThisToo = func(s string) {
	fmt.Printf("ThisToo prints %s\n", s)
}

func aux(s string) int {
	return 1
}

func youCanAlso(f func(string), s string) {
	aux := func(s string) int {
		return len(s)
	}
	aux(s)
	f(s)
}
