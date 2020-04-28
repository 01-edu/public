package correct

import (
	"fmt"
	"unicode"
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

func printLine(arr [10]int, start int) {
	size := len(arr)
	a := start
	var aux, b int

	for a < start+16 && a < size {
		if a%4 == 0 && a != 0 {
			fmt.Println()
		}
		b = 8 - printBase(arr[a])
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
		if unicode.IsPrint(rune(arr[c])) {
			fmt.Printf("%c", rune(arr[c]))
		} else {
			fmt.Print(".")
		}
		c++
	}
	fmt.Println()
}

func PrintMemory(a [10]int) {
	i := 0
	for i < len(a) {
		printLine(a, i)
		i += 16
	}
}
