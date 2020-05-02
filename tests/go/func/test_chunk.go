package main

import (
	"../lib"
	"./correct"
	"./student"
)

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
		lib.Challenge("Chunk", student.Chunk, correct.Chunk, args.slice, args.ch)
	}
}
