## listpushback

### Instructions

Write a function `SortListInsert` that inserts `data_ref` in the linked list, but it as to remain sorted in ascending order.

- The list as to be alredy sorted.

- Use pointers when ever you can.

### Expected function and structure

```go
func SortListInsert(l *Nodee, data_ref int) *Nodee{

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

func PrintList(l *piscine.Nodee) {
	m := l
	for m != nil {
		fmt.Print(m.Data, " -> ")
		m = m.Next
	}
	fmt.Print(nil)
	fmt.Println()
}

func listPushBack(l *piscine.Nodee, data int) *piscine.Nodee {
	n := &piscine.Nodee{Data: data}

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

	var link *piscine.Nodee

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
