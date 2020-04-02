package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//brackets("al|b", "ale atg bar sim nao pro par impar") In JS it's used without brackets
	if len(os.Args) == 3 {
		brackets(os.Args[1], os.Args[2])
	} else {
		fmt.Println()
	}
}

func brackets(regexp, text string) {
	if len(text) == 0 || len(regexp) == 0 {
		fmt.Println()
		return
	}
	reArr := []rune(regexp)

	if len(reArr) != 0 && reArr[0] == '(' && reArr[len(reArr)-1] == ')' {
		reArr = reArr[1 : len(reArr)-1]
		result := simpleSearch(reArr, text)
		for i, results := range result {
			if !isAlphaNum(results[len(results)-1]) {
				results = results[:len(results)-1]
			}
			if !isAlphaNum(results[0]) {
				results = results[1:]
			}
			fmt.Printf("%d: %s\n", i+1, results)
		}
	} else {
		fmt.Println()
	}
}

func isAlphaNum(r byte) bool {
	return (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') ||
		(r >= '0' && r <= '9')
}

func simpleSearch(reArr []rune, text string) []string {
	exp := string(reArr)

	var result []string
	if !(strings.ContainsRune(exp, '|')) {
		helper := []string{exp}
		result = append(singleSearch(helper, text))
	} else {
		expWords := strings.Split(exp, "|")
		result = append(result, singleSearch(expWords, text)...)
	}
	return result
}

func singleSearch(exp []string, text string) []string {
	tArr := strings.Split(text, " ")
	var result []string

	for _, elem := range tArr {
		for _, word := range exp {
			if strings.Contains(elem, word) {
				result = append(result, elem)
			}
		}
	}
	return result
}
