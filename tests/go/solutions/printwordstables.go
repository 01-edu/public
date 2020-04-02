package solutions

import (
	"fmt"
)

func PrintWordsTables(table []string) {
	for _, t := range table {
		fmt.Println(t)
	}
}

func isSeparator(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}
