package main

import (
	"fmt"
	"os"
	"strconv"
)

//This solution is the comparing file of the staff
// Because the solution is a function,
//
//1) here the package is solutions
//2) it does not need an empty func main(){}
//3) its location is 1 level below the folder of the nauuo_test.go file

func min(numbers ...int) int {
	minVal := numbers[0]
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < minVal {
			minVal = numbers[i]
		}
	}
	return minVal
}

func max(numbers ...int) int {
	maxVal := numbers[0]
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > maxVal {
			maxVal = numbers[i]
		}
	}
	return maxVal
}

func Costume_profit(a, b, c, d, e, f int) int {
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

	fmt.Println(Costume_profit(a, b, c, d, e, f))
}
