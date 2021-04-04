package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(a)
	for b != a {
		if a < b {
			a++
		} else {
			a--
		}
		fmt.Print(" ", a)
	}
	fmt.Println()
}
