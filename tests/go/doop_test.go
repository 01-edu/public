package student_test

import (
	"strings"
	"testing"

	"strconv"

	"github.com/01-edu/z01"
)

func TestDoop(t *testing.T) {

	operatorsTable := []string{"+", "-", "*", "/", "%"}

	table := []string{}

	for i := 0; i < 4; i++ {
		firstArg := strconv.Itoa(z01.RandIntBetween(-1000, 1000))
		secondArg := strconv.Itoa(z01.RandIntBetween(0, 1000))

		for _, operator := range operatorsTable {
			table = append(table, firstArg+" "+operator+" "+secondArg)

		}
	}

	table = append(table, "1 + 1")
	table = append(table, "hello + 1")
	table = append(table, "1 p 1")
	table = append(table, "1 / 0")
	table = append(table, "1 # 1")
	table = append(table, "1 % 0")
	table = append(table, "1 * 1")
	table = append(table, "1argument")
	table = append(table, "2 arguments")
	table = append(table, "4 arguments so invalid")
	table = append(table, "9223372036854775807 + 1")
	table = append(table, "9223372036854775809 - 3")
	table = append(table, "9223372036854775807 * 3")
	for _, s := range table {
		z01.ChallengeMain(t, strings.Fields(s)...)
	}
}
