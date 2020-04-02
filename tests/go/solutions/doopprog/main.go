package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 4 {
		var result int
		firstArg, err := strconv.Atoi(os.Args[1])

		if err != nil {
			fmt.Println(0)
			return
		}

		operator := os.Args[2]
		secondArg, err1 := strconv.Atoi(os.Args[3])

		if err1 != nil {
			fmt.Println(0)
			return
		}

		if secondArg == 0 && operator == "/" {
			fmt.Println("No division by 0")
			return
		} else if secondArg == 0 && operator == "%" {
			fmt.Println("No modulo by 0")
			return
		} else if operator == "+" {
			result = firstArg + secondArg
			if !((result > firstArg) == (secondArg > 0)) {
				fmt.Println(0)
				return
			}
		} else if operator == "-" {
			result = firstArg - secondArg
			if !((result < firstArg) == (secondArg > 0)) {
				fmt.Println(0)
				return
			}
		} else if operator == "/" {
			result = firstArg / secondArg
		} else if operator == "*" {
			result = firstArg * secondArg
			if firstArg != 0 && (result/firstArg != secondArg) {
				fmt.Println(0)
				return
			}
		} else if operator == "%" {
			result = firstArg % secondArg
		}
		fmt.Println(result)
	}
}
