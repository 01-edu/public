package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
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

	s := lib.MultRandWords()

	arr = append(arr, []interface{}{s, -len(s) - 10, -len(s) - 5})

	for i := 0; i < 3; i++ {
		s = lib.MultRandWords()
		arr = append(arr, []interface{}{s, lib.RandIntBetween(-len(s)-10, len(s)+10), lib.RandIntBetween(-len(s)-8, len(s)+10)})
	}

	for _, a := range arr {
		lib.Challenge("Slice", student.Slice, correct.Slice, a...)
	}
}
