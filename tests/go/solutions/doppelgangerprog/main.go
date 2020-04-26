package main

import "strings"

func DoppelGanger(s, substr string) int {
	return strings.LastIndex(s, substr)
}

func main() {
}
