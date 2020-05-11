package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isOp(s string) bool {
	return s == "+" ||
		s == "-" ||
		s == "*" ||
		s == "/" ||
		s == "%"
}

func deleteExtraSpaces(arr []string) []string {
	var res []string
	for _, v := range arr {
		if v != "" {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	if len(os.Args) == 2 {
		var values []int
		var n int
		op := strings.Split(os.Args[1], " ")
		op = deleteExtraSpaces(op)
		for _, v := range op {
			val, err := strconv.Atoi(v)

			if err == nil {
				values = append(values, val)
				continue
			}

			n = len(values)
			if isOp(v) && n < 2 {
				fmt.Println("Error")
				os.Exit(0)
			}

			switch v {
			case "+":
				values[n-2] += values[n-1]
				values = values[:n-1]
			case "-":
				values[n-2] -= values[n-1]
				values = values[:n-1]
			case "*":
				values[n-2] *= values[n-1]
				values = values[:n-1]
			case "/":
				values[n-2] /= values[n-1]
				values = values[:n-1]
			case "%":
				values[n-2] %= values[n-1]
				values = values[:n-1]
			default:
				fmt.Println("Error")
				os.Exit(0)
			}
		}
		if len(values) == 1 {
			fmt.Println(values[0])
		} else {
			fmt.Println("Error")
		}
	} else {
		fmt.Println("Error")
	}
}
