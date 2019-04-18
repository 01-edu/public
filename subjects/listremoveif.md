# listpushback

## Instructions

Write a function `ListRemoveIf` that removes all elements that are equal to the `data_ref` introduced in the argument of the function.

- In case the list is empty print the message `no data on list`.

- Use pointers wen ever you can.


## Expected function and structure

```go
type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {

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

func PrintList(l *List) {
	m := l.Head
	for m != nil {
		fmt.Print(m.Data, " -> ")
		m = m.Next
	}

	fmt.Print(nil)
	fmt.Println()
}

func main() {
	link := &List{}
	link2 := &List{}
	link3 := &List{}

	fmt.Println("------answer-----")
	ListRemoveIf(link3, 1)
	fmt.Println()

	fmt.Println("----normal state----")
	student.ListPushBack(link2, 1)
	PrintList(link2)
	ListRemoveIf(link2, 1)
	fmt.Println("------answer-----")
	PrintList(link)
	fmt.Println()

	fmt.Println("----normal state----")
	student.ListPushBack(link, 1)
	student.ListPushBack(link, "Hello")
	student.ListPushBack(link, 1)
	student.ListPushBack(link, "There")
	student.ListPushBack(link, 1)
	student.ListPushBack(link, 1)
	student.ListPushBack(link, "How")
	student.ListPushBack(link, 1)
	student.ListPushBack(link, "are")
	student.ListPushBack(link, "you")
	student.ListPushBack(link, 1)
	PrintList(link)

	ListRemoveIf(link, 1)
	fmt.Println("------answer-----")
	PrintList(link)
}

```

And its output :

```console
student@ubuntu:~/student/test$ go build
student@ubuntu:~/student/test$ ./test
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
