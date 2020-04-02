package main

import (
	"testing"

	solutions "../../solutions"
	"github.com/01-edu/z01"
)

func randomSize() []int {
	randSlice := []int{}
	for i := 0; i <= z01.RandIntBetween(0, 20); i++ {
		randSlice = append(randSlice, z01.RandInt())
	}
	return randSlice
}

func TestChunk(t *testing.T) {
	type node struct {
		slice []int
		ch    int
	}
	table := []node{}

	for i := 0; i <= 7; i++ {
		value := node{
			slice: randomSize(),
			ch:    z01.RandIntBetween(0, 10),
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
		z01.Challenge(t, Chunk, solutions.Chunk, args.slice, args.ch)
	}
}
