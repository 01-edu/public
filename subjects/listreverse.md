## listpushback

### Instructions

Write a function `ListReverse` that reverses the elements order of a given linked list.

- Use pointers when ever you can

### Expected function and structure

```go
type Node struct {
	Data interface{}
	Next *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func ListReverse(l *list) {
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

func main() {
	link := &List{}

	listPushBack(link, 1)
	listPushBack(link, 2)
	listPushBack(link, 3)
	listPushBack(link, 4)

	listReverse(link)

	for link.Head != nil {
		fmt.Println(link.Head.Data)
		link.Head = link.Head.Next
	}
}
```

And its output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
4
3
2
1
student@ubuntu:~/piscine/test$
```
