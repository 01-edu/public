package main

import (
	"os"
	"fmt"
	"strings"
)

func contains(big, small string) bool {
	for i := 0; i <= len(big)-len(small); i++ {
		if big[i] == small[0] {
			for j := 0; j < len(small); j++ {
				if big[i+j] != small[j] {
					break
				}
				if j == len(small)-1 {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	var big string
	args := os.Args[1:]

	for _, arg := range args {
		split := strings.Split(arg, " ")
		ch := split[0]
		small := split[1]

		if ch == "x" {
			break
		}
		if ch == "A" {
			big = big + small
		}
		if ch == "?" {
			if contains(big, small) == true {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
	}
}
