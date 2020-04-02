package main

import (
	"strconv"
	"strings"
	"testing"

	"github.com/01-edu/z01"
)

func isAPrime(nb int) bool {
	if nb <= 0 || nb == 1 {
		return false
	}
	if nb <= 3 {
		return true

	} else if nb%2 == 0 || nb%3 == 0 {
		return false
	}

	i := 5
	for i*i <= nb {
		if nb%i == 0 || nb%(i+2) == 0 {
			return false
		}
		i = i + 6
	}
	return true
}

func TestAddPrimeSum(t *testing.T) {

	var table []string

	// fill with all rpime numbers between 0 and 100

	for i := 0; i < 100; i++ {
		if isAPrime(i) {
			str := strconv.Itoa(i)
			table = append(table, str)
		}
	}

	// adds 15 random numbers
	for i := 0; i < 15; i++ {
		table = append(table, strconv.Itoa(z01.RandIntBetween(1, 10000)))
	}

	// special cases
	table = append(table, "\"\"")
	table = append(table, "1 2")

	for _, s := range table {
		z01.ChallengeMainExam(t, strings.Fields(s)...)
	}
}
