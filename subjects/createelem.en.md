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

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	n := &node{}
	n.CreateElem(1234)
	fmt.Println(n)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
&{1234}
student@ubuntu:~/piscine-go/test$
```
