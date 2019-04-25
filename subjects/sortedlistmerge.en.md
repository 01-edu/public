## listpushback

### Instructions

Write a function `SortedListMerge` that mereges two lists, `l1` and `l2`, but it as to join them in ascending order.

- Tip each list as to be already sorted.

- Use pointers when ever you can.

### Expected function and structure

```go
type node struct {
	data interface{}
	next *node
}

func SortedListMerge(l1 *node, l2 *node) *node {

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

func PrintList(l *list) {
	m := l.head
	for m != nil {
		fmt.Print(m.data, " -> ")
		m = m.next
	}

	fmt.Print(nil)
	fmt.Println()
}

func main() {
	link := &list{}
	link2 := &list{}

	piscine.ListPushBack(link, "5")
	piscine.ListPushBack(link, "3")
	piscine.ListPushBack(link, "7")

	piscine.ListPushBack(link2, "1")
	piscine.ListPushBack(link2, "-2")
	piscine.ListPushBack(link2, "4")
	piscine.ListPushBack(link2, "6")

	PrintList(SortedListMerge(link, link2))
}
```

And its output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
-2 -> 0 -> 0 -> 1 -> 3 -> 4 -> 5 -> 6 -> 7 -> <nil>
student@ubuntu:~/piscine/test$
```
