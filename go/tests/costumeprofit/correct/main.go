package main

import (
	"fmt"
	"os"
	"strconv"
)

func min(numbers ...int) int {
	min := numbers[0]
	for _, i := range numbers {
		if i < min {
			min = i
		}
	}
	return min
}

func max(numbers ...int) int {
	max := numbers[0]
	for _, i := range numbers {
		if i > max {
			max = i
		}
	}
	return max
}

func costumeProfit(a, b, c, d, e, f int) int {
	if d == min(a, b, c, d) {
		return d * max(e, f)
	}
	if e > f {
		ans := min(a, d) * e
		d -= min(a, d)
		ans += min(d, b, c) * f
		return ans
	}
	ans := min(b, c, d) * f
	d -= min(b, c, d)
	ans += min(a, d) * e
	return ans
}

func main() {
	args := os.Args[1:]
	a, _ := strconv.Atoi(args[0])
	b, _ := strconv.Atoi(args[1])
	c, _ := strconv.Atoi(args[2])
	d, _ := strconv.Atoi(args[3])
	e, _ := strconv.Atoi(args[4])
	f, _ := strconv.Atoi(args[5])

	fmt.Println(costumeProfit(a, b, c, d, e, f))
}
