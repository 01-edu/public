package solutions

import (
	"github.com/01-edu/z01"
)

const size = 8

// board is a chessboard composed of boolean squares, a true square means a queen is on it
// a false square means it is a free square
var board [size][size]bool

// goodDirection check that there is no queen on the segment that starts at (x, y)
// coordinates, points into the direction vector (vx, vy) and ends at the edge of the board
func goodDirection(x, y, vx, vy int) bool {
	// x and y are still on board
	for 0 <= x && x < size &&
		0 <= y && y < size {
		if board[x][y] {
			// Not a good line : the square is already occupied
			return false
		}
		x = x + vx // Move x in the right direction
		y = y + vy // Move y in the right direction
	}
	// All clear
	return true
}

// goodSquare makes all the necessary line checks for the queens movements
func goodSquare(x, y int) bool {
	return goodDirection(x, y, +0, -1) &&
		goodDirection(x, y, +1, -1) &&
		goodDirection(x, y, +1, +0) &&
		goodDirection(x, y, +1, +1) &&
		goodDirection(x, y, +0, +1) &&
		goodDirection(x, y, -1, +1) &&
		goodDirection(x, y, -1, +0) &&
		goodDirection(x, y, -1, -1)
}

func printQueens() {
	x := 0
	for x < size {
		y := 0
		for y < size {
			if board[x][y] {
				// We have found a queen, let's print her y
				z01.PrintRune(rune(y) + '1')
			}
			y++
		}
		x++
	}
	z01.PrintRune('\n')
}

// tryX tries, for a given x (column) to find a y (row) so that the queen on (x, y) is a part
// of the solution to the problem
func tryX(x int) {
	y := 0
	for y < size {
		if goodSquare(x, y) {
			// Since the square is good for the queen, let's put one on it:
			board[x][y] = true

			if x == size-1 {
				// x is the biggest possible x, it means that we just placed the last
				// queen on the board, so the solution is complete and we can print it
				printQueens()
			} else {
				// let's try to put another queen on the next empty x (column)
				tryX(x + 1)
			}

			// remove the queen of the board, to try other y values
			board[x][y] = false
		}
		y++
	}
}

func EightQueens() {
	// try the first column
	tryX(0)
}
