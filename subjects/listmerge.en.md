## listmerge

### Instructions

Write a function `ListMerge` that places elements of a list `l2` at the end of another list `l1`.

-   New elements should not be created!

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

func ListMerge(l1 *List, l2 *List) {

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

func PrintList(l *piscine.List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func main() {
	link := &piscine.List{}
	link2 := &piscine.List{}

	piscine.ListPushBack(link, "a")
	piscine.ListPushBack(link, "b")
	piscine.ListPushBack(link, "c")
	piscine.ListPushBack(link, "d")
	fmt.Println("-----first List------")
	PrintList(link)

	piscine.ListPushBack(link2, "e")
	piscine.ListPushBack(link2, "f")
	piscine.ListPushBack(link2, "g")
	piscine.ListPushBack(link2, "h")
	fmt.Println("-----second List------")
	PrintList(link2)

	fmt.Println("-----Merged List-----")
	piscine.ListMerge(link, link2)
	PrintList(link)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
-----first List------
a -> b -> c -> d -> <nil>
-----second List------
e -> f -> g -> h -> <nil>
-----Merged List-----
a -> b -> c -> d -> e -> f -> g -> h -> <nil>
student@ubuntu:~/piscine-go/test$
```
