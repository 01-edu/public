## listpushback

### Instructions

Write a function `ListReverse` that reverses the elements order of a given linked list.

- Use pointers when ever you can

### Expected function and structure

```go
type node struct {
	data interface{}
	next *node
}

type list struct {
	head *node
	tail *node
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
	link := &list{}

	listPushBack(link, 1)
	listPushBack(link, 2)
	listPushBack(link, 3)
	listPushBack(link, 4)

	listReverse(link)

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
4
3
2
1
student@ubuntu:~/piscine/test$
```
