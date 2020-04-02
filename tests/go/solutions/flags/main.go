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

type sortRunes []rune

//to implement the sort for a []rune it is necessary to
//implement Lessm, Swap and Len functions for the sort interface
func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
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
		var str []rune
		strToInsert := ""
		var order bool

		for i := 1; i < size; i++ {
			if strings.Contains(os.Args[i], "--insert") || strings.Contains(os.Args[i], "-i") {
				strToInsert = obtainValues(os.Args[i], "=")
			} else if strings.Contains(os.Args[i], "--order") || strings.Contains(os.Args[i], "-o") {
				order = true
			} else {
				str = []rune(os.Args[i])
			}
		}
		if strToInsert != "" {
			concatStr := string(str) + strToInsert
			str = []rune(concatStr)
		}
		if order == true {
			sort.Sort(sortRunes(str))
		}

		fmt.Println(string(str))
	}
}
