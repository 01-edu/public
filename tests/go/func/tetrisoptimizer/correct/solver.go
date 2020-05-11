package main

func recursion(b board, a []tetrimino, idx int) bool {
	if idx == len(a) {
		return true
	}
	for i := range b {
		for j := range b[i] {
			if !b.check(i, j, a[idx]) {
				continue
			}
			b.put(i, j, idx, a[idx])
			if recursion(b, a, idx+1) {
				return true
			}
			b.remove(i, j, a[idx])
		}
	}
	return false
}

func solve(a []tetrimino) board {
	square := 2
	for square*square < len(a)*4 {
		square++
	}
	for {
		b := makeBoard(square)
		square++
		if !recursion(b, a, 0) {
			continue
		}
		return b
	}
}
