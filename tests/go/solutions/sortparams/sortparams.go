package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	args := os.Args[1:]
	sort.Sort(sort.StringSlice(args))
	for _, v := range args {
		fmt.Println(v)
	}
}
