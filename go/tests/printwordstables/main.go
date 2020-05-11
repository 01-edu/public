package main

import (
	"fmt"

	"./student"
	"github.com/01-edu/public/go/lib"
)

func printWordsTables(a []string) {
	for _, s := range a {
		fmt.Println(s)
	}
}

func main() {
	table := [][]string{}
	// 30 random slice of slice of strings

	for i := 0; i < 30; i++ {
		table = append(table, lib.MultRandASCII())
	}

	table = append(table,
		[]string{"Hello", "how", "are", "you?"})

	for _, arg := range table {
		lib.Challenge("PrintWordsTables", student.PrintWordsTables, printWordsTables, arg)
	}
}
