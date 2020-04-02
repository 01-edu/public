package student_test

import (
	"testing"

	solutions "./solutions"
	student "./student"
	"github.com/01-edu/z01"
)

func TestTrimAtoi(t *testing.T) {
	array := []string{"",
		"12345",
		"str123ing45",
		"012 345",
		"Hello World!",
		"sd+x1fa2W3s4",
		"sd-x1fa2W3s4",
		"sdx1-fa2W3s4",
		z01.RandAlnum()}
	array = stringsToTrimAtoi(array)
	for i := range array {
		z01.Challenge(t, student.TrimAtoi, solutions.TrimAtoi, array[i])
	}
}

func stringsToTrimAtoi(arr []string) []string {
	alpha := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789012345678901234567890123456789"

	for index := 0; index < 4; index++ {
		str := ""
		str += z01.RandStr(z01.RandIntBetween(0, 2), alpha)
		x := z01.RandIntBetween(0, 14)
		if x <= 4 {
			str += "-"
		}
		str += z01.RandStr(z01.RandIntBetween(0, 10), alpha)
		arr = append(arr, str)
	}
	return arr
}
