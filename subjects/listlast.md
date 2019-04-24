## listpushback

### Instructions

Write a function `ListLast` that returns the last element of the linked list.

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

func ListLast(l *list) *list {

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
	link2 := &list{}

	piscine.ListPushBack(link, "three")
	piscine.ListPushBack(link, 3)
	piscine.ListPushBack(link, "1")

	fmt.Println(piscine.ListLast(link).head)
	fmt.Println(piscine.ListLast(link2).head)
}
```

And its output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
&{1 <nil>}
<nil>
student@ubuntu:~/piscine/test$
```
