## sortlistinsert

### Instructions

Write a function `SortListInsert` that inserts `data_ref` in the linked list `l` while keeping the list sorted in ascending order.

-   During the tests the list passed as an argument will be already sorted.

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
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
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
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
1 -> 4 -> 9 -> <nil>
-2 -> 1 -> 2 -> 4 -> 9 -> <nil>
student@ubuntu:~/piscine-go/test$
```
