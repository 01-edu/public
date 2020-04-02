package main

import (
	"testing"

	"github.com/01-edu/z01"

	solutions "../../solutions" // This line is not necessary when testing an exercise with a program
)

func TestIsAnagram(t *testing.T) {
	type node struct {
		s string
		t string
	}

	table := []node{}

	table = append(table,
		node{"listen", "silent"},
		node{"alem", "school"},
		node{"neat", "a net"},
		node{"anna madrigal", "a man and a girl"},
		node{"abcc", "abcd"},
		node{"aaaac", "caaaa"},
		node{"", ""},
		node{"       ", ""},
		node{"lyam", "meow"},
		node{"golang", "lang go"},
		node{"verylongword", "v e r y l o n g w o r d"},
		node{"chess", "ches"},
		node{"anagram", "nnagram"},
		node{"chess", "board"},
		node{"mmm", "m"},
		node{"pulp", "fiction"},
	)

	for i := 0; i < 15; i++ {
		value := node {
			s: z01.RandStr(z01.RandIntBetween(15, 20), "qwertyuiopasdfghjklzxcvbnm "),
			t: z01.RandStr(z01.RandIntBetween(15, 20), "qwertyuiopasdfghjklzxcvbnm "),
		}

		table = append(table, value)
	}

	for _, arg := range(table) {
		z01.Challenge(t, IsAnagram, solutions.IsAnagram, arg.s, arg.t)
	}
}