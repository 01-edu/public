## listpushback

### Instructions

Write a function `SortListInsert` that inserts `Data_ref` in the linked list, but it as to remain sorted in ascending order.

- The list as to be alredy sorted. 

- Use pointers when ever you can.

### Expected function and structure

```go
type Nodee struct {
	Data int
	Next *Nodee
}

func SortListInsert(l *Nodee, Data_ref int) *Nodee{

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

//Prints the list
func PrintList(l *Nodee) {
	m := l
	for m != nil {
		fmt.Print(m.Data, " -> ")
		m = m.Next
	}
	fmt.Print(nil)
	fmt.Println()
}
//insert elements
func listPushBack(l *Nodee, Data int) {
	n := &Nodee{}
	n.Data = Data
	n.Next = nil
	if l == nil {
		l = n
		return
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
}

func main() {

	link := &Nodee{}

	listPushBack(link, 1)
	listPushBack(link, 4)
	listPushBack(link, 9)

	PrintList(link)

	link = sortListInsert(link, -2)
	link = sortListInsert(link, 2)
	PrintList(link)
}

```

And its output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
-2 -> 0 -> 1 -> 2 -> 4 -> 9 -> <nil>
lee@lee:~/Documents/work/day11/11-16-sortlistinsert/so
student@ubuntu:~/piscine/test$
```
