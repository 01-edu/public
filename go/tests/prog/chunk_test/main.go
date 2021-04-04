package main

import (
	"fmt"

	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func chunk(a []int, ch int) {
	slice := []int{}
	if ch <= 0 {
		fmt.Println()
		return
	}
	result := make([][]int, 0, len(a)/ch+1)
	for len(a) >= ch {
		slice, a = a[:ch], a[ch:]
		result = append(result, slice)
	}
	if len(a) > 0 {
		result = append(result, a[:])
	}
	fmt.Println(result)
}

func randomSize() []int {
	randSlice := []int{}
	for i := 0; i <= lib.RandIntBetween(0, 20); i++ {
		randSlice = append(randSlice, lib.RandInt())
	}
	return randSlice
}

func main() {
	type node struct {
		slice []int
		ch    int
	}
	table := []node{}

	for i := 0; i <= 7; i++ {
		value := node{
			slice: randomSize(),
			ch:    lib.RandIntBetween(0, 10),
		}
		table = append(table, value)
	}
	table = append(table, node{
		slice: []int{},
		ch:    0,
	}, node{
		slice: []int{1, 2, 3, 4, 5, 6, 7, 8},
		ch:    0,
	})
	for _, args := range table {
		lib.Challenge("Chunk", student.Chunk, chunk, args.slice, args.ch)
	}
}
