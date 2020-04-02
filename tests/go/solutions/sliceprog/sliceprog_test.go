package main

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "../../solutions"
)

func TestSlice(t *testing.T) {

	arr := [][]interface{}{
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			1,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			-3,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			2, 4,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			-2, -1,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			2, 0,
		},
	}

	s := z01.MultRandWords()

	arr = append(arr, []interface{}{s, -len(s) - 10, -len(s) - 5})

	for i := 0; i < 3; i++ {
		s = z01.MultRandWords()
		arr = append(arr, []interface{}{s, z01.RandIntBetween(-len(s)-10, len(s)+10), z01.RandIntBetween(-len(s)-8, len(s)+10)})
	}

	for _, a := range arr {
		z01.Challenge(t, Slice, solutions.Slice, a...)
	}
}
