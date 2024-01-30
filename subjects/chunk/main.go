package main

import (
	"piscine"
)

func main() {
	piscine.Chunk([]int{}, 10)
	piscine.Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	piscine.Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	piscine.Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	piscine.Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
