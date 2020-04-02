package main

func StrLen(str string) int {
	len := 0

	strConverted := []rune(str)
	for i := range strConverted {
		len = i + 1
	}
	return len
}

func main() {
}
