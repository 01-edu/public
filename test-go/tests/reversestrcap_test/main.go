package main

import (
	"github.com/01-edu/public/test-go/lib"
)

func main() {
	args := [][]string{
		{"First SMALL TesT"},
		{"SEconD Test IS a LItTLE EasIEr", "bEwaRe IT'S NoT HARd WhEN ", " Go a dernier 0123456789 for the road e"},
		{""},
	}

	for i := 0; i < 15; i++ {
		args = append(args, lib.MultRandAlnum())
	}

	for _, v := range args {
		lib.ChallengeMain("reversestrcap", v...)
	}
	lib.ChallengeMain("reversestrcap")
}
