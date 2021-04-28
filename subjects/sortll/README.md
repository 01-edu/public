## sortlinkedlist

### Instructions

You are given a linked list, where each node contains a single digit.
Write a function that sorts the list in descending order and return pointer/reference to new linked list

### Expected function and struct

```go
package main

type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

func Sortll(node *NodeAddL) *NodeAddL {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import "fmt"

// I implemented pushBack for this

func main() {
	num1 := &NodeAddL{Num: 5}
	num1 = pushBack(num1, 1)
	num1 = pushBack(num1, 3)
	num1 = pushBack(num1, 1)
	num1 = pushBack(num1, 3)

	result := Sortll(num1)
	for tmp := result; tmp != nil; tmp = tmp.Next {
		fmt.Print(tmp.Num)
		if tmp.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
```

Its output:

```console
$ go run .
5 -> 3 -> 3 -> 1 -> 1
$
```
