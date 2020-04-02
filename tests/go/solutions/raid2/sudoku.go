package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

// Prints in standard output the sudoku board
func printBoard(board [][]rune) {
	for _, row := range board {
		for i, e := range row {
			z01.PrintRune(e)
			if i != len(row)-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

// Return true if the value 'value' is in the row 'y' of the board 'board'
func isInRow(board [][]rune, value rune, x, y int) bool {
	for i, v := range board[y] {
		if i != x && v == value {
			return true
		}
	}
	return false
}

// Returns true if the value 'value' is in the column 'x' of the board 'board'
func isInColumn(board [][]rune, value rune, x, y int) bool {
	for i := 0; i < 9; i++ {
		if i != y && board[i][x] == value {
			return true
		}
	}
	return false
}

// Receives an int 'x' and returns the beginning and the end
// of the interval of consecutive pairs of multiples of three 'x' is in.
// Ex. for x = 2, x is in (0, 3)
// for x = 4, x is in (3, 6)
func intervalThree(x int, max int) (int, int) {
	var i int
	for i = 0; i < max; i += 3 {
		if x >= i && x < i+3 {
			break
		}
	}
	endi := 3 + i
	return i, endi
}

// Returns true if the value 'value' is allowed in the possition (x,y) of the board 'board'
func isAllowedInBox(board [][]rune, value rune, x, y int) bool {
	n := len(board)
	begi, endi := intervalThree(x, n)
	begj, endj := intervalThree(y, n)
	for j := begj; j < endj; j++ {
		for i := begi; i < endi; i++ {
			if (j != y || i != i) && board[j][i] == value {
				return false
			}
		}
	}
	return true
}

// Returns true if the value 'val' is allowed in the position (x, y) of the board
func isAllowed(board [][]rune, val rune, x, y int) bool {
	return !isInRow(board, val, x, y) &&
		!isInColumn(board, val, x, y) &&
		isAllowedInBox(board, val, x, y)
}

// Returns true if the position doesn't have any value defined
// That is, if the character is a dot '.'
func isEmpty(board [][]rune, x, y int) bool {
	return board[y][x] == '.'
}

// Returns all the empty positions in the board
func availablePos(board [][]rune) [][]int {
	var ava [][]int
	for y, row := range board {
		for x, e := range row {
			if e == '.' {
				ava = append(ava, []int{x, y})
			}
		}
	}
	return ava
}

func validBoard(board [][]rune) bool {
	size := 9
	if len(board) != size {
		return false
	}
	for y, row := range board {
		if len(row) != size {
			return false
		}
		for x, e := range row {
			if (e < '1' || e > '9') && e != '.' {
				return false
			}
			if e != '.' && !isAllowed(board, e, x, y) {
				return false
			}
		}
	}
	return true
}

// Returns true if the sudoku has a solution
func sudokuH(board [][]rune, available [][]int, i int) bool {
	n := len(available)

	if i >= n {
		return true
	}

	x := available[i][0]
	y := available[i][1]

	for c := '1'; c <= '9'; c++ {
		if isAllowed(board, c, x, y) {
			board[y][x] = c
			if sudokuH(board, available, i+1) {
				return true
			}
			board[y][x] = '.'
		}
	}
	return false
}

// Receives a board and fills the empty positions with the correct
// value
func solveSudoku(board [][]rune) {
	available := availablePos(board)
	if sudokuH(board, available, 0) {
		printBoard(board)
	} else {
		fmt.Println("Error")
	}
}

func main() {
	var board [][]rune

	for _, v := range os.Args[1:] {
		board = append(board, []rune(v))
	}

	if validBoard(board) {
		solveSudoku(board)
	} else {
		fmt.Println("Error")
	}
}
