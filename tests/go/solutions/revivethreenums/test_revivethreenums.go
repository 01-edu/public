package main

import (
	"github.com/01-edu/z01"

	solutions "../../solutions"
)

func main() {
	table := [][4]int{
		{3, 6, 5, 4},
		{40, 40, 40, 60},
		{201, 101, 101, 200},
	}

	for i := 0; i < 25; i++ {
		first := z01.RandIntBetween(0, 877)
		second := z01.RandIntBetween(0, 877)
		third := z01.RandIntBetween(0, 877)
		table = append(table, [4]int{
			first + second,
			second + third,
			first + third,
			first + second + third,
		})
	}
	for _, arg := range table {
		z01.Challenge("Revivethreenums", ReviveThreeNums, solutions.ReviveThreeNums, arg[0], arg[1], arg[2], arg[3])
	}
}
