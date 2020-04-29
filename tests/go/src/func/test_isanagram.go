package main

import (
	"../lib"
	"./correct"
	"./student"
)

func main() {
	table := [][2]string{
		{"listen", "silent"},
		{"alem", "school"},
		{"neat", "a net"},
		{"anna madrigal", "a man and a girl"},
		{"abcc", "abcd"},
		{"aaaac", "caaaa"},
		{"", ""},
		{"       ", ""},
		{"lyam", "meow"},
		{"golang", "lang go"},
		{"verylongword", "v e r y l o n g w o r d"},
		{"chess", "ches"},
		{"anagram", "nnagram"},
		{"chess", "board"},
		{"mmm", "m"},
		{"pulp", "fiction"},
	}
	for i := 0; i < 15; i++ {
		table = append(table, [2]string{
			lib.RandStr(lib.RandIntBetween(15, 20), "qwertyuiopasdfghjklzxcvbnm "),
			lib.RandStr(lib.RandIntBetween(15, 20), "qwertyuiopasdfghjklzxcvbnm "),
		})
	}

	for _, arg := range table {
		lib.Challenge("IsAnagram", student.IsAnagram, correct.IsAnagram, arg[0], arg[1])
	}
}
