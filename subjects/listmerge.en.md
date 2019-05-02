## listpushback

### Instructions

Write a function `ListMerge` that places elements of a list `l2` at the end of an other list `l1`.

- You can't create new elements!

- Use pointers when ever you can.

### Expected function and structure

```go
type node struct {
	data interface{}
	next *node
}

type list struct {
	head *node
	tail *node
}

func ListMerge(l1 *list, l2 *list) {

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

	fmt.Print(l.tail)
	fmt.Println()
}

func main() {
	link := &list{}
	link2 := &list{}

	piscine.ListPushBack(link, "a")
	piscine.ListPushBack(link, "b")
	piscine.ListPushBack(link, "c")
	piscine.ListPushBack(link, "d")

	piscine.ListPushBack(link2, "e")
	piscine.ListPushBack(link2, "f")
	piscine.ListPushBack(link2, "g")
	piscine.ListPushBack(link2, "h")

	piscine.ListMerge(link, link2)
	PrintList(link)
}
```

And its output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
a -> b -> c -> d -> e -> f -> g -> h -> <nil>
student@ubuntu:~/piscine/test$
```
