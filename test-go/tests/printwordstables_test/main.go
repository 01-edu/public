package main

import (
	"fmt"

	student "student"

	"github.com/01-edu/public/test-go/lib"
)

func printWordsTables(a []string) {
	for _, s := range a {
		fmt.Println(s)
	}
}

func main() {
	table := [][]string{{"Hello", "how", "are", "you?"}}

	// 30 random slice of slice of strings
	for i := 0; i < 30; i++ {
		table = append(table, lib.MultRandASCII())
	}

	for _, arg := range table {
		lib.Challenge("PrintWordsTables", student.PrintWordsTables, printWordsTables, arg)
	}
}
