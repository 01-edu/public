## changeorder

## **WARNING! VERY IMPORTANT!**

For this exercise a function will be tested **with the exam own main**. However the student **still needs** to submit a structured program:

This means that:

- The package needs to be named `package main`.
- The submitted code needs one declared function main(```func main()```) even if empty.
- The function main declared needs to **also pass** the `Restrictions Checker`(illegal functions tester). It is advised for the student to just empty the function main after its own testing are done.
- Every other rules are obviously the same than for a `program`.

### Instructions

You are given a linked list, where each node contains a single digit.
Change order of linked list so that elements with odd index come first, elements
with even index come afterwards.
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

import (
	"fmt"
)

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
$ go build
$ ./main
1 -> 3 -> 5 -> 2 -> 4
```
