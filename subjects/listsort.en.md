# listpushback

## Instructions

Write a function `ListSort` that sorts the linked list by ascending order.

- This time you only will have the `node` structure.

- Try to use recursive.

- Use pointers when ever you can.

## Expected function and structure

```go
type node struct {
	data int
	next *node
}

func ListSort(l *node) *node {

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

	listPushBack(link, 5)
	listPushBack(link, 4)
	listPushBack(link, 3)
	listPushBack(link, 2)
	listPushBack(link, 1)

	PrintList(piscine.ListSort(link))

}

```

And its output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
0 -> 1 -> 2 -> 3 -> 4 -> 5 -> <nil>
student@ubuntu:~/piscine/test$
```
