## listpushback

### Instructions

Write a function `ListPushBack` that inserts a new element `node` at the beginning of the list using `list`

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

func ListPushFront(l *List, data interface{}) {
}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
	piscine ".."
	"fmt"
)

func main() {

	link := &piscine.List{}

	piscine.ListPushFront(link, "Hello")
	piscine.ListPushFront(link, "man")
	piscine.ListPushFront(link, "how are you")

	it := link.Head
	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
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
