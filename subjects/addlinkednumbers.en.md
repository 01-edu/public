## addlinkednumbers

## **WARNING! VERY IMPORTANT!**

For this exercise a function will be tested **with the exam own main**. However the student **still needs** to submit a structured program:

This means that:

- The package needs to be named `package main`.
- The submitted code needs one declared function main(```func main()```) even if empty.
- The function main declared needs to **also pass** the `Restrictions Checker`(illegal functions tester). It is advised for the student to just empty the function main after its own testing are done.
- Every other rules are obviously the same than for a `program`.

### Instructions

You have two numbers represented by a linked list, where each NodeAddL contains a single digit.
The digits are stored in reverse order, such that the 1’s digit is at the head of the list.
Write a function that adds the two numbers and returns the sum as a linked list

### Expected function and structure

```go
package main

type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

func AddLinkedNumbers(num1, num1 *NodeAddL) *NodeAddL {

}
```

### Usage

Here is a possible program to test your function:

```go
package main

import (
	"fmt"
)

func pushFront(node *NodeAddL, num int) *NodeAddL {

}

func main() {
	// 3 -> 1 -> 5
	num1 := &NodeAddL{Num:5}
	num1 = pushFront(num1, 1)
	num1 = pushFront(num1, 3)

	// 5 -> 9 -> 2
	num2 := &NodeAddL{Num:2}
	num2 = pushFront(num2, 9)
	num2 = pushFront(num2, 5)

	// 9 -> 0 -> 7
	result := AddLinkedNumbers(num1, num2)
	for tmp := result; tmp != nil; tmp = tmp.Next {
		fmt.Print(tmp.Num)
		if tmp.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
```

An its output:

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./main
9 -> 0 -> 7
student@ubuntu:~/[[ROOT]]/test$
```
