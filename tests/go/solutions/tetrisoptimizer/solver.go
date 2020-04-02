package main

func recursion(b board, array []tetrimino, idx int) bool {
	if idx == len(array) {
		return true
	}
	for i := range b {
		for j := range b[i] {
			if !b.check(i, j, array[idx]) {
				continue
			}
			b.put(i, j, idx, array[idx])
			if recursion(b, array, idx+1) {
				return true
			}
			b.remove(i, j, array[idx])
		}
	}
	return false
}

func solve(array []tetrimino) board {
	square := 2
	for square*square < len(array)*4 {
		square++
	}
	for {
		b := makeBoard(square)
		square++
		if !recursion(b, array, 0) {
			continue
		}
		return b
	}
}
