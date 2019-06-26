## sortlistinsert

### Instructions

Write a function `SortListInsert` that inserts `data_ref` in the linked list, but it as to remain sorted in ascending order.

- The list as to be alredy sorted.

- Use pointers when ever you can.

### Expected function and structure

```go
func SortListInsert(l *NodeI, data_ref int) *NodeI{

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

	link = listPushBack(link, 1)
	link = listPushBack(link, 4)
	link = listPushBack(link, 9)

	PrintList(link)

	link = piscine.SortListInsert(link, -2)
	link = piscine.SortListInsert(link, 2)
	PrintList(link)
}
```

And its output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
1 -> 4 -> 9 -> <nil>
-2 -> 1 -> 2 -> 4 -> 9 -> <nil>
student@ubuntu:~/piscine/test$
```
