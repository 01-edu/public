## listpushback

### Instructions

Write a function `ListMerge` that places elements of a list `l2` at the end of an other list `l1`.

- You can't create new elements!

- Use pointers when ever you can.

### Expected function and structure

```go
type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func listMerge(l1 *List, l2 *List) {
	
}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
	"fmt"
	student ".."
)

func PrintList(l *List) {
	m := l.Head
	for m != nil {
		fmt.Print(m.Data, " -> ")
		m = m.Next
	}

	fmt.Print(nil)
	fmt.Println()
}

func main() {
	link := &List{}
	link2 := &List{}

	student.ListPushBack(link, "a")
	student.ListPushBack(link, "b")
	student.ListPushBack(link, "c")
	student.ListPushBack(link, "d")

	student.ListPushBack(link2, "e")
	student.ListPushBack(link2, "f")
	student.ListPushBack(link2, "g")
	student.ListPushBack(link2, "h")

	student.ListMerge(link, link2)
	PrintList(link)
}
```

And its output :

```console
student@ubuntu:~/student/test$ go build
student@ubuntu:~/student/test$ ./test
a -> b -> c -> d -> e -> f -> g -> h -> <nil>
student@ubuntu:~/student/test$
```
