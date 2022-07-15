## listlast

### Instructions

Write a function `ListLast` that returns the `Data` interface of the last element of a linked list `l`.

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

func ListLast(l *List) interface{} {

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
	link2 := &piscine.List{}

	piscine.ListPushBack(link, "three")
	piscine.ListPushBack(link, 3)
	piscine.ListPushBack(link, "1")

	fmt.Println(piscine.ListLast(link))
	fmt.Println(piscine.ListLast(link2))
}
```

And its output :

```console
$ go run .
1
<nil>
$
```
