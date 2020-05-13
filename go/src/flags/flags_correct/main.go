package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type helpMs struct {
	flag        string
	shortenFlag string
	handler     string
}

func obtainValues(value, strsplit string) string {
	values := strings.Split(value, "=")
	return values[len(values)-1]
}

func setMs(flag, shortenFlag, handler string) *helpMs {
	helpMs := &helpMs{
		flag:        flag,
		shortenFlag: shortenFlag,
		handler:     handler,
	}
	return helpMs
}

func main() {
	size := len(os.Args)

	if size == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		table := []helpMs{}

		helpMs := setMs("--insert", "-i", "This flag inserts the string into the string passed as argument.")
		table = append(table, *helpMs)
		helpMs = setMs("--order", "-o", "This flag will behave like a boolean, if it is called it will order the argument.")
		table = append(table, *helpMs)

		for _, v := range table {
			fmt.Println(v.flag)
			fmt.Println(" ", v.shortenFlag)
			fmt.Println("	", v.handler)
		}
	} else if size <= 4 {
		var runes []rune
		strToInsert := ""
		var order bool

		for i := 1; i < size; i++ {
			if strings.Contains(os.Args[i], "--insert") || strings.Contains(os.Args[i], "-i") {
				strToInsert = obtainValues(os.Args[i], "=")
			} else if strings.Contains(os.Args[i], "--order") || strings.Contains(os.Args[i], "-o") {
				order = true
			} else {
				runes = []rune(os.Args[i])
			}
		}
		if strToInsert != "" {
			concatStr := string(runes) + strToInsert
			runes = []rune(concatStr)
		}
		if order {
			sort.Slice(runes, func(i, j int) bool {
				return runes[i] < runes[j]
			})
		}

		fmt.Println(string(runes))
	}
}
