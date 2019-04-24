# listpushback

## Instructions

Write a function `SortListInsert` that inserts `data_ref` in the linked list, but it as to remain sorted in ascending order.

- The list as to be alredy sorted.

- Use pointers when ever you can.

## Expected function and structure

```go
type node struct {
	data int
	next *node
}

func SortListInsert(l *node, data_ref int) *node{

}
```

## Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
	"fmt"
	piscine ".."
)

//Prints the list
func PrintList(l *node) {
	m := l
	for m != nil {
		fmt.Print(m.data, " -> ")
		m = m.next
	}
	fmt.Print(nil)
	fmt.Println()
}
//insert elements
func listPushBack(l *node, data int) {
	n := &node{}
	n.data = data
	n.next = nil
	if l == nil {
		l = n
		return
	}
	iterator := l
	for iterator.next != nil {
		iterator = iterator.next
	}
	iterator.next = n
}

func main() {

	link := &node{}

	listPushBack(link, 1)
	listPushBack(link, 4)
	listPushBack(link, 9)

	PrintList(link)

	link = sortListInsert(link, -2)
	link = sortListInsert(link, 2)
	PrintList(link)
}

```

And its output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
-2 -> 0 -> 1 -> 2 -> 4 -> 9 -> <nil>
lee@lee:~/Documents/work/day11/11-16-sortlistinsert/so
student@ubuntu:~/piscine/test$
```
