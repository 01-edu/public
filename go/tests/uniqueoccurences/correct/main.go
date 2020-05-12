package main

import (
	"fmt"
	"os"
)

func solve(s string) bool {
	var a [26]int
	for _, r := range s {
		a[r-'a']++
	}
	for i := 0; i < 26; i++ {
		for j := i + 1; j < 26; j++ {
			if a[i] == a[j] && a[i] != 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	if len(os.Args) == 2 {
		fmt.Println(solve(os.Args[1]))
	}
}
