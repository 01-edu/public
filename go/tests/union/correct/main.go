package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 3 {
		var res string
		s1 := os.Args[1]
		s2 := os.Args[2]

		for _, v := range s1 {
			if !strings.ContainsRune(res, v) {
				res += string(v)
			}
		}
		for _, v := range s2 {
			if !strings.ContainsRune(res, v) {
				res += string(v)
			}
		}
		fmt.Print(res)
	}
	fmt.Println()
}
