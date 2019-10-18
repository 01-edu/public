## listreverse

### Instructions

Write a function `ListReverse` that reverses the order of the elements of a given linked list `l`.

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

func ListReverse(l *List) {

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

	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, 2)
	piscine.ListPushBack(link, 3)
	piscine.ListPushBack(link, 4)

	piscine.ListReverse(link)

	it := link.Head

	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}

	fmt.Println("Tail", link.Tail)
	fmt.Println("Head", link.Head)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
4
3
2
1
Tail &{1 <nil>}
Head &{4 0xc42000a140}
student@ubuntu:~/piscine-go/test$
```
