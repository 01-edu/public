package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type roman struct {
	num        int
	romanDigit string
}

func main() {
	if len(os.Args) == 2 {
		nbr, err := strconv.Atoi(os.Args[1])
		if err != nil || nbr >= 4000 || nbr == 0 {
			fmt.Println("ERROR: can not convert to roman digit")
			os.Exit(0)
		}
		patter := []roman{
			{num: 1000, romanDigit: "M"},
			{num: 900, romanDigit: "CM"},
			{num: 500, romanDigit: "D"},
			{num: 400, romanDigit: "CD"},
			{num: 100, romanDigit: "C"},
			{num: 90, romanDigit: "XC"},
			{num: 50, romanDigit: "L"},
			{num: 40, romanDigit: "XL"},
			{num: 10, romanDigit: "X"},
			{num: 9, romanDigit: "IX"},
			{num: 5, romanDigit: "V"},
			{num: 4, romanDigit: "IV"},
			{num: 1, romanDigit: "I"},
		}
		sumRoman, romandigit := print(nbr, patter)
		fmt.Println(strings.TrimSuffix(sumRoman, "+"))
		fmt.Println(romandigit)
	}
}

func print(nbr int, patter []roman) (string, string) {
	var sumRomanDigit, result string
	for _, v := range patter {
		for nbr >= v.num {
			sumRomanDigit += v.romanDigit + "+"
			result += v.romanDigit
			nbr -= v.num
		}
	}
	sumRomanDigit = formatsum(sumRomanDigit, patter)
	return sumRomanDigit, result
}

func formatsum(a string, patter []roman) string {
	result2 := strings.Split(a, "+")

	for i, v := range result2 {
		if len(v) == 2 {
			result2[i] = fmt.Sprintf("(%s-%s)", string(result2[i][1]), string(result2[i][0]))
		}
	}
	a = strings.Join(result2, "+")
	return a
}
