## createelem

### Instructions

Write a function `CreateElem` that creates a new element of type`Node`.

### Expected function and structure

```go
type Node struct {
	Data interface{}
}

func CreateElem(n *Node, value int) {

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
	n := &node{}
	n.CreateElem(1234)
	fmt.Println(n)
}
```

And its output :

```console
$ go run .
&{1234}
$
```
