# listpushback

## Instructions

Write a function `SortedListMerge` that mereges two lists, `l1` and `l2`, but you have to join them in ascending order.

- Tip each list as to be already sorted, and initialized with 0.

- Use pointers when ever you can.

## Expected function and structure

```go
type Nodee struct {
	Data interface{}
	Next *Nodee
}

func SortedListMerge(l1 *Nodee, l2 *Nodee) *Nodee {

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

func PrintList(l *Nodee) {
	m := l
	for m != nil {
		fmt.Print(m.Data, " -> ")
		m = m.Next
	}

	fmt.Print(nil)
	fmt.Println()
}

func main() {
	link := &Nodee{}
	link2 := &Nodee{}

	piscine.ListPushBack(link, "5")
	piscine.ListPushBack(link, "3")
	piscine.ListPushBack(link, "7")

	piscine.ListPushBack(link2, "1")
	piscine.ListPushBack(link2, "-2")
	piscine.ListPushBack(link2, "4")
	piscine.ListPushBack(link2, "6")

	PrintList(SortedListMerge(link, link2))
}
```

And its output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
-2 -> 0 -> 0 -> 1 -> 3 -> 4 -> 5 -> 6 -> 7 -> <nil>
student@ubuntu:~/piscine/test$
```
