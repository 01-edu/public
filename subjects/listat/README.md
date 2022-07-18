## listat

### Instructions

Write a function `ListAt` that takes a pointer to the head of the list `l` and an `int pos` as parameters. This function should return the pointer to the `NodeL` in the position `pos` of the linked list `l`.

- In case of error the function should return `nil`.

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

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	"piscine"
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
$ go run .
1
how are
<nil>
$
```
