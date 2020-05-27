package main

import (
	"fmt"
	"os"
)

func ok(s1 string, s2 string) bool {
	runes1 := []rune(s1)
	runes2 := []rune(s2)
	var rest string
	count := 0
	for i := 0; i < len(runes1); i++ {
		for j := count; j < len(runes2); j++ {
			if runes1[i] == runes2[j] {
				rest += string(runes1[i])
				j = len(runes2) - 1
			}
			count++
		}
	}
	return s1 == rest
}

func main() {
	if len(os.Args) == 3 {
		if ok(os.Args[1], os.Args[2]) {
			fmt.Println(os.Args[1])
		}
	}
}
