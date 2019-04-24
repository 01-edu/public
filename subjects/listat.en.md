## listpushback

### Instructions

Write a function `ListAt` that haves one pointer to the list, `l`, and an `int` as parameters. This function should print a `Node` of the linked list, depending on the number, `nbr`.

- In case of error it should print `nil`

### Expected function and structure

```go
type Node struct {
	Data interface{}
	Next *Node
}


func ListAt(l *Node, nbr int) *Node{

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
	link := &Node{}

	ListPushBack(link, "hello")
	ListPushBack(link, "how are")
	ListPushBack(link, "you")
	ListPushBack(link, 1)

	fmt.Println()

	fmt.Println(ListAt(link, 3).Data)
	fmt.Println(ListAt(link, 1).Data)
	fmt.Println(ListAt(link, 7))
}

```

And its output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
you
hello
<nil>
student@ubuntu:~/piscine/test$
```
