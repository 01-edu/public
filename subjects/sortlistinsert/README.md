## sortlistinsert

### Instructions

Create a function named `SortListInsert`, which accepts a pointer to the head of a _sorted_ linked list and an integer.

The function should insert a new element into the list, with its `Data` set to the value of the integer. The element should be inserted so that the linked list remains sorted in ascending order by `Data`. The function should return the head of the list.

Your function will only be tested with sorted linked lists.

### Expected function and structure

```go
func SortListInsert(l *NodeI, data_ref int) *NodeI{

}
```

> You have already defined the `NodeI` structure in the `listsort` exercise.

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"

	"piscine"
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
$ go run .
1 -> 4 -> 9 -> <nil>
-2 -> 1 -> 2 -> 4 -> 9 -> <nil>
$
```
