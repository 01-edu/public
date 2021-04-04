package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func singleSearch(exp []string, text string) []string {
	items := strings.Split(text, " ")
	var result []string

	for _, item := range items {
		for _, word := range exp {
			if strings.Contains(item, word) {
				result = append(result, item)
			}
		}
	}
	return result
}

func simpleSearch(runes []rune, text string) []string {
	exp := string(runes)

	var result []string
	if !strings.ContainsRune(exp, '|') {
		helper := []string{exp}
		result = append(singleSearch(helper, text))
	} else {
		expWords := strings.Split(exp, "|")
		result = append(result, singleSearch(expWords, text)...)
	}
	return result
}

func brackets(regexp, text string) {
	if text == "" || regexp == "" {
		return
	}
	runes := []rune(regexp)

	if runes[0] == '(' && runes[len(runes)-1] == ')' {
		runes = runes[1 : len(runes)-1]
		result := simpleSearch(runes, text)
		for i, s := range result {
			if !unicode.IsLetter(rune(s[len(s)-1])) {
				s = s[0 : len(s)-1]
			}
			fmt.Printf("%d: %s\n", i+1, s)
		}
	}
}

func main() {
	if len(os.Args) == 3 {
		brackets(os.Args[1], os.Args[2])
	}
}
