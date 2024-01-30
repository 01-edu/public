package main

import (
	"fmt"
	"piscine"
)

func main() {
	table := []int{1, 2, 3}
	ac := 93
	piscine.FoldInt(Add, table, ac)
	piscine.FoldInt(Mul, table, ac)
	piscine.FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	piscine.FoldInt(Add, table, ac)
	piscine.FoldInt(Mul, table, ac)
	piscine.FoldInt(Sub, table, ac)
}
