package main

import (
	"fmt"
	"os"
	"strings"
)

func deleteExtraSpaces(a []string) (res []string) {
	for _, v := range a {
		if v != "" {
			res = append(res, v)
		}
	}
	return
}

func main() {
	if len(os.Args) == 2 {
		arg := strings.Split(os.Args[1], " ")
		arg = deleteExtraSpaces(arg)
		fmt.Println(strings.Join(arg, "   "))
	}
}
