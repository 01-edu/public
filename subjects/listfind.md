## listpushback

### Instructions

Write a function `ListFind` that returns the value of the first link that the function in the arguments its equal.

- For this you shoud use the function `CompStr`.

- Use pointers when ever you can.

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

func CompStr(l *list) bool {

}

func ListFind(l *List, comp func(l *List) bool) *interface{} {

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
	link := &List{}

	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, "hello")
	piscine.ListPushBack(link, "hello2")
	piscine.ListPushBack(link, "hello3")

	fmt.Println(piscine.ListFind(link, CompStr))
}
```

And its output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
hello
student@ubuntu:~/piscine/test$
```
