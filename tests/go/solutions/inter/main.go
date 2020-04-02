package main

import (
	"fmt"
	"os"
	"strings"
)

func result(str1 string, str2 string) string {

	arraystr1 := []rune(str1)
	arraystr2 := []rune(str2)
	sizeStr1 := len(arraystr1)
	sizeStr2 := len(arraystr2)
	var rest []rune

	for i := 0; i < sizeStr1; i++ {
		for j := 0; j < sizeStr2; j++ {
			if arraystr1[i] == arraystr2[j] && !strings.ContainsRune(string(rest), arraystr1[i]) {
				rest = append(rest, arraystr1[i])
			}
		}
	}
	return string(rest)
}

func main() {

	if len(os.Args) == 3 {
		fmt.Println(result(os.Args[1], os.Args[2]))
	} else {
		fmt.Println()
	}
}
