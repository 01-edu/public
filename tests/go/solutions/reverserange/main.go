package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 3 {
		secondArg := 0
		firstArg, err := strconv.Atoi(os.Args[1])
		if err == nil {
			secondArg, err = strconv.Atoi(os.Args[2])
		}
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(reverseRange(firstArg, secondArg))

	} else {
		fmt.Println()
	}
}

func reverseRange(start, end int) []int {
	var rran []int
	if end >= start {
		for i := end; i >= start; i-- {
			rran = append(rran, i)
		}
		return rran
	}
	for i := end; i <= start; i++ {
		rran = append(rran, i)
	}
	return rran
}
