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
	fmt.Print(b)
	for a != b {
		if b < a {
			b++
		} else {
			b--
		}
		fmt.Print(" ", b)
	}
	fmt.Println()
}
