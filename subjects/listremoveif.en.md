# listpushback

## Instructions

Write a function `ListRemoveIf` that removes all elements that are equal to the `data_ref` introduced in the argument of the function.

- In case the list is empty print the message `no data on list`.

- Use pointers wen ever you can.

## Expected function and structure

```go
type node struct {
	data interface{}
	next *node
}

type list struct {
	head *node
	tail *node
}

func ListRemoveIf(l *list, data_ref interface{}) {

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

func PrintList(l *list) {
	m := l.head
	for m != nil {
		fmt.Print(m.data, " -> ")
		m = m.next
	}

	fmt.Print(l.tail)
	fmt.Println()
}

func main() {
	link := &list{}
	link2 := &list{}
	link3 := &list{}

	fmt.Println("------answer-----")
	ListRemoveIf(link3, 1)
	fmt.Println()

	fmt.Println("----normal state----")
	piscine.ListPushBack(link2, 1)
	PrintList(link2)
	ListRemoveIf(link2, 1)
	fmt.Println("------answer-----")
	PrintList(link)
	fmt.Println()

	fmt.Println("----normal state----")
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, "Hello")
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, "There")
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, "How")
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, "are")
	piscine.ListPushBack(link, "you")
	piscine.ListPushBack(link, 1)
	PrintList(link)

	ListRemoveIf(link, 1)
	fmt.Println("------answer-----")
	PrintList(link)
}

```

And its output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
------answer-----
no data on list

----normal state----
1 -> <nil>
------answer-----
<nil>

----normal state----
1 -> Hello -> 1 -> There -> 1 -> 1 -> How -> 1 -> are -> you -> 1 -> <nil>
------answer-----
Hello -> There -> How -> are -> you -> <nil>
student@ubuntu:~/piscine/test$
```
