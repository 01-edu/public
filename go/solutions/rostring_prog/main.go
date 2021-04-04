package main

import (
	"fmt"
	"os"
	"strings"
)

func deleteExtraSpaces(a []string) []string {
	var res []string
	for _, v := range a {
		if v != "" {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	if len(os.Args) == 2 {
		words := strings.Split(os.Args[1], " ")
		words = deleteExtraSpaces(words)
		if len(words) >= 1 {
			for _, v := range words[1:] {
				fmt.Print(v, " ")
			}
			fmt.Print(words[0])
		}
	}
	fmt.Println()
}
