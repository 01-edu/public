## listpushback

### Instructions

Write a function `ListPushBack` that inserts a new element `NodeL` at the end of the list `l` while using the structure `List`.

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

func ListPushBack(l *List, data interface{}) {

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

	piscine.ListPushBack(link, "Hello")
	piscine.ListPushBack(link, "man")
	piscine.ListPushBack(link, "how are you")

	for link.Head != nil {
		fmt.Println(link.Head.Data)
		link.Head = link.Head.Next
	}
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
Hello
man
how are you
student@ubuntu:~/piscine-go/test$
```
