package student_test

import (
	"math/rand"
	"testing"

	"github.com/01-edu/z01"
)

//randomValidBase function is used to create the tests (input VALID bases here)
func randomValidBase() string {
	validBases := []string{
		"01",
		"CHOUMIisDAcat!",
		"choumi",
		"0123456789",
		"abc", "Zone01",
		"0123456789ABCDEF",
		"WhoAmI?",
	}
	index := rand.Intn(len(validBases))
	return validBases[index]
}

//randomInvalidBase function is used to create the tests (input INVALID bases here)
func randomInvalidBase() string {
	invalidBases := []string{
		"0",
		"1",
		"CHOUMIisdacat!",
		"choumiChoumi",
		"01234567890",
		"abca",
		"Zone01Zone01",
		"0123456789ABCDEF0",
		"WhoAmI?IamWhoIam",
	}
	index := z01.RandIntBetween(0, len(invalidBases)-1)
	return invalidBases[index]
}

//randomStringFromBase function is used to create the random STRING number from VALID BASES
func randomStringFromBase(base string) string {
	letters := []rune(base)
	size := z01.RandIntBetween(1, 10)
	r := make([]rune, size)
	for i := range r {
		r[i] = letters[rand.Intn(len(letters))]
	}
	return string(r)
}

// this is the function that creates the TESTS
func TestAtoiBaseProg(t *testing.T) {
	type node struct {
		s    string
		base string
	}

	table := []node{}

	// 5 random pairs of string numbers with valid bases
	for i := 0; i < 5; i++ {
		validBaseToInput := randomValidBase()
		val := node{
			s:    randomStringFromBase(validBaseToInput),
			base: validBaseToInput,
		}
		table = append(table, val)
	}
	// 5 random pairs of string numbers with invalid bases
	for i := 0; i < 5; i++ {
		invalidBaseToInput := randomInvalidBase()
		val := node{
			s:    "thisinputshouldnotmatter",
			base: invalidBaseToInput,
		}
		table = append(table, val)
	}
	table = append(table,
		node{s: "125", base: "0123456789"},
		node{s: "1111101", base: "01"},
		node{s: "7D", base: "0123456789ABCDEF"},
		node{s: "uoi", base: "choumi"},
		node{s: "bbbbbab", base: "-ab"},
	)
	for _, arg := range table {
		z01.ChallengeMain(t, arg.s, arg.base)
	}
	z01.ChallengeMain(t)
	z01.ChallengeMain(t, "125", "0123456789", "something")
}
