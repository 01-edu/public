## listsort

### Instructions

Write a function `ListSort` that sorts the linked list by ascending order.

- This time you only will have the `node` structure.

- Try to use recursive.

- Use pointers when ever you can.

### Expected function and structure

```go
type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {

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

func PrintList(l *piscine.NodeI) {
	m := l
	for m != nil {
		fmt.Print(m.Data, " -> ")
		m = m.Next
	}
	fmt.Print(nil)
	fmt.Println()
}

func listPushBack(l *piscine.NodeI, data int) *piscine.NodeI {
	n := &piscine.NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {
	var link *piscine.NodeI

	link = listPushBack(link, 5)
	link = listPushBack(link, 4)
	link = listPushBack(link, 3)
	link = listPushBack(link, 2)
	link = listPushBack(link, 1)

	PrintList(piscine.ListSort(link))
}
```

And its output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
1 -> 2 -> 3 -> 4 -> 5 -> <nil>
student@ubuntu:~/piscine/test$
```
