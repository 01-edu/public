package main

import (
	"./student"
	"github.com/01-edu/public/go/lib"
)

func more(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func max(a, b, c, d int) int {
	if a >= b && a >= c && a >= d {
		return a
	}
	if b >= a && b >= c && b >= d {
		return b
	}
	if c >= a && c >= b && c >= d {
		return c
	}
	if d >= a && d >= b && d >= c {
		return d
	}
	return -1
}

func reviveThreeNums(a, b, c, d int) int {
	maxi := -111
	if a != max(a, b, c, d) {
		maxi = more(maxi, max(a, b, c, d)-a)
	}
	if b != max(a, b, c, d) {
		maxi = more(maxi, max(a, b, c, d)-b)
	}
	if c != max(a, b, c, d) {
		maxi = more(maxi, max(a, b, c, d)-c)
	}
	if d != max(a, b, c, d) {
		maxi = more(maxi, max(a, b, c, d)-d)
	}
	return maxi
}

func main() {
	table := [][4]int{
		{3, 6, 5, 4},
		{40, 40, 40, 60},
		{201, 101, 101, 200},
	}

	for i := 0; i < 25; i++ {
		first := lib.RandIntBetween(0, 877)
		second := lib.RandIntBetween(0, 877)
		third := lib.RandIntBetween(0, 877)
		table = append(table, [4]int{
			first + second,
			second + third,
			first + third,
			first + second + third,
		})
	}
	for _, arg := range table {
		lib.Challenge("Revivethreenums", student.ReviveThreeNums, reviveThreeNums, arg[0], arg[1], arg[2], arg[3])
	}
}
