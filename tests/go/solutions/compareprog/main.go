package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 3 {
		fmt.Println(compare(os.Args[1], os.Args[2]))
	} else {
		fmt.Println()
	}
}

func compare(s string, toCompare string) int {
	result := strings.Compare(s, toCompare)
	return result
}
