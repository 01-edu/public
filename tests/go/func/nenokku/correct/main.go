package main

import (
	"fmt"
	"os"
	"strings"
)

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
			big += small
		}
		if ch == "?" {
			if strings.Contains(big, small) {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
	}
}
