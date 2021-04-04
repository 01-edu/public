package main

import (
	student "student"

	"github.com/01-edu/public/go/tests/lib"
)

func isAnagram(s, t string) bool {
	alph := make([]int, 26)
	for i := 0; i < len(s); i++ {
		if s[i] < 'a' || s[i] > 'z' {
			continue
		}
		alph[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		if t[i] < 'a' || t[i] > 'z' {
			continue
		}
		alph[t[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if alph[i] != 0 {
			return false
		}
	}
	return true
}

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
		lib.Challenge("IsAnagram", student.IsAnagram, isAnagram, arg[0], arg[1])
	}
}
