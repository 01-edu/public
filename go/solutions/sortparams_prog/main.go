package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	args := os.Args[1:]
	sort.Strings(args)
	for _, v := range args {
		fmt.Println(v)
	}
}
