## reverse

### Instructions

You are given a linked list, where each node contains a single digit.
Write a function that reverses the list and returns pointer/reference to new linked list

### Expected function and structure

```go
package main

type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

func Reverse(node *NodeAddL) *NodeAddL {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import "fmt"

func pushBack(n *NodeAddL, num int) *NodeAddL{

}

func main() {
	num1 := &piscine.NodeAddL{Num: 1}
	num1 = pushBack(num1, 3)
	num1 = pushBack(num1, 2)
	num1 = pushBack(num1, 4)
	num1 = pushBack(num1, 5)

	result := piscine.Reverse(num1)
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
5 -> 4 -> 2 -> 3 -> 1
$
```
