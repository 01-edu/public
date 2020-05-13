package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 4 {
		a, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(0)
			return
		}
		b, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println(0)
			return
		}
		operator := os.Args[2]
		switch operator {
		case "+":
			result := a + b
			if (result > a) != (b > 0) {
				fmt.Println(0)
			} else {
				fmt.Println(result)
			}
		case "-":
			result := a - b
			if (result < a) != (b > 0) {
				fmt.Println(0)
			} else {
				fmt.Println(result)
			}
		case "*":
			result := a * b
			if a != 0 && (result/a != b) {
				fmt.Println(0)
			} else {
				fmt.Println(result)
			}
		case "/":
			if b == 0 {
				fmt.Println("No division by 0")
			} else {
				fmt.Println(a / b)
			}
		case "%":
			if b == 0 {
				fmt.Println("No modulo by 0")
			} else {
				fmt.Println(a % b)
			}
		}
	}
}
