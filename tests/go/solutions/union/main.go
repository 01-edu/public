package main

import (
	"fmt"
	"os"
)

func isIn(a rune, arr []rune) bool {
	for _, v := range arr {
		if a == v {
			return true
		}
	}
	return false
}

func main() {

	if len(os.Args) == 3 {
		var res []rune
		str1 := []rune(os.Args[1])
		str2 := []rune(os.Args[2])

		for _, v := range str1 {
			if !isIn(v, res) {
				res = append(res, v)
			}
		}

		for _, v := range str2 {
			if !isIn(v, res) {
				res = append(res, v)
			}
		}

		fmt.Print(string(res))
	}

	fmt.Println()
}
