## listsize

### Instructions

Write a function `ListSize` that returns the number of elements in a linked list `l`.

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

func ListSize(l *List) int {

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
	link := &piscine.List{}

	piscine.ListPushFront(link, "Hello")
	piscine.ListPushFront(link, "2")
	piscine.ListPushFront(link, "you")
	piscine.ListPushFront(link, "man")

	fmt.Println(piscine.ListSize(link))
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
4
student@ubuntu:~/piscine-go/test$
```
