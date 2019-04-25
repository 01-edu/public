## listpushback

### Instructions

Write a function `ListPushBack` that inserts a new element `node` at the beginning of the list using `list`

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

func ListPushFront(l *list, data interface{}) {
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

	link := &list{}

	piscine.ListPushFront(link, "Hello")
	piscine.ListPushFront(link, "man")
	piscine.ListPushFront(link, "how are you")

	for link.head != nil {
		fmt.Println(link.head.data)
		link.head = link.head.next
	}
}
```

And its output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
how are you
man
Hello
student@ubuntu:~/piscine/test$
```
