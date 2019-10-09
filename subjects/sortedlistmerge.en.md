## sortedlistmerge

### Instructions

Write a function `SortedListMerge` that merges two lists `n1` and `n2` in ascending order.

-   During the tests `n1` and `n2` will already be initially sorted.

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
	var link2 *piscine.NodeI

	link = listPushBack(link, 3)
	link = listPushBack(link, 5)
	link = listPushBack(link, 7)

	link2 = listPushBack(link2, -2)
	link2 = listPushBack(link2, 9)

	PrintList(piscine.SortedListMerge(link2, link))
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
-2 -> 3 -> 5 -> 7 -> 9 -> <nil>
student@ubuntu:~/piscine-go/test$
```
