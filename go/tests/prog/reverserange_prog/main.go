package main

import (
	"fmt"
	"os"
	"sort"
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
	suite := lib.IntRange(a, b)
	sort.Sort(sort.Reverse(sort.IntSlice(suite)))
	for _, i := range suite {
		fmt.Print(i)
	}
	fmt.Println()
}
