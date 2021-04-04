package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/public/test-go/lib/is"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("0")
	} else {
		argument, _ := strconv.Atoi(os.Args[1])

		if argument < 0 {
			fmt.Println("0")
		} else {
			result := 0
			for ; argument >= 0; argument-- {
				if is.Prime(argument) {
					result += argument
				}
			}
			fmt.Println(result)
		}
	}
}
