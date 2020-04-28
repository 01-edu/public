package correct

import (
	"unicode"

	"github.com/01-edu/z01"
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
		z01.PrintRune(result[j])
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
			z01.PrintRune('\n')
		}
		b = 8 - printBase(arr[a])
		for aux != b {
			if b == 6 {
				z01.PrintRune('0')
			}
			if aux == 1 {
				z01.PrintRune(' ')
			}
			if b < 6 {
				z01.PrintRune('0')
			}
			aux++
		}
		z01.PrintRune(' ')
		aux = 0
		a++
	}
	z01.PrintRune('\n')
	c := start
	for c < start+16 && c < size {
		if unicode.IsPrint(rune(arr[c])) {
			z01.PrintRune(rune(arr[c]))
		} else {
			z01.PrintRune('.')
		}
		c++
	}
	z01.PrintRune('\n')
}

func PrintMemory(arr [10]int) {
	i := 0
	for i < len(arr) {
		printLine(arr, i)
		i += 16
	}
}
