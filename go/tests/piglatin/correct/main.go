package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "" {
		return
	}
	s := os.Args[1]
	vowel := strings.IndexAny(s, "aeiouAEIOU")
	if vowel == -1 {
		fmt.Println("No vowels")
	} else {
		fmt.Println(string(s[vowel:]) + string(s[:vowel]) + "ay")
	}
}
