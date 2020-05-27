package main

import "fmt"

type board [][]byte

func makeBoard(size int) board {
	var b board
	for i := 0; i < size; i++ {
		row := make([]byte, size)
		for j := range row {
			row[j] = dot
		}
		b = append(b, row)
	}
	return b
}

func (b board) print() {
	for i := range b {
		fmt.Println(string(b[i]))
	}
}

func (b board) check(i, j int, t tetrimino) bool {
	for y := range t {
		for x := range t[y] {
			if t[y][x] != hashTag {
				continue
			}
			if i+y >= len(b) || j+x >= len(b[i+y]) {
				return false
			}
			if b[i+y][j+x] != dot {
				return false
			}
		}
	}
	return true
}

func (b board) put(i, j, idx int, t tetrimino) {
	for y := range t {
		for x := range t[y] {
			if t[y][x] != hashTag {
				continue
			}
			b[i+y][j+x] = byte(idx + 'A')
		}
	}
}

func (b board) remove(i, j int, t tetrimino) {
	for y := range t {
		for x := range t[y] {
			if t[y][x] != hashTag {
				continue
			}
			b[i+y][j+x] = dot
		}
	}
}
