package main

import (
	student "student"

	"github.com/01-edu/public/test-go/lib"
)

func max(ints ...int) int {
	if len(ints) == 0 {
		panic("max() is invalid")
	}
	max := ints[0]
	for _, i := range ints[1:] {
		if i > max {
			max = i
		}
	}
	return max
}

func reviveThreeNums(a, b, c, d int) int {
	maxi := -111
	if a != max(a, b, c, d) {
		maxi = max(maxi, max(a, b, c, d)-a)
	}
	if b != max(a, b, c, d) {
		maxi = max(maxi, max(a, b, c, d)-b)
	}
	if c != max(a, b, c, d) {
		maxi = max(maxi, max(a, b, c, d)-c)
	}
	if d != max(a, b, c, d) {
		maxi = max(maxi, max(a, b, c, d)-d)
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
