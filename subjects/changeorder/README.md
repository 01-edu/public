## changeorder

### Instructions

You are given a linked list, where each node contains a single digit.
Change order of linked list so that elements with odd index come first, elements with even index come afterwards.
You have to return pointer/reference to the beginning of new list

### Expected function and structure

```go
package main

type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

func Changeorder(node *NodeAddL) *NodeAddL {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import "fmt"

// I implemented pushBack for this

func main() {
	num1 := &NodeAddL{Num: 1}
	num1 = pushBack(num1, 2)
	num1 = pushBack(num1, 3)
	num1 = pushBack(num1, 4)
	num1 = pushBack(num1, 5)

	result := Changeorder(num1)
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
1 -> 3 -> 5 -> 2 -> 4
$
```
