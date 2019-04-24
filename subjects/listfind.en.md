# listpushback

## Instructions

Write a function `ListFind` that returns the address of the first link that the function in the arguments its equal.

- For this you shoud use the function `CompStr`.

- Use pointers wen ever you can.

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

func CompStr(l *list) bool {

}

func ListFind(l *list, comp func(l *list) bool) *interface{} {

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

	piscine.ListPushBack(link, "hello")
	piscine.ListPushBack(link, "hello1")
	piscine.ListPushBack(link, "hello2")
	piscine.ListPushBack(link, "hello3")

	fmt.Println(piscine.ListFind(link, compStr))
}
```

And its output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
0xc42000a0a0
student@ubuntu:~/piscine/test$
```
