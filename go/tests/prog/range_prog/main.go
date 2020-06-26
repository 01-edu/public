package main

import (
	"fmt"
	"os"
	"strconv"

	"lib"
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
	for _, i := range lib.IntRange(a, b) {
		fmt.Print(i)
	}
	fmt.Println()
}
