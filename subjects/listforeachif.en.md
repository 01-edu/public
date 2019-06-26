## listpushback

### Instructions

Write a function `ListForEachIf` that applies a function given as argument to the information within some nodes of the list.

- This functions receives two functions:

  - `f` is a functions that is applied to the node.

  - `comp` is a predicate (a function that returns true or false) and will be use to determine if the function `f` would be applied to the node.

- The function given as argument must have a pointer as argument: `*NodeL`.

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

func CompStr(l *NodeL) bool {

}

func ListForEachIf(l *List, f func(*NodeL), comp func(*NodeL) bool) {

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

func PrintElem(l *List) {
	fmt.Println(l.Head.Data)
}

func StringToInt(l *List) {
	count := 1
	l.Head.Data = count
}

func PrintList(l *List) {
	m := l.Head
	for m != nil {
		fmt.Print(m.Data, " -> ")
		m = m.Next
	}

	fmt.Print(l.Tail)
}
func main() {
	link := &List{}

	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, "hello")
	piscine.ListPushBack(link, 3)
	piscine.ListPushBack(link, "there")
	piscine.ListPushBack(link, 23)
	piscine.ListPushBack(link, "!")
	piscine.ListPushBack(link, 54)

	PrintAllList(link)

	fmt.Println()
	fmt.Println("--------function applied--------")
	piscine.ListForEachIf(link, PrintElem, CompStr)

	piscine.ListForEachIf(link, StringToInt, CompStr)

	fmt.Println("--------function applied--------")
	PrintAllList(link)

	fmt.Println()
}
```

And its output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
1 -> hello -> 3 -> there -> 23 -> ! -> 54 -> <nil>
--------function applied--------
hello
there
!
--------function applied--------
1 -> 1 -> 3 -> 1 -> 23 -> 1 -> 54 -> <nil>
student@ubuntu:~/piscine/test$
```
