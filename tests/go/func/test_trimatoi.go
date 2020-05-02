package main

import (
	"../lib"
	"./correct"
	"./student"
)

func stringsToTrimAtoi(a []string) []string {
	alpha := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789012345678901234567890123456789"

	for index := 0; index < 4; index++ {
		s := ""
		s += lib.RandStr(lib.RandIntBetween(0, 2), alpha)
		x := lib.RandIntBetween(0, 14)
		if x <= 4 {
			s += "-"
		}
		s += lib.RandStr(lib.RandIntBetween(0, 10), alpha)
		a = append(a, s)
	}
	return a
}

func main() {
	a := []string{
		"",
		"12345",
		"str123ing45",
		"012 345",
		"Hello World!",
		"sd+x1fa2W3s4",
		"sd-x1fa2W3s4",
		"sdx1-fa2W3s4",
		lib.RandAlnum(),
	}
	a = stringsToTrimAtoi(a)
	for _, elem := range a {
		lib.Challenge("TrimAtoi", student.TrimAtoi, correct.TrimAtoi, elem)
	}
}
