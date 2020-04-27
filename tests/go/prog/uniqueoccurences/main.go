package main

import (
	"fmt"
	"os"
)

func solve(str string) bool {
	arr := make([]int, 26)
	for i := 0; i < len(str); i++ {
		arr[str[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		for j := i + 1; j < 26; j++ {
			if arr[i] == arr[j] && arr[i] != 0 {
				return false
			}
		}
	}
	return true
}

func print(str string) {
	for _, v := range str {
		fmt.Printf("%c", v)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println()
		return
	}
	result := solve(args[0])
	if result {
		print("true\n")
	} else {
		print("false\n")
	}
}
