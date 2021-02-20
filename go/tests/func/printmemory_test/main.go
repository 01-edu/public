package main

import (
	"fmt"
	"unicode"

	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func printBase(nbr int) int {
	var result []rune
	base := "0123456789abcdef"
	typeBase := []rune(base)
	a := 0
	pos := 0
	i := 0
	if nbr == 0 {
		result = append(result, '0', '0')
		i = 2
	}
	for nbr > 0 {
		pos = nbr % 16
		nbr = nbr / 16
		result = append(result, typeBase[pos])
		i++
	}
	if i == 1 {
		result = append(result, '0')
		i = 2
	}
	for j := i - 1; j >= 0; j-- {
		fmt.Printf("%c", result[j])
		a++
	}
	return a
}

func printLine(elems [10]int, start int) {
	size := len(elems)
	a := start
	var aux, b int

	for a < start+16 && a < size {
		if a%4 == 0 && a != 0 {
			fmt.Println()
		}
		b = 8 - printBase(elems[a])
		for aux != b {
			if b == 6 {
				fmt.Print("0")
			}
			if aux == 1 {
				fmt.Print(" ")
			}
			if b < 6 {
				fmt.Print("0")
			}
			aux++
		}
		fmt.Print(" ")
		aux = 0
		a++
	}
	fmt.Println()
	c := start
	for c < start+16 && c < size {
		if unicode.IsGraphic(rune(elems[c])) {
			fmt.Printf("%c", rune(elems[c]))
		} else {
			fmt.Print(".")
		}
		c++
	}
	fmt.Println()
}

func printMemory(a [10]int) {
	i := 0
	for i < len(a) {
		printLine(a, i)
		i += 16
	}
}

func main() {
	var table [10]int

	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			table[i] = lib.RandIntBetween(0, 1000)
		}
		lib.Challenge("PrintMemory", student.PrintMemory, printMemory, table)
	}
	table2 := [10]int{104, 101, 108, 108, 111, 16, 21, 42}
	lib.Challenge("PrintMemory", student.PrintMemory, printMemory, table2)
}

// TODO: this can be simplified a lot
