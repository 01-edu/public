## btreemax

### Instructions

Write a function, BTreeMax, that returns the node with the maximum value in the tree given by root

This function must have the following signature.

### Expected function

```go
func BTreeMax(root *TreeNode) *TreeNode {
	
}
```

### Usage

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
	max := student.BTreeMax(root)
	fmt.Println(max.Data)
}
```

And its output :

```console
student@ubuntu:~/student/btreemax$ go build
student@ubuntu:~/student/btreemax$ ./btreemax
7
student@ubuntu:~/student/btreemax$ 
```
