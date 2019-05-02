## listpushback

### Instructions

Write a function `ListForEachIf` that applies a function given as argument to the information within some links of the list.

- For this you will have to create a function `CompStr`, that returns a `bool`, to compare each elemente of the linked list, to see if it is a string, and than apply the function in the argument of `ListForEachIf`.

- The function given as argument as to have a pointer as argument: `l *list`.

- Use pointers wen ever you can.

### Expected function and structure

```go
type node struct {
	data interface{}
	next *node
}

type list struct {
	head *node
	tail *node
}

func CompStr(l *list) bool {

}

func ListForEachIf(l *list, f func(l *list), comp func(l *list) bool) {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
	"fmt"
	piscine ".."
)

func PrintElem(l *list) {
	fmt.Println(l.head.data)
}

func StringToInt(l *list) {
	count := 1
	l.head.data = count
}

func PrintList(l *list) {
	m := l.head
	for m != nil {
		fmt.Print(m.data, " -> ")
		m = m.next
	}

	fmt.Print(l.tail)
}
func main() {
	link := &list{}

	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, "hello")
	piscine.ListPushBack(link, 3)
	piscine.ListPushBack(link, "there")
	piscine.ListPushBack(link, 23)
	piscine.ListPushBack(link, "!")
	piscine.ListPushBack(link, 54)

	PrintAllList(link)

	fmt.Println()
	fmt.Println("--------function applied--------")
	piscine.ListForEachIf(link, PrintElem, CompStr)

	piscine.ListForEachIf(link, StringToInt, CompStr)

	fmt.Println("--------function applied--------")
	PrintAllList(link)

	fmt.Println()
}
```

And its output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
1 -> hello -> 3 -> there -> 23 -> ! -> 54 -> <nil>
--------function applied--------
hello
there
!
--------function applied--------
1 -> 1 -> 3 -> 1 -> 23 -> 1 -> 54 -> <nil>
student@ubuntu:~/piscine/test$
```
