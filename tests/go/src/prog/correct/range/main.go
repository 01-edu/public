package main

import (
	"fmt"
	"os"
	"strconv"
)

func rangeOf(start, end int) []int {
	var ran []int
	if start >= end {
		for i := start; i >= end; i-- {
			ran = append(ran, i)
		}
		return ran
	}
	for i := start; i <= end; i++ {
		ran = append(ran, i)
	}
	return ran
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	secondArg := 0
	firstArg, err := strconv.Atoi(os.Args[1])
	if err == nil {
		secondArg, err = strconv.Atoi(os.Args[2])
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rangeOf(firstArg, secondArg))
}
