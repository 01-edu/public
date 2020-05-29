import (
	"math/rand"

	"lib"
)

func Valid() string {
	valid := []string{
		"01",
		"CHOUMIisDAcat!",
		"choumi",
		"0123456789",
		"abc",
		"Zone01",
		"0123456789ABCDEF",
		"WhoAmI?",
	}
	i := rand.Intn(len(valid))
	return valid[i]
}

func Invalid() string {
	invalid := []string{
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
	i := rand.Intn(len(invalid))
	return invalid[i]
}

func StringFrom(base string) string {
	letters := []rune(base)
	size := lib.RandIntBetween(1, 10)
	runes := make([]rune, size)
	for i := range runes {
		runes[i] = letters[rand.Intn(len(letters))]
	}
	return string(runes)
}

func ConvertNbr(n int, base string) string {
	var result string
	length := len(base)

	for n >= length {
		result = string(base[(n%length)]) + result
		n = n / length
	}
	result = string(base[n]) + result

	return result
}

func Convert(nbr, baseFrom, baseTo string) string {
	resultIntermediary := AtoiBase(nbr, baseFrom)

	resultFinal := ConvertNbr(resultIntermediary, baseTo)

	return resultFinal
}
