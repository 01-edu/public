## listpushback

### Instructions

Write a function `ListSort` that sorts the linked list by ascending order.

- This time you only will have the `node` structure.

- Try to use recursive.

- Use pointers when ever you can.

### Expected function and structure

```go
type Nodee struct {
	Data int
	Next *Nodee
}

func ListSort(l *Nodee) *Nodee {

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

func PrintList(l *piscine.Nodee) {
	m := l
	for m != nil {
		fmt.Print(m.Data, " -> ")
		m = m.Next
	}
	fmt.Print(nil)
	fmt.Println()
}

func listPushBack(l *piscine.Nodee, data int) *piscine.Nodee {
	n := &piscine.Nodee{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {
	var link *piscine.Nodee

	link = listPushBack(link, 5)
	link = listPushBack(link, 4)
	link = listPushBack(link, 3)
	link = listPushBack(link, 2)
	link = listPushBack(link, 1)

	PrintList(piscine.ListSort(link))
}
```

And its output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
1 -> 2 -> 3 -> 4 -> 5 -> <nil>
student@ubuntu:~/piscine/test$
```
