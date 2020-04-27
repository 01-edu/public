package main

import (
	"fmt"
	"os"
)

const SIZE = 2048

func main() {
	if len(os.Args) != 2 {
		return
	}
	progpoint := []byte(os.Args[1])
	var arby [SIZE]byte
	pos := 0
	openBr := 0         // opened brackets
	i := 0              // iterates through the source code passed in the argument
	N := len(progpoint) // length of the source code
	for i >= 0 && i < N {
		switch progpoint[i] {
		case '>':
			// Increment the pointer
			pos++
		case '<':
			// decrement the pointes
			pos--
		case '+':
			// increment the pointed byte
			arby[pos]++
		case '-':
			// decrement the pointed byte
			arby[pos]--
		case '.':
			// print the pointed byte on std output
			fmt.Printf("%c", rune(arby[pos]))
		case '[':
			// go to the matching ']' if the pointed byte is 0 (while start)
			openBr = 0
			if arby[pos] == 0 {
				for i < N && (progpoint[i] != byte(']') || openBr > 1) {
					if progpoint[i] == byte('[') {
						openBr++
					} else if progpoint[i] == byte(']') {
						openBr--
					}
					i++
				}
			}
		case ']':
			// go to the matching '[' if the pointed byte is not 0 (while end)
			openBr = 0
			if arby[pos] != 0 {
				for i >= 0 && (progpoint[i] != byte('[') || openBr > 1) {
					if progpoint[i] == byte(']') {
						openBr++
					} else if progpoint[i] == byte('[') {
						openBr--
					}
					i--
				}
			}
		}
		i++
	}
}
