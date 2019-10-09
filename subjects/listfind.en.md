## listfind

### Instructions

Write a function `ListFind` that returns the address of the first node in the list `l` that is determined to be equal to `ref` by the function `CompStr`.

-   For this exercise the function `CompStr` must be used.

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

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {

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
	piscine.ListPushBack(link, "hello1")
	piscine.ListPushBack(link, "hello2")
	piscine.ListPushBack(link, "hello3")

	found := piscine.ListFind(link, interface{}("hello2"), piscine.CompStr)

	fmt.Println(found)
	fmt.Println(*found)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
0xc42000a0a0
hello2
student@ubuntu:~/piscine-go/test$
```

### Note

-   The address may be different in each execution of the program.
