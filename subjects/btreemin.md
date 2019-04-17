# btreemin
## Instructions

Write a function, BTreeMin, that returns the node with the minimum value in the tree given by root

This function must have the following signature.

## Expected function

```go
func BTreeMin(root *TreeNode) *TreeNode {
	
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

func main() {
	root := &student.TreeNode{Data: "4"}
	student.BTreeInsertData(root, "1")
	student.BTreeInsertData(root, "7")
	student.BTreeInsertData(root, "5")
	min := student.BTreeMin(root)
	fmt.Println(min.Data)
}
```

And its output :

```console
student@ubuntu:~/student/btreemin$ go build
student@ubuntu:~/student/btreemin$ ./btreemin
1
student@ubuntu:~/student/btreemin$ 
```
