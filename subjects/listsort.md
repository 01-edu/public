# listpushback

## Instructions

Write a function `ListSort` that sorts the linked list by ascending order.

- This time you only will have the `Nodee` structure. 

- Try to use recursive.

- Use pointers when ever you can.

## Expected function and structure

```go
type Nodee struct {
	Data int
	Next *Nodee
}

func ListSort(l *NodeL) *NodeL {

}
```

## Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
	"fmt"
	student ".."
)

//Prints the list
func PrintList(l *Nodee) {
	m := l
	for m != nil {
		fmt.Print(m.Data, " -> ")
		m = m.next
	}

	fmt.Print(nil)
	fmt.Println()
}

//insert elements
func listPushBack(l *Nodee, Data int) {

	n := &Nodee{}
	n.Data = Data
	n.next = nil

	if l == nil {
		l = n
		return
	}

	iterator := l
	for iterator.next != nil {
		iterator = iterator.next
	}
	iterator.next = n
}

func main() {
	link := &Nodee{}

	listPushBack(link, 5)
	listPushBack(link, 4)
	listPushBack(link, 3)
	listPushBack(link, 2)
	listPushBack(link, 1)

	PrintList(student.ListSort(link))

}

```

And its output :

```console
student@ubuntu:~/student/test$ go build
student@ubuntu:~/student/test$ ./test
0 -> 1 -> 2 -> 3 -> 4 -> 5 -> <nil>
student@ubuntu:~/student/test$
```
