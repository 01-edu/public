package main

import (
	"strings"
)

func DoppelGanger(big, little string) int {
	if little == "" {
		return -1
	}
	return strings.LastIndex(big, little)
}

func main() {

}
