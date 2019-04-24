# listpushback

## Instructions

Write a function `ListForEach` that applies a function given as argument to the information within each of the list's links.

- The function given as argument must have a pointer as argument: `l *list`

## Expected function and structure

```go
type node struct {
	data interface{}
	next *node
}

type list struct {
	head *node
	tail *node
}

func ListForEach(l *list, f func(l *list)) {
}
```

## Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	link := &list{}

	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, 2)
	piscine.ListPushBack(link, 3)
	piscine.ListPushBack(link, 4)

	piscine.ListForEach(link, piscine.ListReverse)

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
