package main

import (
	"fmt"
	"os"
	"strings"
)

func deleteExtraSpaces(arr []string) []string {
	var res []string
	for _, v := range arr {
		if v != "" {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		os.Exit(0)
	}

	arg := strings.Split(os.Args[1], " ")
	arg = deleteExtraSpaces(arg)
	for i, v := range arg {
		fmt.Print(v)
		if i < len(arg)-1 {
			fmt.Print("   ")
		}
	}
	fmt.Println()
}
