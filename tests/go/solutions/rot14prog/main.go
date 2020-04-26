package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	var result string

	for _, r := range os.Args[1] {
		if r >= 'a' && r <= 'z' {
			if r >= 'm' {
				r -= 12
			} else {
				r += 14
			}
		} else if r >= 'A' && r <= 'Z' {
			if r >= 'M' {
				r -= 12
			} else {
				r += 14
			}
		}
		result += string(r)
	}
	fmt.Println(result)
}
