## listclear

### Instructions

Write a function `ListClear` that deletes all `nodes` from a linked list `l`.

-   Tip: assign the list's pointer to `nil`.

### Expected function and structure

```go
func ListClear(l *List) {

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

type List = piscine.List
type Node = piscine.NodeL

func PrintList(l *List) {
	link := l.Head
	for link != nil {
		fmt.Print(link.Data, " -> ")
		link = link.Next
	}
	fmt.Println(nil)
}

func main() {
	link := &List{}

	piscine.ListPushBack(link, "I")
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, "something")
	piscine.ListPushBack(link, 2)

	fmt.Println("------list------")
	PrintList(link)
	piscine.ListClear(link)
	fmt.Println("------updated list------")
	PrintList(link)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
------list------
I -> 1 -> something -> 2 -> <nil>
------updated list------
<nil>
student@ubuntu:~/piscine-go/test$
```
