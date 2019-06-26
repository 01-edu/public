## listpushback

### Instructions

Write a function `SortedListMerge` that mereges two lists, `n1` and `n2`, but it as to join them in ascending order.

- Tip each list as to be already sorted.

- Use pointers when ever you can.

### Expected function and structure

```go
func SortedListMerge(n1 *Nodee, n2 *Nodee) *Nodee {

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

type node = piscine.Nodee
type nodes = piscine.Nodee

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
