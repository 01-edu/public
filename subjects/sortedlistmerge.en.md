## sortedlistmerge

### Instructions

Write a function `SortedListMerge` that mereges two lists, `n1` and `n2`, but it as to join them in ascending order.

- Tip each list as to be already sorted.

### Expected function and structure

```go
func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {

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

type node = piscine.NodeI
type nodes = piscine.NodeI

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
	var link *node
	var link2 *nodes

	link = listPushBack(link, 5)
	link = listPushBack(link, 3)
	link = listPushBack(link, 7)

	link2 = listPushBack(link2, -2)
	link2 = listPushBack(link2, 4)

	PrintList(piscine.SortedListMerge(link2, link))
}

```

And its output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
-2 -> 3 -> 4 -> 5 -> 7 -> <nil>
student@ubuntu:~/piscine/test$
```
