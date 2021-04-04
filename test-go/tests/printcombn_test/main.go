package main

import (
	"fmt"

	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func show(n int, table [9]int, tmax [9]int) {
	i := 0
	for i < n {
		fmt.Print(table[i])
		i++
	}
	if table[0] != tmax[0] {
		fmt.Print(", ")
	}
}

func printComb1() {
	table := [9]int{0}
	tmax := [9]int{9}
	for table[0] <= tmax[0] {
		show(1, table, tmax)
		table[0]++
	}
}

func printCombN(n int) {
	table := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	tmax := [9]int{}

	if n == 1 {
		printComb1()
	} else {
		i := n - 1
		j := 9
		for i >= 0 {
			tmax[i] = j
			i--
			j--
		}
		i = n - 1
		for table[0] < tmax[0] {
			if table[i] != tmax[i] {
				show(n, table, tmax)
				table[i]++
			}
			if table[i] == tmax[i] {
				if table[i-1] != tmax[i-1] {
					show(n, table, tmax)
					table[i-1]++
					j = i
					for j < n {
						table[j] = table[j-1] + 1
						j++
					}
					i = n - 1
				}
			}
			for table[i] == tmax[i] && table[i-1] == tmax[i-1] && i > 1 {
				i--
			}
		}
		show(n, table, tmax)
	}
	fmt.Println()
}

func main() {
	table := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, arg := range table {
		lib.Challenge("PrintCombN", student.PrintCombN, printCombN, arg)
	}
}
