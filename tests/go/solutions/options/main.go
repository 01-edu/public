package main

import (
	"fmt"
	"os"
)

func printOptions(op [32]int, i int) {
	for _, v := range op {
		if i == 8 {
			fmt.Print(" ")
			i = 0
		}
		fmt.Print(v)
		i++
	}
	fmt.Println()
}

func main() {
	var options [32]int
	pos := 1
	size := len(os.Args)
	if size < 2 {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
	} else {
		for pos < size {
			j := 1
			if os.Args[pos][0] == '-' {
				if os.Args[pos][1] == 'h' {
					fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
					return
				}
				for j < len(os.Args[pos]) && os.Args[pos][j] >= 'a' && os.Args[pos][j] <= 'z' {
					posOption := 'z' - os.Args[pos][j] + 6
					options[posOption] = 1
					j++
				}

				if j < len(os.Args[pos]) && os.Args[pos][j] <= 'a' && os.Args[pos][j] <= 'z' {
					fmt.Println("Invalid Option")
					return
				}
				j++
			}
			pos++
		}
		printOptions(options, 0)
	}
}
