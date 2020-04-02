package main

import (
	"fmt"
	"os"
)

func result(str1 string, str2 string) string {
	arraystr1 := []rune(str1)
	arraystr2 := []rune(str2)
	var rest string
	count := 0
	for i := 0; i < len(arraystr1); i++ {
		for j := count; j < len(arraystr2); j++ {
			if arraystr1[i] == arraystr2[j] {
				rest += string(arraystr1[i])
				j = len(arraystr2) - 1
			}
			count++
		}
	}
	if rest != str1 {
		return ""
	}
	return rest
}

func main() {
	if len(os.Args) == 3 {
		fmt.Println(result(os.Args[1], os.Args[2]))
	} else {
		fmt.Println()
	}
}
