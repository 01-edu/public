package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// counts in the Cartesian coordinates x and y
func countXY(output []rune) (x, y int) {
	countX := 0
	countY := 0
	flag := false
	for _, s := range output {
		if s == '\n' {
			countY++
			flag = true
		} else if !flag {
			countX++
		}
	}
	return countX, countY
}

func determineInput(output []rune) {
	x, y := countXY(output)
	// the leftC and rightC -> position of the corners
	// this will be to see what raid1 is being used
	// for the lower left corner
	leftC := ((x * y) + y - 1) - x
	// for the lower right corner
	rightC := (x * y) + y - 2
	// for the upper right corner
	rightUpC := x - 1

	X := strconv.Itoa(x)
	Y := strconv.Itoa(y)
	result := isPipedWith(output, leftC, rightC, rightUpC, X, Y)
	fmt.Println(strings.Join(result, " || "))
}

func isPipedWith(output []rune, leftC, rightC, rightUpC int, x, y string) []string {
	result := []string{}

	if output[0] == 'o' {
		result = append(result, "[raid1a] ["+x+"] ["+y+"]")
	} else if output[0] == '/' {
		result = append(result, "[raid1b] ["+x+"] ["+y+"]")
	} else if output[0] == 'A' {
		if (output[rightUpC] == 'A' && output[rightC] == 'C') ||
			(output[rightUpC] == 'A' && y == "1") ||
			(x == "1" && y == "1") {
			result = append(result, "[raid1c] ["+x+"] ["+y+"]")
		}
		if (output[rightUpC] == 'C' && output[leftC] == 'A') ||
			(output[leftC] == 'A' && x == "1") ||
			(x == "1" && y == "1") {
			result = append(result, "[raid1d] ["+x+"] ["+y+"]")
		}
		if (output[leftC] == 'C' && x == "1") ||
			(output[rightUpC] == 'C' && y == "1") ||
			(output[rightUpC] == 'C' && output[leftC] == 'C') ||
			(x == "1" && y == "1") {
			result = append(result, "[raid1e] ["+x+"] ["+y+"]")
		}
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var output []rune
	for {
		input, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		output = append(output, input)
	}
	if output[0] != 'o' && output[0] != '/' && output[0] != 'A' {
		fmt.Println("Not a Raid function")
		return
	}
	determineInput(output)
}
