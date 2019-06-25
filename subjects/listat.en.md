## listpushback

### Instructions

Write a function `ListAt` that takes one pointer to the list, `l`, and an `int` as parameters. This function should print a `NodeL` in the position `pos` in the linked list.

- In case of error it should print `nil`

### Expected function and structure

```go
type NodeL struct {
	Data interface{}
	Next *NodeL
}


func ListAt(l *NodeL, pos int) *NodeL{

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

	piscine.ListPushBack(link, "hello")
	piscine.ListPushBack(link, "how are")
	piscine.ListPushBack(link, "you")
	piscine.ListPushBack(link, 1)

	fmt.Println(piscine.ListAt(link.Head, 3).Data)
	fmt.Println(piscine.ListAt(link.Head, 1).Data)
	fmt.Println(piscine.ListAt(link.Head, 7))
}

```

And its output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
1
how are
<nil>
student@ubuntu:~/piscine/test$
```